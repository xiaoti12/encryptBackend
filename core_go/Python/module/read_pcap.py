import pyshark
import pandas as pd
import numpy as np


# 读取pcap文件
def read_pcap(filepath, tsharkPath=None):
    feature_list = []
    RAN_IP = ['192.168.3.10', '20.0.2.10', '172.30.132.20']
    pcap_read = pyshark.FileCapture(
        filepath, display_filter='ngap && !http', tshark_path="D:\\Program Files\\Wireshark\\tshark.exe")
    print("Begin", filepath)
    for pkt in pcap_read:
        if "NGAP" in pkt:
            if pkt.ip.src in RAN_IP:
                Dirseq = 1
            else:
                Dirseq = -1

            packetLength = int(pkt.length)
            timeStamps = float(pkt.sniff_timestamp)
            try:
                ran_ue_ngap_id = int(pkt.ngap.ran_ue_ngap_id)
            except:
                ran_ue_ngap_id = -1

            feature_list.append([ran_ue_ngap_id, packetLength, timeStamps, Dirseq, 1, int(timeStamps / 10)])
    print("Finish", filepath)
    pcap_read.close()
    feature_list.sort(key=lambda x: x[2], reverse=False)
    df = pd.DataFrame(feature_list, columns=[
        "RAN-UE-NGAP-ID", "Length", "Time", "DirSeq", "Label", "Slice_id"])
    return df
