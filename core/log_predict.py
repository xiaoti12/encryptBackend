import parseTLS
import RandomForest
from FeatureConstant import INFO_COLUMN
import os
import clickhouse_connect
from clickhouse_connect.driver.client import Client
import argparse

# log = "tls.log"
work_dir = os.path.dirname(os.path.abspath(__file__))
model_path = os.path.join(work_dir, "train", "rfc_nomkv.joblib")
scaler_path = os.path.join(work_dir, "train", "scaler_nomkv.joblib")
# csv_path = "/home/fsc/liujy/realtime/results.csv"


def write_results(data, result):
    data = data[INFO_COLUMN].copy()
    res_column = result.tolist()
    # res_column = [round(i[1], 3) for i in res_column]
    data['malware'] = res_column
    data.to_csv(None, mode='a', header=False, index=False)


def write_database(c: Client, results_pred, results_db, task_id):
    task_results = parseTLS.get_task_results(results_pred)
    olddata = c.query(f"SELECT `normal`,`abnormal`,`total` from Task WHERE `taskId`='{task_id}'")
    olddata = olddata.result_set[0]
    c.command(
        f"ALTER TABLE Task UPDATE `normal`={task_results[0]+olddata[0]},`abnormal`={task_results[1]+olddata[1]},`total`={task_results[2]+olddata[2]} WHERE `taskId`='{task_id}'")
    c.insert('TLSFlow', results_db)


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("--log", type=str, help="path of log")
    parser.add_argument("--taskid", type=str, help="task id")

    args = parser.parse_args()
    log = args.log
    task_id = args.taskid

    client = clickhouse_connect.get_client(host='localhost', password='password', database='TLS')
    rfc = RandomForest.RandomForestPredictor(model_path, scaler_path)

    data = parseTLS.log_to_df(logName=log)
    # print("lines:", data.shape[0])
    results_pred = rfc.predict(data)
    results_db = parseTLS.process_results(data, results_pred, task_id=task_id)
    write_database(client, results_pred, results_db, task_id)

    # write_results(data, results)
