import pandas as pd
from datetime import datetime
import numpy as np
import statistics as stats
from collections import Counter
from datetime import datetime


def rm_Unestablished_ssl(df):  # 过滤未成功建立ssl连接的数据
    df = df.loc[(df['ssl_established'] == 'T') | (df['ssl_established'] == 1), :]
    df = df.drop(['ssl_established'], axis=1)
    df.reset_index(drop=True, inplace=True)
    return df


def getCipherSuits(series):  # 提取客户端握手特征
    cs_List = str(series['client_ciphers']).split(',')
    for cs in cs_List:
        try:
            series['cs_' + str(cs)] += 1
        except:
            series['cs_unknown'] += 1
    return series


def getExtensions(series):  # 提取客户端扩展特征
    ext_List = str(series['ssl_client_exts']).split(',')
    for ext in ext_List:
        try:
            series['ext_' + str(ext)] += 1
        except:
            series['ext_unassigned'] += 1
    return series


def getPacketLenFea(series, state_size=10):  # 包长度转移矩阵
    state_matrix = [0] * state_size * state_size
    if (series['or_spl'] == '(empty)'):
        return series
    else:
        spl = str(series['or_spl']).split(',')
        spl = list(map(int, spl))
        series['packet_num'] = len(spl)
        up_packet = 0
        down_packet = 0
        up_byte = 0
        down_byte = 0

        for i in spl:
            if i < 0:
                up_packet += 1
                up_byte += (-1) * i
            else:
                down_packet += 1
                down_byte += i
        series['up_packet'] = up_packet
        series['up_byte'] = up_byte
        series['down_packet'] = down_packet
        series['down_byte'] = down_byte
        series['up_down_packet_ratio'] = (up_packet + 0.01) / down_packet
        series['up_down_byte_ratio'] = (up_byte + 0.01) / down_byte

        if (len(spl) > 50):
            spl = spl[0:50]
        # packet长度>1500的，一律放到最后一个state，防止后续计算转移矩阵越界
        for i in spl:
            if (int(i) >= 1500):
                i = 1499

        for i in range(len(spl) - 1):
            # 映射到[0,150),[150,300)....
            packet_i = int(abs(int(spl[i])) / 150)
            packet_j = int(abs(int(spl[i + 1])) / 150)
            # 计算转移矩阵
            state_matrix[packet_i * state_size + packet_j] += 1

        for i in range(state_size):
            all_count = sum(state_matrix[i * state_size:(i + 1) * state_size])
            if all_count <= 0:
                continue
            for j in range(state_size):
                state_matrix[i * state_size + j] /= all_count

        # matrix赋值给series
        for i in range(len(state_matrix)):
            series['mkv_' + str(i)] = state_matrix[i]
    return series


def getClientPortFea(x):
    if x >= 49152 and x <= 65535:
        x = 0
    else:
        x = 1
    return x


def getServerPortFea(x):
    if x == 443 or x == 465 or x == 563 or x == 636 or x == 853 or x == 989 or x == 990 or x == 992 or x == 993 or x == 995:
        x = 0
    else:
        x = 1
    return x


def getServerVersion(series):
    if series['server_version'] == 771 and series['server_supported_version'] == 772:
        return 7
    elif series['server_version'] == 766:
        return 1
    elif series['server_version'] == 767:
        return 2
    elif series['server_version'] == 768:
        return 3
    elif series['server_version'] == 769:
        return 4
    elif series['server_version'] == 770:
        return 5
    elif series['server_version'] == 771:
        return 6
    else:
        return 0


def getClientVersion(x):
    if x == 766:
        return 1
    elif x == 767:
        return 2
    elif x == 768:
        return 3
    elif x == 769:
        return 4
    elif x == 770:
        return 5
    elif x == 771:
        return 6
    elif x == 772:
        return 7
    else:
        return 0


def isSelfSigned(series):  # -1代表缺失证书
    if series['subject'] != 0 and series['issuer'] != 0:
        if series['subject'] == series['issuer']:
            return 1
        else:
            return 0
    else:
        return -1


