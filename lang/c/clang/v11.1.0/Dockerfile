FROM debian:buster AS BASE

RUN mkdir /build
WORKDIR /build

# 安装所有的依赖
RUN apt update && apt install -y make wget xz

# 下载 clang
RUN wget https://github.com/llvm/llvm-project/releases/download/llvmorg-11.1.0/clang-11.1.0.src.tar.xz
RUN tar -xvJf clang-11.1.0.src.tar.xz

########################################
# 准备编译 clang
#
WORKDIR /build/clang-11.1.0
RUN mkdir /lang
RUN apt -y install gcc g++ lbzip2

RUN ./configure --prefix=/lang
RUN make
RUN make install


FROM debian:buster

COPY --from=BASE /lang /lang
