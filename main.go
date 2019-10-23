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
	"log"
	"net"
)

func init() {
	conf.InitConfig()
}

func main() {

	list, err := net.Listen("tcp", ":"+conf.Config.Server.Port)
	if err != nil {
		log.Fatalf("net.listen err:%v", err.Error())
	}
	//TLS 认证
	serverTLS := gtls.NewServerTLS()
	creds, err := serverTLS.GetTLSCredentials()
	if err != nil {
		log.Fatalf("credentials.NewClientTLSFromFile 异常：%+v", err)
	}
	//grpc.UnaryInterceptor()
	server := grpc.NewServer(
		grpc.Creds(creds),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(mid.AuthRequestToken),
		),
	)
	pb.RegisterSearchServiceServer(server, &srv.SearchService{})

	fmt.Println("服务启动成功")

	if err := server.Serve(list); err != nil {
		log.Fatalf("grpc异常：%+v", err)
	}
}
