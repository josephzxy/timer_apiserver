FROM golang:1.17.1-bullseye AS build

COPY . /app
WORKDIR /app
ARG ARCH=amd64

RUN make build PLATFORM=linux_${ARCH} && \
    mv /app/_output/build/linux/${ARCH}/apiserver /app/apiserver

FROM alpine:latest
COPY --from=build /app/apiserver /app/apiserver
WORKDIR /app
RUN mkdir -p /app/log && \
    mkdir -p /app/config
ENTRYPOINT ["/bin/sh", "-c", "/app/apiserver 1>>/app/log/stdout.log 2>>/app/log/stderr.log"]
