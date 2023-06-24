from clickhouse_driver import Client

from datetime import datetime
import threading
import joblib
import pandas as pd
import module.feature_extraction as fe

def insert_result(task_id, status):
    conn = clickhouse_driver.connect(
        host='your_host',
        port='your_port',
        user='your_username',
        password='your_password',
        database='your_database',
    )
    
    try:
        cursor = conn.cursor()
        cursor.execute('INSERT INTO result (task_id, status) VALUES (%s, %s)', (task_id, status))
        conn.commit()
    finally:
        conn.close()


def monitor_status():
    conn = clickhouse_driver.connect(
        host='10.21.244.125', 
        user='default', 
        password='qwe123', 
        database='helloworld'
    )
    cursor = conn.cursor()
    cursor.execute('SELECT task_id, status FROM task')
    previous_status = {row[0]: row[1] for row in cursor.fetchall()}
    
    while True:
        cursor.execute('SELECT task_id, status FROM task')
        current_status = {row[0]: row[1] for row in cursor.fetchall()}
        
        for task_id, status in current_status.items():
            if task_id not in previous_status or status != previous_status[task_id]:
                threading.Thread(target=insert_result, args=(task_id, status)).start()
        
        previous_status = current_status
        time.sleep(1)  # 适当的睡眠时间，避免频繁查询数据库

    conn.close()



client = Client(host='10.3.242.84', user='default', password='password', database='SCTP')

select = "SELECT * FROM SCTP.Packet WHERE FlowTimeID = 13226433187010039149 ORDER BY ArriveTime"
result = client.execute(select)
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
feature,label = fe.feature_extract(df)

model_path = "python/model/model_multi_UEID.pkl"
model = joblib.load(model_path)
y_predict = model.predict([feature])[0]
print(y_predict)
