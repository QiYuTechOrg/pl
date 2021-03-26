FROM debian:buster as BASE


RUN mkdir /build
WORKDIR /build

ARG PY_VERSION=3.9.2

RUN apt update && apt install -y make wget

RUN wget https://www.python.org/ftp/python/${PY_VERSION}/Python-${PY_VERSION}.tgz
RUN tar -zxvf Python-${PY_VERSION}.tgz

RUN mkdir /lang
RUN apt -y install gcc
WORKDIR /build/Python-${PY_VERSION}

################################
# 编译 Python
RUN ./configure --enable-optimizations --prefix=/lang
