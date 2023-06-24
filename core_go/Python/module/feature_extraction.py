# -*- coding: utf-8 -*-
import random

import numpy as np
import pandas as pd


def feature_extract(item):
    """
    item 是一个字典，包含 'Dirseq','Time','Length','Label'
    """
    label = (max(set(list(item['Label'])), key=list(item['Label']).count))
    
    # 数据方向 +1：outgoing，client to server；-1：incoming，server to client
    dirseq = list(item["DirSeq"])
    timeseq = list(item["Time"])  # 时间戳列表
    length = list(item["Length"])  # 包长度列表
    total_packet = len(dirseq)  # 总数据包数

    timeseq = [item - timeseq[0] + 1e-5 for item in timeseq]

    # f1 f2 First Request Content Size and RTT.
    outgoing_index = [j for j, x in enumerate(dirseq) if int(x) == 1]  # 传出数据包的index
    incoming_index = [j for j, x in enumerate(dirseq) if int(x) == -1]  # 传入数据包的index
    # f2:第一个传出数据包和第二个传出数据包之间传入数据包的大小
    if len(outgoing_index) <= 0:  # 如果全是传入/传出数据包
        f2 = 0
    else:
        # 第二个传出数据包与第一个传出数据包之间所有传入数据包的长度
        incoming_length = length[outgoing_index[0]:outgoing_index[1]]
        if len(incoming_length) == 1:
            f2 = 0
        else:
            f2 = sum(incoming_length[1:])

    # f1:第一个传出数据包和第一个传入数据包之间的延迟
    if len(outgoing_index) <= 0 or len(incoming_index) <= 0:
        f1 = 999999
    else:
        t1 = timeseq[[j for j, x in enumerate(
            dirseq) if int(x) == 1][0]]  # 第一个传出数据包的时间
        f1 = 999999  # 若只有第一个传出数据包, f1设为999999
        for i in [j for j, x in enumerate(dirseq) if int(x) == -1]:
            t2 = timeseq[i]  # 第一个在传出数据包后传入数据包的时间
            if t2 > t1:
                f1 = t2 - t1  # f1
                break

    # f3~f10 Statistics of packets size and number.
    f3 = sum([length[j]
                for j, x in enumerate(dirseq) if int(x) == 1])  # f3:传出数据包大小
    f303 = sum([length[j] for j, x in enumerate(
        dirseq) if int(x) == -1])  # f303:传入数据包大小
    f6 = len([j for j, x in enumerate(dirseq)
                if int(x) == -1])  # f6:传入数据包数
    f7 = len([j for j, x in enumerate(dirseq) if int(x) == 1])  # f7:传出数据包数
    f8 = total_packet  # f8:总数据包数
    f4 = f303 / sum(length)  # f4:传入数据包大小/总数据包大小
    f5 = f3 / sum(length)  # f5:传出数据包大小/总数据包大小
    f9 = f6 / f8  # f9:传入数据包数/总数据包数
    f10 = f7 / f8  # f10:传出数据包数/总数据包数

    # f11~f14 The number of incoming, outgoing packets, the fraction of the number of incoming packets, and the fraction of the number of outgoing packets in the first 20 packets of the network flows.
    f11 = len([j for j, x in enumerate(dirseq[:20]) if int(x) == -1]) if (total_packet > 20 or total_packet == 20) else len(
        [j for j, x in enumerate(dirseq) if int(x) == -1])  # 前20个包的传入数据包数
    f12 = len([j for j, x in enumerate(dirseq[:20]) if int(x) == 1]) if (total_packet > 20 or total_packet == 20) else len(
        [j for j, x in enumerate(dirseq) if int(x) == 1])  # 前20个包的传出数据包数
    f13 = f11 / 20 if (total_packet > 20 or total_packet ==
                        20) else f11 / len(dirseq)  # 传入数据包占前20个包的比例
    f14 = f12 / 20 if (total_packet > 20 or total_packet ==
                        20) else f12 / len(dirseq)  # 传出数据包占前20个包的比例

    # f15~f18 We generate two lists by recording the number of packets before every incoming and outgoing packet. Then we compute the average and standard deviation values of these two lists respectively
    l1 = [j for j, x in enumerate(dirseq) if int(x) == 1]  # 每个传出数据包前的数据包数量
    l2 = [j for j, x in enumerate(dirseq) if int(x) == -1]  # 每个传入数据包前的数据包数量
    f15 = np.mean(l1)  # 列表平均值
    f16 = np.mean(l2)  # 列表平均值
    f17 = np.std(l1)  # 列表标准差
    f18 = np.std(l2)  # 列表标准差

    # f19~f33 Statistics of packet inter-arrival time.
    lll1 = timeseq  # 总数据包时间序列
    lll1 = [data - lll1[0] for data in lll1]
    lll4 = [lll1[1] - lll1[0]]
    lll4.extend([lll1[n] - lll1[n - 1] for n in range(2, len(lll1))])
    f19 = np.max(lll4)
    f20 = np.min(lll4)
    f21 = np.mean(lll4)
    f22 = np.std(lll4)
    f23 = np.quantile(lll4, .75)

    if len([timeseq[j] for j, x in enumerate(dirseq) if int(x) == 1]) <= 1:  # 如果没有传出数据包/只有一个
        lll2 = [999999]
        lll5 = [999999]
        f24 = 999999
        f25 = 999999
        f26 = 999999
        f27 = 0
        f28 = 999999
    else:
        lll2 = [timeseq[j]
                for j, x in enumerate(dirseq) if int(x) == 1]  # 传出数据包时间序列
        lll2 = [data - lll1[0] for data in lll2]
        lll5 = [lll2[1] - lll2[0]]
        lll5.extend([lll2[n] - lll2[n - 1] for n in range(2, len(lll2))])
        f24 = np.max(lll5)
        f25 = np.min(lll5)
        f26 = np.mean(lll5)
        f27 = np.std(lll5)
        f28 = np.quantile(lll5, .75)

    if len([timeseq[j] for j, x in enumerate(dirseq) if int(x) == -1]) <= 1:
        lll3 = [999999]
        lll6 = [999999]
        f29 = 999999
        f30 = 999999
        f31 = 999999
        f32 = 0
        f33 = 999999
    else:
        lll3 = [timeseq[j]
                for j, x in enumerate(dirseq) if int(x) == -1]  # 传入数据包时间序列
        lll3 = [data - lll1[0] for data in lll3]
        lll6 = [lll3[1] - lll3[0]]
        lll6.extend([lll3[n] - lll3[n - 1] for n in range(2, len(lll3))])
        f29 = np.max(lll6)
        f30 = np.min(lll6)
        f31 = np.mean(lll6)
        f32 = np.std(lll6)
        f33 = np.quantile(lll6, .75)

    # f34~f42 Statistics of transmission time
    f34 = np.quantile(lll1, .25)
    f35 = np.quantile(lll1, .5)
    f36 = np.quantile(lll1, .75)
    f37 = np.quantile(lll2, .25)
    f38 = np.quantile(lll2, .5)
    f39 = np.quantile(lll2, .75)
    f40 = np.quantile(lll3, .25)
    f41 = np.quantile(lll3, .5)
    f42 = np.quantile(lll3, .75)
    feature = [f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22,
                f23, f24, f25, f26, f27, f28, f29, f30, f31, f32, f33, f34, f35, f36, f37, f38, f39, f40, f41, f42]

    # f43~102 The quantity and the transmission size speed of incoming, outgoing and total packets sequences.（用1除的）
    lll7 = [1 / j for j in lll4 if j != 0]
    lll8 = [1 / j for j in lll5 if j != 0]
    lll9 = [1 / j for j in lll6 if j != 0]
    f43_62 = random.sample(lll7, 20) if len(lll7) > 20 or len(
        lll7) == 20 else lll7 + [0] * (20 - len(lll7))
    f63_82 = random.sample(lll8, 20) if len(lll8) > 20 or len(
        lll8) == 20 else lll8 + [0] * (20 - len(lll8))
    f83_102 = random.sample(lll9, 20) if len(lll9) > 20 or len(
        lll9) == 20 else lll9 + [0] * (20 - len(lll9))
    feature.extend(f43_62)
    feature.extend(f63_82)
    feature.extend(f83_102)

    # f103~162 The quantity and the transmission size speed of incoming, outgoing and total packets sequences.
    llll4 = length
    lll10 = [llll4[j + 1] / x for j, x in enumerate(lll4) if x != 0]
    f103_122 = random.sample(lll10, 20) if len(lll10) > 20 or len(
        lll10) == 20 else lll10 + [0] * (20 - len(lll10))

    if len([timeseq[j] for j, x in enumerate(dirseq) if int(x) == 1]) <= 1:  # 如果传出数据包数为0/1
        f123_142 = [0] * 20
    else:
        llll5 = [length[j] for j, x in enumerate(dirseq) if int(x) == 1]
        lll11 = [llll5[j + 1] / x for j, x in enumerate(lll5) if x != 0]
        f123_142 = random.sample(lll11, 20) if len(lll11) > 20 or len(
            lll11) == 20 else lll11 + [0] * (20 - len(lll11))

    if len([timeseq[j] for j, x in enumerate(dirseq) if int(x) == -1]) <= 1:
        f143_162 = [0] * 20
    else:
        llll6 = [length[j] for j, x in enumerate(dirseq) if int(x) == -1]
        lll12 = [llll6[j + 1] / x for j, x in enumerate(lll6) if x != 0]
        f143_162 = random.sample(lll12, 20) if len(lll12) > 20 or len(
            lll12) == 20 else lll12 + [0] * (20 - len(lll12))

    feature.extend(f103_122)
    feature.extend(f123_142)
    feature.extend(f143_162)

    # f163~f262 The cumulative size of packets
    s = [0]
    for i in range(total_packet):
        s.append(int(dirseq[i])*length[i] + s[i])
    x = np.linspace(0, total_packet, total_packet+1)
    xvals = np.linspace(0, total_packet, 100)
    yinterp = np.interp(xvals, x, np.array(s))
    feature.extend(yinterp.tolist())

    # f263~302 Burst sizes and quantity
    a = [j for j, x in enumerate(dirseq) if int(x) == -1]
    
    brustSize = [sum(length[a[n - 1] + 1:a[n] + 1]) for n in range(1, len(a))]
    brustQuantity = [a[n] - a[n - 1] for n in range(1, len(a))]
    if len(brustSize) > 20 or len(brustSize) == 20:
        f263_282 = random.sample(brustSize, 20)
        f283_302 = random.sample(brustQuantity, 20)
    else:
        f263_282 = brustSize + [0] * (20 - len(brustSize))
        f283_302 = brustQuantity + [0] * (20 - len(brustQuantity))
    feature.extend(f263_282)
    feature.extend(f283_302)

    for i in range(len(feature)):
        if np.isnan(feature[i]):
            feature[i] = 0

    return feature, label

def get_dataset(a, client, taskid, groudByTime=False):
    """
    读入csv文件, 返回特征跟标签

    参数
    ----------
    dicloc: csv文件路径

    返回
    ----------
    x: 特征
    y: 标签
    """
    random.seed(27)
    X = []
    y = []
    if groudByTime:
        a = list(a.groupby(["Slice_id"]))
    else:
        a = list(a.groupby(["Slice_id", "RAN-UE-NGAP-ID"]))

    for item in a:
        if len(list(item[1]['Label'])) <= 10:
            continue

        feature, label = feature_extract(item[1])

        X.append(feature)
        y.append(label)

    return np.array(X), np.array(y)

