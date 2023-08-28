import argparse
from FourTuple.FourTuplePredictorClass import FourTuplePredictor
from FiveTuple.FiveTuplePredictorClass import FiveTuplePredictor
import os
import clickhouse_connect

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("--file_path", type=str, help="path of pcap")
    parser.add_argument("--model", type=str, default='five', choices=['four', 'five'], help="type of model")
    parser.add_argument("--taskid", type=str, help="task id")

    args = parser.parse_args()
    model_type = args.model
    pcap_path = args.file_path
    task_id = args.taskid

    core_path=os.path.join(os.getcwd(),"core_python")
    os.chdir(core_path)

    pcap_path = os.path.abspath(os.path.join("upload",pcap_path))
    # print(f"=======WORKING DIR:{os.getcwd()}=======")
    # print(f"==========={pcap_path}===========")
    if not os.path.exists(pcap_path):
        raise FileNotFoundError("pcap does not exist!")

    # pcap_path = "/home/fsc/liujy/yahong/pcap/test.pcap"
    # model_type = "five"
    # task_id = "b41ffd1c-a873-43b4-b58f-a56de1c50148"

    client = clickhouse_connect.get_client(host='10.3.242.84', password='password', database='TLS')

    if model_type == 'four':
        preditor = FourTuplePredictor()
    elif model_type == 'five':
        preditor = FiveTuplePredictor()
    preditor.predict_result(pcap_path)

    task_results = preditor.get_task_results()
    # flow_results = preditor.process_results(os.path.basename(pcap_path))
    flow_results = preditor.process_results(task_id=task_id)
    # print(f"===========TASK ID:{task_id} {task_results}===========")
    client.command(
        f"ALTER TABLE Task UPDATE `normal`={task_results[0]},`abnormal`={task_results[1]},`total`={task_results[2]} WHERE `taskId`='{task_id}'")
    client.insert('TLSFlow', flow_results)
