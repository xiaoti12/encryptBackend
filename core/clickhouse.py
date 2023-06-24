import clickhouse_driver
from clickhouse_driver import Client
from datetime import datetime
import threading

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



client = Client(host='10.21.244.125', user='default', password='qwe123', database='helloworld')

select = "Select * from helloworld.task"

print(client.execute(select))