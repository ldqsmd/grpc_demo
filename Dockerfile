FROM golang:latest
#配置环境变量代理
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /grpc_demo
COPY . /grpc_demo
RUN go build .
EXPOSE 8000
ENTRYPOINT ["./grpc_demo"]