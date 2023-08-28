import pandas as pd

FEATURES=["ts","uid","id.orig_h","id.orig_p","id.resp_h","id.resp_p","duration","or_spl","server_version","client_version","cipher","client_ciphers","sni","ssl_client_exts","ssl_server_exts","ssl_established","client_supported_versions","server_supported_version","issuer","subject","cert_num","not_before","not_after","san_num","ext_num"]

FEATURES_TYPES={
    "ts": "float",
    "uid": "str",
    "id.orig_h": "str",
    "id.orig_p": "int",
    "id.resp_h": "str",
    "id.resp_p": "int",
    "duration": "float",
    "or_spl": "str",
    "server_version": "int",
    "client_version": "int",
    "cipher": "int",
    "client_ciphers": "str",
    "sni": "str",
    "ssl_client_exts": "str",
    "ssl_server_exts": "str",
    "ssl_established": "str",
    "client_supported_versions": "str",
    "server_supported_version": "str",
    "issuer": "str",
    "subject": "str",
    "cert_num": "int",
    "not_before": "float",
    "not_after": "float",
    "san_num": "int",
    "ext_num": "int"
}

def conn_to_df(conn_list):
    # df=pd.DataFrame(conn_list,columns=FEATURES)
    conn_dict={i[0]:i[1] for i in zip(FEATURES,conn_list)}
    df=pd.DataFrame(conn_dict,index=[0])
    df=df.astype(FEATURES_TYPES)
    df=df.astype({c:"datetime64[s]" for c in ["ts","not_before","not_after"]})
    print(df)

if __name__=="__main__":
    conn_str="1681656965.302255	CoErlH3ag2mcaOFSYj	10.1.10.101	49221	31.24.228.175	443	0.129982	-121,1234,-134,59,-229,181	769	769	49172	47,53,5,10,49171,49172,49161,49162,50,56,19,4	jupiterpluto.best	65281,0,10,11	65281,11	T	-	-	CN=rfnvPDdyNK.com	CN=rfnvPDdyNK.com	1	1578618051.000000	1610154051.000000	-	3"
    conn_list=conn_str.split("\t")
    conn_to_df(conn_list)
