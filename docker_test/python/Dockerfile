FROM golang:1.16.2-buster as BASE

RUN mkdir /build
COPY . /build
WORKDIR /build
RUN go build .
RUN cp runner /bin/

FROM python:3.9

COPY --from=BASE /bin/runner /bin/runner
