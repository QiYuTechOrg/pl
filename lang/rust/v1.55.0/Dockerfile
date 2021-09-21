FROM ghcr.io/qiyutechdev/runner:v0.2.0 as BASE
# 因为我们使用 github Action
# 所以这儿使用 github container 地址 [速度比较快]

FROM rust:1.55.0

COPY --from=BASE /bin/runner /bin/runner
RUN useradd runner
