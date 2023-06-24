import joblib
import argparse
import pandas as pd
import module.feature_extraction as fe
import module.read_pcap as rp
import os
from clickhouse_driver import Client

client = Client(host='10.3.242.84', port=9000, user='default', password='password')


def detect(model_type, file_path, taskid):
    file_path = "E:\\Code\\web\\backendspringboot3\\core\\upload\\" + file_path
    suffix = file_path.split(".")[-1]
    if suffix == "pcap" or suffix == "pcapng":
        a = rp.read_pcap(file_path)
    elif suffix == "csv":
        a = pd.read_csv(file_path)
    X, y = fe.get_dataset(a, client, taskid)
    if model_type == 'bin':
        model_path = "E:\\Code\\web\\backendspringboot3\\core\\model\\model_bin_UEID.pkl"
    else:
        model_path = "E:\\Code\\web\\backendspringboot3\\core\\model\\model_multi_UEID.pkl"
    model = joblib.load(model_path)
    if len(X) > 0:
        y_predict = model.predict(X)
        print("Total traffic:", len(X), " Abnormal traffic: ", sum(y_predict >= 1))
        query = 'ALTER TABLE SCTP.Task UPDATE normal = %(normal)s, abnormal = %(abnormal)s, total = %(total)s WHERE taskId = %(taskid)s'
        params = {
            'normal': sum(y_predict == 0),
            'abnormal': sum(y_predict >= 1),
            'total': len(X),
            'taskid': taskid
        }
        client.execute(query, params)
    else:
        print("All normal traffic")
        query = 'ALTER TABLE SCTP.Task UPDATE normal = %(normal)s, abnormal = %(abnormal)s, total = %(total)s WHERE taskId = %(taskid)s'
        params = {
            'normal': 0,
            'abnormal': 0,
            'total': 0,
            'taskid': taskid
        }
        client.execute(query, params)
    client.disconnect()


def main(parser):
    args = parser.parse_args()
    detect(args.model_type, args.file_path, args.taskid)


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('--file_path', required=True, type=str)
    # parser.add_argument('--file_path',default=".\dataset\mix_dataset_test.csv",type=str)
    parser.add_argument('--model_type', default='multi', type=str)
    parser.add_argument('--taskid', required=True, type=str)
    main(parser)
