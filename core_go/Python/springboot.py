import random
import time
import joblib
import argparse
import pandas as pd
import module.feature_extraction as fe
import module.feature_extraction_taskid as fego
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
        model_path = "/core/python/model/model_bin_UEID.pkl"
    else:
        model_path = "E:/Code/web/backendspringboot3/core_go/Python/model/model_multi_UEID.pkl"
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


def detect_taskid(model_type, taskid):
    if model_type == 'bin':
        model_path = "/core/python/model/model_bin_UEID.pkl"
    else:
        model_path = "E:/Code/web/backendspringboot3/core_go/Python/model/model_multi_UEID.pkl"
    model = joblib.load(model_path)

    random.seed(27)
    X = []
    y = []
    taskid_params = {'taskid': taskid}
    select = "SELECT status FROM SCTP.Task WHERE taskId = %(taskid)s"
    task_status = 0

    while True:
        task_status = client.execute(select, taskid_params)[0][0]
        if task_status !=3:
            time.sleep(1)
        elif task_status == 3:
            tasking = "ALTER TABLE SCTP.Task UPDATE status = 4 WHERE taskId = %(taskid)s"
            client.execute(tasking, taskid_params)
            flow = "SELECT FlowId FROM SCTP.UEFlow WHERE TaskID = %(taskid)s AND StatusFlow = 0"
            flow_id = client.execute(flow, taskid_params)
            for id in flow_id:
                packet = "SELECT * FROM SCTP.Packet WHERE FlowUEID = %(flowid)s ORDER BY ArriveTime"
                packet_params = {'flowid': id}
                result = client.execute(packet, packet_params)
                X = []
                for row in result:
                    dirseq = row[10]
                    if dirseq == 1:
                        dirseq = 1
                    else:
                        dirseq = -1
                    X.append([row[2],row[3],row[5],dirseq,1])
                df = pd.DataFrame(X, columns=["RAN-UE-NGAP-ID", "Length", "Time", "DirSeq", "Label"])
                df['Time'] = df['Time'].astype('int64')/10**9
                feature, label = fe.feature_extract(df)
                y_predict = model.predict([feature])[0]
                if y_predict == 0:
                    predict_code = 100
                else:
                    predict_code = 200
                update_flow_query = 'ALTER TABLE SCTP.UEFlow UPDATE StatusFlow = %(ypredict)s WHERE FlowId = %(flowid)s'
                result_params = {
                    'ypredict': predict_code,
                    'flowid': id
                }
                client.execute(update_flow_query, result_params)
            break
        
    normal_query = "SELECT COUNT(*) FROM SCTP.UEFlow WHERE TaskID = %(taskid)s AND StatusFlow == 100"
    abnormal_query = "SELECT COUNT(*) FROM SCTP.UEFlow WHERE TaskID = %(taskid)s AND StatusFlow == 200"
    normal = client.execute(normal_query, taskid_params)[0][0]
    abnormal = client.execute(abnormal_query, taskid_params)[0][0]
    update_task_query = 'ALTER TABLE SCTP.Task UPDATE normal = %(normal)s, abnormal = %(abnormal)s, total = %(total)s WHERE taskId = %(taskid)s'
    
    final_params = {
        'normal': normal,
        'abnormal': abnormal,
        'total': normal + abnormal,
        'taskid': taskid
    }

    client.execute(update_task_query, final_params)
    client.disconnect()


def main(parser):
    args = parser.parse_args()
    # detect(args.model_type, args.file_path, args.taskid)
    detect_taskid(args.model_type, args.taskid)


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('--file_path', required=True, type=str)
    # parser.add_argument('--file_path',default=".\dataset\mix_dataset_test.csv",type=str)
    parser.add_argument('--model_type', default='multi', type=str)
    parser.add_argument('--taskid', required=True, type=str)
    main(parser)
