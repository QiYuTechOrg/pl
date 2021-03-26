FROM debian:buster as BASE

ARG PHP_VERSION=8.0.3

RUN mkdir /build
WORKDIR /build

RUN apt update && apt install -y wget
RUN wget https://www.php.net/distributions/php-${PHP_VERSION}.tar.gz

RUN tar -zxvf php-${PHP_VERSION}.tar.gz

WORKDIR /build/php-${PHP_VERSION}
RUN apt install -y gcc pkg-config libxml2 libxml2-dev sqlite3 libsqlite3-dev
RUN ./configure --prefix=/lang

FROM debian:buster as BASE

COPY --from=BASE /lang /lang
