import pandas as pd
from datetime import datetime
# 过滤未成功建立ssl连接的数据


def rm_Unestablished_ssl(df):
    df = df.loc[(df['ssl_established'] == 'T') | (df['ssl_established'] == 1), :]
    df = df.drop(['ssl_established'], axis=1)
    df.reset_index(drop=True, inplace=True)
    return df
# 提取客户端握手特征


def getCipherSuits(series):
    # cipherSuitList = series['client_ciphers']
    # print(type(series['client_ciphers']))
    cs_List = str(series['client_ciphers']).split(',')
    for cs in cs_List:
        try:
            series['cs_' + str(cs)] += 1
        except:
            series['cs_unknown'] += 1
    return series
# 提取客户端扩展特征


def getExtensions(series):
    ext_List = str(series['ssl_client_exts']).split(',')
    for ext in ext_List:
        try:
            series['ext_' + str(ext)] += 1
        except:
            series['ext_unassigned'] += 1
    return series
# 包长度转移矩阵


def getBytePacketStatics(series):
    if (series['or_spl'] == '(empty)'):
        return series
    spl = str(series['or_spl']).split(',')
    spl = list(map(int, spl))
    series['packet_num'] = len(spl)
    up_packet = 0
    down_packet = 0
    up_byte = 0
    down_byte = 0

    up = [i for i in spl if i < 0]
    down = [i for i in spl if i > 0]

    up_packet = len(up)
    up_byte = (-1) * sum(up)
    down_packet = len(down)
    down_byte = sum(down)

    series['up_packet'] = up_packet
    series['up_byte'] = up_byte
    series['down_packet'] = down_packet
    series['down_byte'] = down_byte
    series['up_down_packet_ratio'] = (up_packet + 0.01) / down_packet
    series['up_down_byte_ratio'] = (up_byte + 0.01) / down_byte
    return series


def getPacketLenFea(series, state_size=10):
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
    if x in [443, 465, 563, 636, 853, 989, 990, 992, 993, 995]:
        x = 0
    else:
        x = 1
    return x


def getServerVersion(series):
    server_version = series['server_version']
    supported_version = series['server_supported_version']

    if server_version == 771 and supported_version == 772:  # tls1.3
        return 7
    elif server_version >= 766 and server_version <= 771:
        return server_version - 765
    else:
        return 0


def getClientVersion(x):
    if x >= 766 and x <= 772:
        return x - 765
    else:
        return 0


# -1代表缺失证书
def isSelfSigned(series):
    if series['subject'] != 0 and series['issuer'] != 0:
        if series['subject'] == series['issuer']:
            return 1
        else:
            return 0
    else:
        return -1
# 1代表有效，0代表无效


def isCertValid(series):
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


if __name__ == "__main__":
    data = pd.DataFrame({'or_spl': ['1,2,-3,-4,-5', '3,-4,5,-6,7']})
    data = data.apply(getBytePacketStatics, axis=1)
