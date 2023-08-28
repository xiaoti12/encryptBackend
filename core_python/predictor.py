import joblib
import os
import sqlite3
import pandas as pd
import numpy as np
from typing import Tuple
from json import dumps
# work_dir = os.path.abspath(os.path.dirname(__file__))
# os.chdir(work_dir)
import subprocess
import fcntl


class Predictor():
    def __init__(self, type: str, folder: str) -> None:
        if type not in ['fourtuple', 'fivetuple']:
            raise ValueError("type must be 'fourtuple' or 'fivetuple'")
        # with open(f"{type}_rfc.joblib", 'rb') as f:
        with open(os.path.join(folder, f"{type}_rfc.joblib"), 'rb') as f:
            self.model = joblib.load(f)
        with open(os.path.join(folder, f"{type}_scaler.joblib"), 'rb') as f:
            self.scaler = joblib.load(f)
        self.type = type
        self.predict_results: np.ndarray = None
        self.info_df: pd.DataFrame = None
        # else:
        #     self.zeek_path = "./FiveTuple/fivetuple.zeek"
        #     self.conn = sqlite3.connect("./FiveTuple/fivetuple.sqlite")

        self.zeek_path = os.path.join(os.path.join(folder, f'{type}.zeek'))
        self.conn = sqlite3.connect(os.path.join(folder, f"{type}.sqlite"))
        self.sqlite_file=os.path.join(folder, f"{type}.sqlite")

        self.feature = None

    def zeek_process(self, pcap_path: str) -> None:
        c = self.conn.cursor()
        c.execute(f"delete from {self.type}")
        self.conn.commit()

        sqlite_f=open(self.sqlite_file,'r')
        fcntl.flock(sqlite_f,fcntl.LOCK_EX)
        cmd = f'zeek -C -r {pcap_path} {self.zeek_path}'
        # os.system(cmd)  # 这里如果用os.popen()会异步执行，还未写入数据库时后续命令就开始读
        result = subprocess.run(cmd, shell=True, capture_output=True, text=True)
        # 获取 stdout
        print(result.stdout,result.stderr)
        fcntl.flock(sqlite_f,fcntl.LOCK_UN)
        sqlite_f.close()

    def read(self,pcap_path:str) -> pd.DataFrame:
        column_types = {
            'ts': 'datetime64[s]',
            'not_before': 'datetime64[s]',
            'not_after': 'datetime64[s]',
        }
        query = f"SELECT * FROM {self.type} WHERE pcap='{pcap_path}'"
        df = pd.read_sql_query(query, self.conn, dtype=column_types)
        # print(f"====={pcap_path} columns:{df.shape[0]}======")

        return df

    def df_process(self, df) -> Tuple[pd.DataFrame, pd.DataFrame]:
        return df, df

    def get_task_results(self) -> list:
        results = self.predict_results
        total = len(results)
        labels = [np.argmax(result) for result in results]
        abnormal = sum(labels)

        return [total - abnormal, abnormal, total]

    def predict_result(self, pcap_path: str) -> None:
        self.zeek_process(pcap_path)
        zeek_df = self.read(pcap_path)
        feature_df, info_df = self.df_process(zeek_df)

        if feature_df.empty:
            print("no tls-established flow in pcap")
            return

        feature = self.scaler.transform(feature_df.values)
        results = self.model.predict_proba(feature)
        self.predict_results = results
        self.info_df = info_df

    def print_results(self, results: np.ndarray, info_df: pd.DataFrame) -> None:
        pass

    # def database_record(self, results: np.ndarray, info_df: pd.DataFrame) -> list:
    #     info_df['label'] = results
    #     info_df['ip_2_tuple'] = info_df['id.orig_h'] + '-' + info_df['id.resp_h']

    #     ip_2_tuple = list(set(info_df['ip_2_tuple']))
    #     port_dst = list(set(info_df['id.resp_p']))
    #     malware_num = sum(info_df['label'])
    #     total_num = info_df.shape[0]
    #     normal_num = total_num - malware_num

    #     return [ip_2_tuple, port_dst, malware_num, normal_num, total_num]


# 四元组特征
conn_column = ['duration_mean', 'duration_min', 'duration_max', 'short_dur', 'interval_mean', 'total_conn']
data_column = ['total_bytes', 'up_bytes', 'down_bytes', 'up_pkts', 'down_pkts']
data_column = [f'{i}_{j}' for i in data_column for j in ['mean', 'median', 'mode']]
tls_column = ['sni_ratio', 'sni_server_num', 'sslv2_ratio', 'sslv3_ratio', 'tls10_ratio', 'tls11_ratio', 'tls12_ratio', 'tls13_ratio']
cert_column = ['uniq_cert_num', 'chain_len_mean', 'cert_valid_mean', 'self_signed_num', 'cert_len_mean']
port_column = ['non_std_dstp_num', 'common_non_std_dstp']

FOUR_TUPLE_FEATURE = conn_column + data_column + tls_column + cert_column + port_column


