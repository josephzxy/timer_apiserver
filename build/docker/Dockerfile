FROM golang:1.17.1-bullseye AS build

COPY . /app
WORKDIR /app
ARG OS=linux
ARG ARCH=arm64
RUN make build PLATFORM=${OS}_${ARCH} && \
    mv /app/_output/build/${OS}/${ARCH}/apiserver /app/apiserver

FROM alpine:latest
COPY --from=build /app/apiserver /app/apiserver
WORKDIR /app
ENTRYPOINT ["/app/apiserver"]
