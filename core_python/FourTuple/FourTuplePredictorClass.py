import os
import sys
sys.path.append(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))
from predictor import *
import dataproc as dp


class FourTuplePredictor(Predictor):
    def __init__(self) -> None:
        # work_dir = os.path.abspath(os.path.dirname(__file__))
        # os.chdir(work_dir)
        folder = os.path.dirname(os.path.abspath(__file__))
        super().__init__(type='fourtuple', folder=folder)

    def df_process(self, zeek_df: pd.DataFrame) -> Tuple[pd.DataFrame, pd.DataFrame]:
        # 预处理
        zeek_df.replace(['(empty)', '-'], np.nan, inplace=True)
        zeek_df = dp.rm_Unestablished_ssl(zeek_df)
        if zeek_df.empty:
            return pd.DataFrame(), pd.DataFrame()
        zeek_df['selfSigned'] = zeek_df[['issuer', 'subject']].apply(dp.isSelfSigned, axis=1)
        zeek_df['total_bytes'] = zeek_df[['up_bytes', 'down_bytes']].apply(sum, axis=1)

        zeek_df.fillna(0, inplace=True)

        zeek_df['valid_time'] = zeek_df['valid_time'].apply(dp.getCertValidTime)  # 单位：天
        zeek_df['tls_version'] = zeek_df[['server_version', 'server_supported_version']].apply(dp.getServerVersion, axis=1)

        zeek_df = dp.calFourTupleId(zeek_df)
        zeek_df = zeek_df.groupby('four_tuple_id').agg(lambda x: ','.join(map(str, x)))

        fea_df = pd.DataFrame(columns=FOUR_TUPLE_FEATURE)
        info_df = zeek_df.copy()
        zeek_df = pd.concat([zeek_df, fea_df], axis=1)
        zeek_df.fillna(0, inplace=True)

        # 四元组标签和聚合
        zeek_df['duration_mean'] = zeek_df['duration'].apply(dp.getMean)
        zeek_df['duration_min'] = zeek_df['duration'].apply(dp.getMin)
        zeek_df['duration_max'] = zeek_df['duration'].apply(dp.getMax)
        zeek_df['short_dur'] = zeek_df['duration'].apply(dp.getShortDuration)
        zeek_df['interval_mean'] = zeek_df['ts'].apply(dp.getIntervalMean)
        zeek_df['total_conn'] = zeek_df['duration'].apply(lambda x: len(x.split(',')))

        for column in ['total_bytes', 'up_bytes', 'down_bytes', 'up_pkts', 'down_pkts']:
            zeek_df[f'{column}_mean'] = zeek_df[column].apply(dp.getMean)
            zeek_df[f'{column}_median'] = zeek_df[column].apply(dp.getMedian)
            zeek_df[f'{column}_mode'] = zeek_df[column].apply(dp.getMode)

        zeek_df['sni_ratio'] = zeek_df[['sni', 'total_conn']].apply(dp.getSniRatio, axis=1)
        zeek_df['sni_server_num'] = zeek_df['sni'].apply(lambda x: len(set(x.split(','))))

        for i, tls_ver_column in enumerate(tls_column[-6:]):
            # sslv2 -> 2
            zeek_df[tls_ver_column] = zeek_df[['tls_version', 'total_conn']].apply(dp.getSSLVersionRatio, args=(i + 2,), axis=1)

        zeek_df['uniq_cert_num'] = zeek_df['cert_serial'].apply(lambda x: len(set(x.split(','))))
        zeek_df['chain_len_mean'] = zeek_df['cert_num'].apply(dp.getMean)
        zeek_df['cert_valid_mean'] = zeek_df['valid_time'].apply(dp.getMean)
        zeek_df['self_signed_num'] = zeek_df['selfSigned'].apply(lambda x: x.split(',').count('1'))
        zeek_df['cert_len_mean'] = zeek_df['cert_len'].apply(dp.getMean)

        zeek_df['non_std_dstp_num'] = zeek_df['id.resp_p'].apply(dp.getNonStdPortNum)
        zeek_df['common_non_std_dstp'] = zeek_df['id.resp_p'].apply(dp.getNonStdMode)

        df = zeek_df[FOUR_TUPLE_FEATURE]
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
            predict_log['start_time'] = info['ts'].split(',')[0]
            predict_log['ip_src'] = info['id.orig_h'].split(',')[0]
            predict_log['ip_dst'] = info['id.resp_h'].split(',')[0]
            predict_log['port_src'] = None
            predict_log['port_dst'] = info['id.resp_p'].split(',')[0]

            issuer = info['issuer']
            issuer_list = issuer.split(',') if issuer else []
            cn = [ele for ele in issuer_list if 'CN=' in ele]
            cn = cn[0].replace('CN=', '') if cn else None

            predict_log['issuer'] = ','.join(set(issuer_list)) if issuer else None
            predict_log['common_name'] = cn
            predict_log['validity'] = None

            predict_log['white_proba'] = round(result[0], 5)
            predict_log['black_proba'] = round(result[1], 5)
            label = np.argmax(result)
            predict_log['classification'] = 'black' if label else 'white'
            # if label:
            #     malware_num += 1
            #     predict_log['classification'] = 'black'
            # else:
            #     predict_log['classification'] = 'white'

            flow_result = list(predict_log.values())
            pcap_results.append(flow_result)
            flow_id += 1
        return pcap_results


if __name__ == "__main__":
    pcap_path = "/home/fsc/liujy/yahong/pcap/mta_tls.pcap"
    pre = FourTuplePredictor()
    pre.predict_result(pcap_path)