# 五元组特征
cs_column = ['cs_0', 'cs_1', 'cs_2', 'cs_3', 'cs_4', 'cs_5', 'cs_6',
             'cs_7', 'cs_8', 'cs_9', 'cs_10', 'cs_11', 'cs_12', 'cs_13',
             'cs_14', 'cs_15', 'cs_16', 'cs_17', 'cs_18', 'cs_19', 'cs_20',
             'cs_21', 'cs_22', 'cs_23', 'cs_24', 'cs_25', 'cs_26', 'cs_27',
             'cs_30', 'cs_31', 'cs_32', 'cs_33', 'cs_34', 'cs_35', 'cs_36',
             'cs_37', 'cs_38', 'cs_39', 'cs_40', 'cs_41', 'cs_42', 'cs_43',
             'cs_44', 'cs_45', 'cs_46', 'cs_47', 'cs_48', 'cs_49', 'cs_50',
             'cs_51', 'cs_52', 'cs_53', 'cs_54', 'cs_55', 'cs_56', 'cs_57',
             'cs_58', 'cs_59', 'cs_60', 'cs_61', 'cs_62', 'cs_63', 'cs_64',
             'cs_65', 'cs_66', 'cs_67', 'cs_68', 'cs_69', 'cs_70', 'cs_103',
             'cs_104', 'cs_105', 'cs_106', 'cs_107', 'cs_108', 'cs_109', 'cs_132',
             'cs_133', 'cs_134', 'cs_135', 'cs_136', 'cs_137', 'cs_138', 'cs_139',
             'cs_140', 'cs_141', 'cs_142', 'cs_143', 'cs_144', 'cs_145', 'cs_146',
             'cs_147', 'cs_148', 'cs_149', 'cs_150', 'cs_151', 'cs_152', 'cs_153',
             'cs_154', 'cs_155', 'cs_156', 'cs_157', 'cs_158', 'cs_159', 'cs_160',
             'cs_161', 'cs_162', 'cs_163', 'cs_164', 'cs_165', 'cs_166', 'cs_167',
             'cs_168', 'cs_169', 'cs_170', 'cs_171', 'cs_172', 'cs_173', 'cs_174',
             'cs_175', 'cs_176', 'cs_177', 'cs_178', 'cs_179', 'cs_180', 'cs_181',
             'cs_182', 'cs_183', 'cs_184', 'cs_185', 'cs_186', 'cs_187', 'cs_188',
             'cs_189', 'cs_190', 'cs_191', 'cs_192', 'cs_193', 'cs_194', 'cs_195',
             'cs_196', 'cs_197', 'cs_198', 'cs_199', 'cs_255', 'cs_4865', 'cs_4866',
             'cs_4867', 'cs_4868', 'cs_4869', 'cs_22016', 'cs_49153', 'cs_49154',
             'cs_49155', 'cs_49156', 'cs_49157', 'cs_49158', 'cs_49159', 'cs_49160',
             'cs_49161', 'cs_49162', 'cs_49163', 'cs_49164', 'cs_49165', 'cs_49166',
             'cs_49167', 'cs_49168', 'cs_49169', 'cs_49170', 'cs_49171', 'cs_49172',
             'cs_49173', 'cs_49174', 'cs_49175', 'cs_49176', 'cs_49177', 'cs_49178',
             'cs_49179', 'cs_49180', 'cs_49181', 'cs_49182', 'cs_49183', 'cs_49184',
             'cs_49185', 'cs_49186', 'cs_49187', 'cs_49188', 'cs_49189', 'cs_49190',
             'cs_49191', 'cs_49192', 'cs_49193', 'cs_49194', 'cs_49195', 'cs_49196',
             'cs_49197', 'cs_49198', 'cs_49199', 'cs_49200', 'cs_49201', 'cs_49202',
             'cs_49203', 'cs_49204', 'cs_49205', 'cs_49206', 'cs_49207', 'cs_49208',
             'cs_49209', 'cs_49210', 'cs_49211', 'cs_49212', 'cs_49213', 'cs_49214',
             'cs_49215', 'cs_49216', 'cs_49217', 'cs_49218', 'cs_49219', 'cs_49220',
             'cs_49221', 'cs_49222', 'cs_49223', 'cs_49224', 'cs_49225', 'cs_49226',
             'cs_49227', 'cs_49228', 'cs_49229', 'cs_49230', 'cs_49231', 'cs_49232',
             'cs_49233', 'cs_49234', 'cs_49235', 'cs_49236', 'cs_49237', 'cs_49238',
             'cs_49239', 'cs_49240', 'cs_49241', 'cs_49242', 'cs_49243', 'cs_49244',
             'cs_49245', 'cs_49246', 'cs_49247', 'cs_49248', 'cs_49249', 'cs_49250',
             'cs_49251', 'cs_49252', 'cs_49253', 'cs_49254', 'cs_49255', 'cs_49256',
             'cs_49257', 'cs_49258', 'cs_49259', 'cs_49260', 'cs_49261', 'cs_49262',
             'cs_49263', 'cs_49264', 'cs_49265', 'cs_49266', 'cs_49267', 'cs_49268',
             'cs_49269', 'cs_49270', 'cs_49271', 'cs_49272', 'cs_49273', 'cs_49274',
             'cs_49275', 'cs_49276', 'cs_49277', 'cs_49278', 'cs_49279', 'cs_49280',
             'cs_49281', 'cs_49282', 'cs_49283', 'cs_49284', 'cs_49285', 'cs_49286',
             'cs_49287', 'cs_49288', 'cs_49289', 'cs_49290', 'cs_49291', 'cs_49292',
             'cs_49293', 'cs_49294', 'cs_49295', 'cs_49296', 'cs_49297', 'cs_49298',
             'cs_49299', 'cs_49300', 'cs_49301', 'cs_49302', 'cs_49303', 'cs_49304',
             'cs_49305', 'cs_49306', 'cs_49307', 'cs_49308', 'cs_49309', 'cs_49310',
             'cs_49311', 'cs_49312', 'cs_49313', 'cs_49314', 'cs_49315', 'cs_49316',
             'cs_49317', 'cs_49318', 'cs_49319', 'cs_49320', 'cs_49321', 'cs_49322',
             'cs_49323', 'cs_49324', 'cs_49325', 'cs_49326', 'cs_49327', 'cs_49328',
             'cs_49329', 'cs_49330', 'cs_49331', 'cs_49332', 'cs_49333', 'cs_49408',
             'cs_49409', 'cs_49410', 'cs_49411', 'cs_49412', 'cs_49413', 'cs_49414',
             'cs_52392', 'cs_52393', 'cs_52394', 'cs_52395', 'cs_52396', 'cs_52397',
             'cs_52398', 'cs_53249', 'cs_53250', 'cs_53251', 'cs_53253', 'cs_unknown']
