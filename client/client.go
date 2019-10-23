package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc_demo/proto"
	"log"
)

const (
	port   = "9001"
	srvPem = "./conf/cert/tls/tls.pem"
)

type Authentication struct {
	AppId  string
	AppKey string
}

//Token认证
func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{"appid": a.AppId, "appkey": a.AppKey}, nil
}
func (a *Authentication) RequireTransportSecurity() bool {
	return false
}

func main() {

	auth := Authentication{
		AppId:  "hello",
		AppKey: "world",
	}

	creds, err := credentials.NewClientTLSFromFile(srvPem, "grpc_demo")
	if err != nil {
		log.Fatalf("credentials.NewClientTLSFromFile 异常：%+v", err)
	}
	conn, err := grpc.Dial(":"+port, grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatalf("rpc dial 错误:%v", err.Error())
	}
	defer conn.Close()

	client := pb.NewSearchServiceClient(conn)
	res, err := client.Search(context.Background(), &pb.SearchRequest{
		Request: "hello world 怎么就请求失败呢",
	})
	if err != nil {
		fmt.Printf("search Err:%v \n", err.Error())
	}
	fmt.Printf("resaa:%+v \n", res.GetResponse())
}
