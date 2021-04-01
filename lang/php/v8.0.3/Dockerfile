FROM ghcr.io/qiyutechdev/runner:v0.0.6 as BASE
# 因为我们使用 github Action
# 所以这儿使用 github container 地址 [速度比较快]

FROM php:8.0.3

COPY --from=BASE /bin/runner /bin/runner
RUN useradd runner