ext_column = ['ext_0', 'ext_1',
              'ext_2', 'ext_3', 'ext_4', 'ext_5', 'ext_6', 'ext_7', 'ext_8',
              'ext_9', 'ext_10', 'ext_11', 'ext_12', 'ext_13', 'ext_14',
              'ext_15', 'ext_16', 'ext_17', 'ext_18', 'ext_19', 'ext_20',
              'ext_21', 'ext_22', 'ext_23', 'ext_24', 'ext_25', 'ext_26',
              'ext_27', 'ext_28', 'ext_29', 'ext_30', 'ext_31', 'ext_32',
              'ext_33', 'ext_34', 'ext_35', 'ext_36', 'ext_37', 'ext_38',
              'ext_39', 'ext_40', 'ext_41', 'ext_42', 'ext_43', 'ext_44',
              'ext_45', 'ext_46', 'ext_47', 'ext_48', 'ext_49', 'ext_50',
              'ext_51', 'ext_52', 'ext_53', 'ext_55', 'ext_56', 'ext_65281',
              'ext_unassigned']
mkv_column = ['mkv_' + str(i) for i in range(0, 100)]
stat_column = [
    'packet_num',  # flow中包总数
    'up_packet',  # flow中上行包数
    'up_byte',  # flow中上行字节数
    'down_packet',
    'down_byte',
    'up_down_packet_ratio',
    'up_down_byte_ratio',
    'srcP',  # 源ip是否x >= 49152 and x <= 65535
    'dstP',  # 目的ip是否为常见tls协议端口
    'selfSigned',  # 1：自签名证书；0：非自签名证书；-1：缺失证书
    'certValid',  # 1：有效，0：无效；-1：缺失证书
    # 'duration', # flow持续时间，单位s ,重放会破坏该特征
    'client_version',  # tls协议版本
    'server_version',
    'cert_num',  # 证书链中证书的数量
    'san_num',
    'ext_num',
]
extra_fea_column = ['packet_num', 'up_packet', 'up_byte', 'down_packet', 'down_byte', 'up_down_packet_ratio',
                    'up_down_byte_ratio', 'srcP', 'dstP', 'selfSigned', 'certValid', 'isBenign'] + cs_column + ext_column + mkv_column

fea_column = extra_fea_column + ['duration', 'client_version', 'server_version', 'cert_num', 'san_num', 'ext_num']   # 计算得到的特征

info_column = ['uid', 'id.orig_h', 'id.orig_p', 'id.resp_h', 'id.resp_p', 'sni']

# FIVE_TUPLE_FEATURE = cs_column + ext_column + mkv_column + stat_column
FIVE_TUPLE_FEATURE = stat_column + cs_column + mkv_column + ext_column

INFO_COLUMN = ['ts', 'id.orig_h', 'id.orig_p', 'id.resp_h', 'id.resp_p', 'issuer', 'selfSigned']

csv_column = ['ts', 'id.orig_h', 'id.orig_p', 'id.resp_h', 'id.resp_p', 'duration',
              'or_spl', 'up_bytes', 'down_bytes', 'up_pkts', 'down_pkts',
              'server_version', 'client_version', 'cipher', 'client_ciphers', 'sni', 'ssl_established',
              'ssl_client_exts', 'ssl_server_exts', 'client_supported_versions', 'server_supported_version',
              'issuer', 'subject', 'cert_num', 'not_before', 'not_after', 'valid_time', 'san_num', 'ext_num', 'cert_serial', 'cert_len']
