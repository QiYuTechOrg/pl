FROM golang:1.16.2-buster as BASE

RUN mkdir /build
WORKDIR build
COPY . /build
RUN go build .
COPY runner /bin/

FROM debian:buster

COPY --from=BASE /bin/runner /bin/runner
