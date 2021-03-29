FROM golang:1.16.2-buster as BASE

RUN mkdir /build
COPY . /build
WORKDIR /build
RUN go build .
RUN ls -lh
RUN cp runner /bin/

FROM debian:buster

COPY --from=BASE /bin/runner /bin/runner
