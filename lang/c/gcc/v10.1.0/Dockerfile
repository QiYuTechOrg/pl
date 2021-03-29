FROM ghcr.io/qiyutechdev/runner:v0.0.1p0 as BASE
# 因为我们使用 github Action
# 所以这儿使用 github container 地址 [速度比较快]

FROM gcc:10.1.0

COPY --from=BASE /bin/runner /bin/runner
RUN useradd runner
