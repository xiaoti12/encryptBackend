import time
from watchdog.observers import Observer
from watchdog.events import FileSystemEventHandler
import os
import parseTLS
import RandomForest
from FeatureConstant import INFO_COLUMN
import pandas as pd

from datetime import datetime

model_path = "/home/fsc/liujy/file_watch/train/rfc_nomkv.joblib"
scaler_path = "/home/fsc/liujy/file_watch/train/scaler_nomkv.joblib"
csv_path = "/home/fsc/liujy/file_watch/results.csv"


class MyHandler(FileSystemEventHandler):
    def __init__(self):
        super().__init__()

    def on_created(self, event):
        log_name = os.path.basename(event.src_path)
        if not event.is_directory and log_name.startswith("tls") and log_name.endswith("log"):
            start = time.time()
            print("process", event.src_path)
            # data = log_to_df(logName=event.src_path)
            # self.write_results(data, result)
            print("process time:", time.time() - start)


class TLSLogCreateHandler(FileSystemEventHandler):
    def __init__(self) -> None:
        super().__init__()
        # self.rfc = RandomForest.RandomForestPredictor(model_path, scaler_path)

    def on_created(self, event):
        # print("creat:", event)
        return
        # log_name = os.path.basename(event.src_path)
        # if not event.is_directory and log_name.startswith("tls") and log_name.endswith("log"):
        #     start = time.time()
        #     print("process", event.src_path)
        #     data = parseTLS.log_to_df(logName=event.src_path)
        #     result = self.rfc.predict(data)
        #     # self.write_results(data, result)
        #     print("process time:", time.time() - start)
        # elif log_name.endswith("log"):
        #     os.remove(event.src_path)

    def on_moved(self, event):
        print(event)
        log_name = os.path.basename(event.dest_path)
        if not event.is_directory and log_name.startswith("tls") and log_name.endswith("log"):
            print("===process", event.dest_path)
            # return
            start = time.time()
            # data = parseTLS.log_to_df(logName=event.dest_path)
            # print("lines:", data.shape[0])
            # result = self.rfc.predict(data)
            # self.write_results(data, result)
            os.system(f"python time_test.py {event.dest_path}")
            print("process time:", round(time.time() - start, 2), end="\n")

    def write_results(self, data: pd.DataFrame, result):
        data = data[INFO_COLUMN].copy()
        res_column = result.tolist()
        data['malware'] = res_column
        data.to_csv(csv_path, mode='a', header=False, index=False)
        pass


if __name__ == "__main__":
    file_path = '/home/fsc/liujy/file_watch/tls.log'
    dir_path = '/home/fsc/liujy/file_watch/pcap'

    # event_handler = MyHandler()
    event_handler = TLSLogCreateHandler()

    observer = Observer()
    observer.schedule(event_handler, dir_path, recursive=False)
    observer.start()
    print('start watching...')
    try:
        while True:
            pass
    except KeyboardInterrupt:
        observer.stop()
    observer.join()
