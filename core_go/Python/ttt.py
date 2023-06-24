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

def detect_taskid(model_type, taskid):
    if model_type == 'bin':
        model_path = "python/model/model_bin_UEID.pkl"
    else:
        model_path = "E:/Code/web/backendspringboot3/core_go/Python/model/model_multi_UEID.pkl"
    model = joblib.load(model_path)

    random.seed(27)
    X = []
    y = []
    taskid_params = {'taskid': taskid}

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
        
    normal_query = "SELECT COUNT(*) FROM SCTP.UEFlow WHERE StatusFlow == 0"
    abnormal_query = "SELECT COUNT(*) FROM SCTP.UEFlow WHERE StatusFlow >= 1"
    normal = client.execute(normal_query)[0][0]
    abnormal = client.execute(abnormal_query)[0][0]
    
    client.disconnect()


def main(parser):
    args = parser.parse_args()
    # detect(args.model_type, args.file_path, args.taskid)
    detect_taskid(args.model_type, args.taskid)


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    # parser.add_argument('--file_path', required=True, type=str)
    parser.add_argument('--file_path',default=".\dataset\mix_dataset_test.csv",type=str)
    parser.add_argument('--model_type', default='multi', type=str)
    parser.add_argument('--taskid', default='2000', type=str)
    main(parser)
