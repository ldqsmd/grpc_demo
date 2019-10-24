package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	pb "grpc_demo/proto"
	"grpc_demo/util/auth"
	"grpc_demo/util/conf"
	"grpc_demo/util/gtls"
	"log"
	"net/http"
)

func init() {
	conf.InitConfig()
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	client := gtls.NewClientTLS()
	creds, err := client.GetCredentiasCA()
	if err != nil {
		log.Fatal(err)
	}

	option := []grpc.DialOption{
		//证书校验
		grpc.WithTransportCredentials(creds),
		//token认证
		grpc.WithPerRPCCredentials(&auth.Authentication{
			conf.Config.Client.AppId,
			conf.Config.Client.AppKey,
		}),
	}
	err = pb.RegisterSearchServiceHandlerFromEndpoint(ctx, mux, "localhost:9001", option)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("rpc  rest-api start")
	err = http.ListenAndServe(":8899", mux)
	if err != nil {
		log.Fatal(err)
	}

}
