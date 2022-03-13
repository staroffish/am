FROM golang:1.17 AS builder

COPY .  /src
WORKDIR /src

RUN GOPROXY=https://goproxy.io make build

FROM debian:buster-slim

RUN echo 'deb http://mirrors.163.com/debian/ buster main non-free contrib\n\
deb http://mirrors.163.com/debian/ buster-updates main non-free contrib\n\
deb http://mirrors.163.com/debian/ buster-backports main non-free contrib\n\
deb-src http://mirrors.163.com/debian/ buster main non-free contrib\n\
deb-src http://mirrors.163.com/debian/ buster-updates main non-free contrib\n\
deb-src http://mirrors.163.com/debian/ buster-backports main non-free contrib\n\
deb http://mirrors.163.com/debian-security/ buster/updates main non-free contrib\n\
deb-src http://mirrors.163.com/debian-security/ buster/updates main non-free contrib\n' > /etc/apt/sources.list

RUN cat /etc/apt/sources.list

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin /app

RUN ls /app

WORKDIR /app
