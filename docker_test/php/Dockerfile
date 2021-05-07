FROM golang:1.16.2-buster as BASE

RUN mkdir /build
COPY runner /build
WORKDIR /build
RUN go build .
RUN cp runner /bin/

FROM php:8.0.5

COPY --from=BASE /bin/runner /bin/runner
