FROM ubuntu:20.04
VOLUME [ "/data" ]

RUN sed -i s@/archive.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list

# COPY sources.list /etc/apt/sources.list
COPY target/*.jar backend.jar
COPY core /core

ENV DEBIAN_FRONTEND=noninteractive
# 更新apt软件包列表并安装必要的工具
RUN apt-get clean && apt-get update && apt-get install -y curl gnupg2

# 安装OpenJDK 17 and Python
RUN apt-get install -y openjdk-17-jdk python3 python3-pip

RUN pip3 install -r core/requirements.txt

# 设置JAVA_HOME环境变量
ENV JAVA_HOME=/usr/lib/jvm/java-17-openjdk-amd64

# 设置PATH环境变量
ENV PATH=$PATH:$JAVA_HOME/bin

ENTRYPOINT [ "java","-jar","backend.jar" ]