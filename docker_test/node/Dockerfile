FROM golang:1.16.2-buster as BASE

RUN mkdir /build
COPY runner /build
WORKDIR /build
RUN go build .
RUN cp runner /bin/

FROM node:v16

COPY --from=BASE /bin/runner /bin/runner
