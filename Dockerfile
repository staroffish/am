FROM registry.holygrail.com:5000/debian:buster-slim

RUN echo 'deb http://mirrors.aliyun.com/debian-archive/debian/ buster main non-free contrib\n\
deb http://mirrors.aliyun.com/debian-archive/debian-security buster/updates main\n\
deb http://mirrors.aliyun.com/debian-archive/debian/ buster-updates main non-free contrib\n\
deb-src http://mirrors.aliyun.com/debian-archive/debian/ buster main non-free contrib\n\
deb-src http://mirrors.aliyun.com/debian-archive/debian-security buster/updates main\n\
deb-src http://mirrors.aliyun.com/debian-archive/debian/ buster-updates main non-free contrib' > /etc/apt/sources.list

RUN cat /etc/apt/sources.list

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY bin/* /app/

RUN ls /app

WORKDIR /app