def isCertValid(series):  # 1代表有效，0代表无效
    if series['not_before'] == 0 or series['not_after'] == 0:
        return -1
    # now = datetime.strptime(series['ts'].split('.')[0], "%Y-%m-%d %H:%M:%S")
    if isinstance(series['not_before'], pd.Timestamp):
        before = series['not_before']
        after = series['not_after']
    elif isinstance(series['not_before'], str):
        before = datetime.strptime(series['not_before'], "%Y-%m-%d %H:%M:%S")
        after = datetime.strptime(series['not_after'], "%Y-%m-%d %H:%M:%S")
    now = series['ts']
    # before = series['not_before']
    # after = series['not_after']
    if (now > before and now < after):
        return 1
    else:
        return 0


def getCertValidTime(x):
 # 将每个有效时长转换为list[int]，单位为天
    if x == 0:
        return 0
    x = valueToList(x, lambda i: int(float(i) / 86400))
    return int(sum(x) / len(x))


def calFourTupleId(df):
    df['four_tuple_id'] = -1
    tuple_dict = {}  # structure: {four_tuple_id:[(start_time1,id1),(start_time2,id2),...],...}
    id = 0
    for index, row in df.iterrows():
        row_id = -1
        if type(row['ts']) == str:
            start_time = float(row['ts'])
        else:
            start_time = row['ts'].timestamp()
        four_tuple = f"{row['id.orig_h']} {row['id.resp_h']} {row['id.resp_p']}"
        if four_tuple not in tuple_dict:  # 新的四元组
            tuple_dict[four_tuple] = [(start_time, id)]
            row_id = id
            id += 1
        else:  # 已有四元组
            for time_id in tuple_dict[four_tuple]:
                if start_time - time_id[0] < 360:  # 同一四元组同一时间片
                    row_id = time_id[1]
                    break
            if row_id < 0:  # 同一四元组新的时间片:
                tuple_dict[four_tuple].append((start_time, id))
                row_id = id
                id += 1
        df.at[index, 'four_tuple_id'] = row_id
    return df


def valueToList(x, func):
    x = x.split(',')
    x = list(map(func, x))
    return x


def getListMode(l):
    count = Counter(l)
    mode = sorted(count.items(), key=lambda x: x[1], reverse=True)[0]
    if mode[1] > 1:
        return mode[0]
    else:
        return -1


def getMean(x):
    x = valueToList(x, float)
    return round(np.mean(x), 5)


def getMin(x):
    x = valueToList(x, float)
    return round(min(x), 5)


def getMax(x):
    x = valueToList(x, float)
    return round(max(x), 5)


def getMedian(x):
    x = valueToList(x, lambda x: int(float(x)))
    return stats.median(x)


def getMode(x):
    x = valueToList(x, lambda x: int(float(x)))
    return getListMode(x)


def getShortDuration(x):
    x = valueToList(x, float)
    short_dur = [i for i in x if i <= 60]
    return len(short_dur)


def getIntervalMean(x):
    x = valueToList(x, lambda x: datetime.strptime(x.split('.')[0], '%Y-%m-%d %H:%M:%S').timestamp())
    interval = []
    if len(x) == 1:
        return 0
    for i in range(len(x) - 1):
        interval.append(x[i + 1] - x[i])
    return round(np.mean(interval), 5)


def getSniRatio(series):
    sni = valueToList(series['sni'], str)
    sni = [i for i in sni if i != '0']
    conn_num = series['total_conn']
    return round(len(sni) / conn_num, 5)


def getSniNum(x):
    sni = valueToList(x, str)
    sni = [i for i in sni if i != '0']
    return len(set(x.split(',')))


def getSSLVersionRatio(series, target_ver):
    version = valueToList(series['tls_version'], int)
    conn_num = series['total_conn']
    ratio = version.count(target_ver) / conn_num
    return round(ratio, 5)


def getNonStdPortNum(x):
    standard_port = [443, 80, 8080, 9001, 9010, 9050, 9150, 12443]
    port = valueToList(x, int)
    non_std_port = [
        x for x in port if (x not in standard_port and x > 1023)
    ]
    return len(non_std_port)


def getNonStdMode(x):
    standard_port = [443, 80, 8080, 9001, 9010, 9050, 9150, 12443]
    port = valueToList(x, int)
    non_std_port = [
        x for x in port if (x not in standard_port and x > 1023)
    ]
    if non_std_port:
        return getListMode(non_std_port)
    else:
        return -1
