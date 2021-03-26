FROM debian:buster AS BASE

RUN mkdir /build
WORKDIR /build

ARG GCC_VERSION=gcc-10.1.0

# 安装所有的依赖
RUN apt update && apt install -y make wget

# 下载 gcc
RUN wget https://ftp.gnu.org/gnu/gcc/${GCC_VERSION}/${GCC_VERSION}.tar.gz
RUN tar -zxvf ${GCC_VERSION}.tar.gz

########################################
# 准备编译 GCC
#
WORKDIR /build/${GCC_VERSION}
RUN mkdir /lang
RUN apt -y install gcc g++ lbzip2
# 下载依赖
RUN ./contrib/download_prerequisites
# 不需要支持 32bit 的架构
RUN ./configure --prefix=/lang --disable-multilib
RUN make
RUN make install


FROM debian:buster

COPY --from=BASE /lang /lang
