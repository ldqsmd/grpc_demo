package main

import (
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"grpc_demo/mid"
	pb "grpc_demo/proto"
	"grpc_demo/server"
	"grpc_demo/util/conf"
	"grpc_demo/util/gtls"
	"grpc_demo/util/trace"
	"log"
	"net"
)

func init() {
	conf.InitConfig()
	trace.JaegerConfigInit()
}

func main() {

	list, err := net.Listen("tcp", ":"+conf.Config.Server.Port)
	if err != nil {
		log.Fatalf("net.listen err:%v", err.Error())
	}
	serverTLS := gtls.NewServerTLS()
	//TLS认证
	//creds, err := serverTLS.GetTLSCredentials()
	//CA证书认证
	creds, err := serverTLS.GetCredentiasCA()
	if err != nil {
		log.Fatalf("tls 异常 异常：%+v", err)
	}

	server := grpc.NewServer(
		grpc.Creds(creds),
		grpc.UnaryInterceptor(
			//引入开源RPC中间件
			grpc_middleware.ChainUnaryServer(
				mid.AuthRequestToken), //token 认证
		),
	)
	pb.RegisterSearchServiceServer(server, &srv.SearchService{})

	fmt.Println("服务启动成功")

	if err := server.Serve(list); err != nil {
		log.Fatalf("grpc异常：%+v", err)
	}

	defer trace.TraceClose()
}
