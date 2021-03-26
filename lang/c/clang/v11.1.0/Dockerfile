FROM debian:buster AS BASE

RUN mkdir /build
WORKDIR /build

# 安装所有的依赖
RUN apt update && apt install -y make wget xz-utils

# 下载 llvm && clang
RUN wget https://github.com/llvm/llvm-project/releases/download/llvmorg-11.1.0/llvm-11.1.0.src.tar.xz
RUN tar -xJf llvm-11.1.0.src.tar.xz

RUN wget https://github.com/llvm/llvm-project/releases/download/llvmorg-11.1.0/clang-11.1.0.src.tar.xz
RUN tar -xJf clang-11.1.0.src.tar.xz

########################################
# 准备编译 llvm && clang
#
RUN mkdir /lang
RUN apt -y install gcc g++ lbzip2 cmake python3

# build llvm
RUN mkdir /build/llvm
WORKDIR /build/llvm
RUN cmake -DCMAKE_BUILD_TYPE=Release -DCMAKE_INSTALL_PREFIX:PATH=/lang ../llvm-11.1.0.src
RUN make
RUN make install

RUN mkdir /build/clang
WORKDIR /build/clang
RUN cmake -DCMAKE_BUILD_TYPE=Release -DCMAKE_INSTALL_PREFIX:PATH=/lang ../clang-11.1.0.src
RUN make  .
RUN make install


FROM debian:buster

COPY --from=BASE /lang /lang
