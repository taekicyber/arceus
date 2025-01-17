FROM node:15.5 as front
WORKDIR /work
ADD . .
RUN make ui

FROM golang:1.16-alpine as builder
ENV GOPROXY=https://goproxy.io,direct
ENV GO111MODULE=on

RUN apk add gcc g++ make git libffi-dev openssl-dev libtool

WORKDIR /go/cache
ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /work
ADD . .
COPY --from=front /work/static/ui /work/static/ui
RUN GOOS=linux CGO_ENABLED=0 make build

FROM gruebel/upx:latest as upx
WORKDIR /work
COPY --from=builder /work/arceus /work/arceus.pre
RUN upx --best --lzma -o /work/arceus arceus.pre

FROM alpine:3.6 as alpine
RUN apk add -U --no-cache ca-certificates

FROM alpine
MAINTAINER zc
LABEL maintainer="zc" \
    email="zc2638@qq.com" \
    version="v0.6"

ENV TS="Asia/Shanghai"
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=upx /work/arceus /work/arceus

WORKDIR /work
CMD ["./arceus"]
