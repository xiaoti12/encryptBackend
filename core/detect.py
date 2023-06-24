import argparse
import paramiko


def run_remote_script(ip, username, password, script_path, script_args, python_path):
    # 创建SSH客户端对象
    ssh_client = paramiko.SSHClient()
    # 自动添加服务器的主机密钥
    ssh_client.set_missing_host_key_policy(paramiko.AutoAddPolicy())

    try:
        # 连接到远程服务器
        ssh_client.connect(ip, username=username, password=password)

        # 构建要执行的命令
        command = f'{python_path} {script_path} {script_args}'

        # 执行远程命令
        stdin, stdout, stderr = ssh_client.exec_command(command)

        # 获取命令输出
        output = stdout.read().decode('utf-8')
        error = stderr.read().decode('utf-8')

        # 输出执行结果
        print('执行结果:')
        if output:
            print(output)
        if error:
            print(f'错误信息: {error}')
    finally:
        # 关闭SSH连接
        ssh_client.close()


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('--file_path', required=True, type=str)
    parser.add_argument('--model_type', default='bin', type=str)

    args = parser.parse_args()
    # 服务器信息
    server_ip = '10.3.242.94'
    username = 'heweichun'
    password = 'Hwc123!'

    # 要执行的脚本和参数
    python_path = '/data/heweichun/anaconda3/envs/xgboost38/bin/python'
    script_path = '/data/heweichun/Code/N2_multi_web/springboot.py'
    file_path = '/data/heweichun/Code/N2_multi_web/pcap/oai1.pcapng'
    file_path = args.file_path

    # 将参数拼接为一个字符串，使用空格分隔
    script_args = f'--file_path {file_path}'

    # 运行远程脚本
    run_remote_script(server_ip, username, password, script_path, script_args, python_path)
