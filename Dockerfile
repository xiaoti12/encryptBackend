FROM ubuntu:20.04
VOLUME [ "/data" ]

RUN sed -i s@/archive.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list

# COPY sources.list /etc/apt/sources.list
COPY target/*.jar backend.jar
COPY core /core

ENV DEBIAN_FRONTEND=noninteractive
# 更新apt软件包列表并安装必要的工具
RUN apt-get clean && apt-get update && apt-get install -y curl gnupg2 libpcap-dev libssl-dev openjdk-17-jdk python3 python3-pip && pip3 install -r core/requirements.txt

# 更新zeek仓库源并安装zeek
RUN echo 'deb http://download.opensuse.org/repositories/security:/zeek/xUbuntu_20.04/ /' | tee /etc/apt/sources.list.d/security:zeek.list 
RUN curl -fsSL https://download.opensuse.org/repositories/security:zeek/xUbuntu_22.04/Release.key | gpg --dearmor | tee /etc/apt/trusted.gpg.d/security_zeek.gpg > /dev/null
RUN apt update && apt install zeek-lts -y

# 设置JAVA_HOME环境变量
ENV JAVA_HOME=/usr/lib/jvm/java-17-openjdk-amd64
# 设置ZEEK_HOME环境变量
ENV ZEEK_HOME=/opt/zeek/bin
# 设置PATH环境变量
ENV PATH=$PATH:$JAVA_HOME/bin:$ZEEK_HOME

ENTRYPOINT [ "java","-jar","backend.jar" ]