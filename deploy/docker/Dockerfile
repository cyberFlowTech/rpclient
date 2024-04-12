FROM golang:alpine as builder

WORKDIR /app
COPY . .
RUN apk add curl  && go env -w GO111MODULE=on \
&& go env -w GOPROXY=https://goproxy.cn,direct \
&& go env -w GOPRIVATE=github.com/cyberFlowTech \
&& go env -w CGO_ENABLED=0 \
&& go install github.com/zeromicro/goctl-swagger@latest \
&& go install github.com/zeromicro/go-zero/tools/goctl@latest
