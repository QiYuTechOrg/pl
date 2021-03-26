FROM debian:buster

RUN apt update && apt install -y curl
RUN curl https://sh.rustup.rs -sSf | sh -s -- -y --profile=minimal --default-toolchain 1.51.0
