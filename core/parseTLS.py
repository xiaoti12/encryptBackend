from zat.log_to_dataframe import LogToDataFrame
import pandas as pd
import numpy as np
import dataproc as dp
from FeatureConstant import *
from pandarallel import pandarallel

pandarallel.initialize(nb_workers=4, verbose=0)


def log_to_df(zeek_df: pd.DataFrame = None, logName: str = None) -> pd.DataFrame:
    if logName:
        log_to_df = LogToDataFrame()
        zeek_df = log_to_df.create_dataframe(logName, ts_index=False, aggressive_category=False)

    if zeek_df.empty:
        print("===EMPTY DATA===")
        return pd.DataFrame()
    zeek_df.replace(['(empty)'], np.nan, inplace=True)
    zeek_df['duration'] = zeek_df['duration'].astype('timedelta64[s]')
    zeek_df = dp.rm_Unestablished_ssl(zeek_df)
    zeek_df = pd.concat([zeek_df, extra_fea_df], axis=1)
    zeek_df = zeek_df.fillna(0)

    # zeek_df = zeek_df.apply(dp.getCipherSuits, axis=1)
    # zeek_df = zeek_df.apply(dp.getExtensions, axis=1)

    zeek_df = zeek_df.parallel_apply(dp.getCipherSuits, axis=1)
    zeek_df = zeek_df.parallel_apply(dp.getExtensions, axis=1)
    zeek_df = zeek_df.parallel_apply(dp.getBytePacketStatics, axis=1)
    # zeek_df = zeek_df.parallel_apply(dp.getPacketLenFea, axis=1)

    # 协议版本映射766:ssl1.0,767:2,768:3,769:tls1.0,770:tls1,771:tls2,772:tls3
    zeek_df['client_version'] = zeek_df['client_version'].map(dp.getClientVersion)
    zeek_df['server_version'] = zeek_df[['server_version', 'server_supported_version']].apply(dp.getServerVersion, axis=1)
    zeek_df['srcP'] = zeek_df['id.orig_p'].map(dp.getClientPortFea)
    zeek_df['dstP'] = zeek_df['id.resp_p'].map(dp.getServerPortFea)
    zeek_df['selfSigned'] = zeek_df[['issuer', 'subject']].apply(dp.isSelfSigned, axis=1)
    zeek_df['certValid'] = zeek_df[['ts', 'not_before', 'not_after']].apply(dp.isCertValid, axis=1)

    df = zeek_df[INFO_COLUMN + FEATURE]
    # csvName="/home/fsc/liujy/file_watch/test.csv"
    # df.to_csv(csvName,index=False)
    df = df.fillna(0)
    return df

def process_results(info:pd.DataFrame,results,task_id)->list:
    pcap_results=[]
    flow_id=0
    for result, (_, info) in zip(results, info.iterrows()):
        predict_log={}
        predict_log['task_id'] = task_id
        predict_log['flow_id'] = str(flow_id)
        # predict_log['start_time'] = info['ts'].strftime('%Y-%m-%d %X') # string type
        predict_log['start_time'] = info['ts'] # 
        predict_log['ip_src'] = info['id.orig_h']
        predict_log['ip_dst'] = info['id.resp_h']
        predict_log['port_src'] = info['id.orig_p']
        predict_log['port_dst'] = info['id.resp_p']

        issuer = info['issuer']
        issuer_list = issuer.split(',') if issuer else []
        cn = [ele for ele in issuer_list if 'CN=' in ele]
        cn = cn[0].replace('CN=', '') if cn else None

        predict_log['issuer'] = issuer if issuer else None
        predict_log['common_name'] = cn
        predict_log['validity'] = 1 if not info['selfSigned'] else 0

        predict_log['white_proba'] = round(result[0], 5)
        predict_log['black_proba'] = round(result[1], 5)
        label = 1 if result[0]<result[1] else 0
        predict_log['classification'] = 'black' if label else 'white'

        flow_result = list(predict_log.values())
        # flow_result.insert(0, pcap_id)
        pcap_results.append(flow_result)
        flow_id += 1
    return pcap_results

def get_task_results(results):
    total=len(results)
    labels = [np.argmax(result) for result in results]
    abnormal = sum(labels)

    return [total - abnormal, abnormal, total]