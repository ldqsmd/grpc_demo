package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc_demo/proto"
	"grpc_demo/util/auth"
	"grpc_demo/util/conf"
	"grpc_demo/util/gtls"
	"log"
)

func init() {
	conf.InitConfig()
}

func main() {
	cliTLS := gtls.NewClientTLS()
	//creds, err := cliTLS.GetTLSCredentials()
	creds, err := cliTLS.GetCredentiasCA()
	if err != nil {
		log.Fatalf("GetCredentiasCA：%+v", err)
	}
	conn, err := grpc.Dial(":"+conf.Config.Client.ServerPort,
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(&auth.Authentication{
			conf.Config.Client.AppId,
			conf.Config.Client.AppKey,
		}))
	if err != nil {
		log.Fatalf("rpc dial 错误:%v", err.Error())
	}
	defer conn.Close()

	client := pb.NewSearchServiceClient(conn)
	res, err := client.Search(context.Background(), &pb.SearchRequest{})
	if err != nil {
		fmt.Printf("search Err:%v \n", err.Error())
		return
	}
	fmt.Printf("resaa:%+v \n", res.GetResponse())

	e, _ := client.Echo(context.Background(), &pb.StringMessage{
		Words: "哎哟哟",
	})
	fmt.Printf("echo :%+v \n", e.GetWords())

}
