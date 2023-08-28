
import os
import sys
sys.path.append(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))

import dataproc as dp
from predictor import *
from pandarallel import pandarallel
pandarallel.initialize(nb_workers=2, verbose=0)


class FiveTuplePredictor(Predictor):
    def __init__(self) -> None:
        # work_dir = os.path.abspath(os.path.dirname(__file__))
        # os.chdir(work_dir)
        folder = os.path.dirname(os.path.abspath(__file__))
        super().__init__(type='fivetuple', folder=folder)

    def df_process(self, zeek_df: pd.DataFrame) -> Tuple[pd.DataFrame, pd.DataFrame]:
        zeek_df.replace(['(empty)', '-'], np.nan, inplace=True)
        zeek_df = dp.rm_Unestablished_ssl(zeek_df)
        if zeek_df.empty:
            return pd.DataFrame(), pd.DataFrame()

        extra_fea_df = pd.DataFrame(columns=extra_fea_column)
        zeek_df = pd.concat([zeek_df, extra_fea_df], axis=1)
        zeek_df = zeek_df.fillna(0)

        zeek_df = zeek_df.apply(dp.getCipherSuits, axis=1)
        zeek_df = zeek_df.apply(dp.getExtensions, axis=1)
        zeek_df = zeek_df.parallel_apply(dp.getPacketLenFea, axis=1)

        # 协议版本映射766:ssl1.0,767:2,768:3,769:tls1.0,770:tls1,771:tls2,772:tls3
        zeek_df['client_version'] = zeek_df['client_version'].map(dp.getClientVersion)
        zeek_df['server_version'] = zeek_df[['server_version', 'server_supported_version']].apply(dp.getServerVersion, axis=1)
        zeek_df['srcP'] = zeek_df['id.orig_p'].map(dp.getClientPortFea)
        zeek_df['dstP'] = zeek_df['id.resp_p'].map(dp.getServerPortFea)
        zeek_df['selfSigned'] = zeek_df[['issuer', 'subject']].apply(dp.isSelfSigned, axis=1)
        zeek_df['certValid'] = zeek_df[['ts', 'not_before', 'not_after']].apply(dp.isCertValid, axis=1)

        df = zeek_df[FIVE_TUPLE_FEATURE]
        info_df = zeek_df[INFO_COLUMN]
        return df, info_df

    def process_results(self, task_id: str) -> list:
        results = self.predict_results
        info_df = self.info_df
        flow_id = 0
        pcap_results = []
        for result, (_, info) in zip(results, info_df.iterrows()):
            predict_log = {}
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
            label = np.argmax(result)
            predict_log['classification'] = 'black' if label else 'white'

            flow_result = list(predict_log.values())
            # flow_result.insert(0, pcap_id)
            pcap_results.append(flow_result)
            flow_id += 1
        return pcap_results


if __name__ == "__main__":
    pcap_path = "/home/fsc/liujy/yahong/pcap/mta_tls.pcap"
    pre = FiveTuplePredictor()
    pre.predict_result(pcap_path)
