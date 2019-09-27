package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc_demo/proto"
	"io/ioutil"
	"log"
)

const (
	port  = "9001"
	cliPem = "./conf/cert/client/client.pem"
	cliKey = "./conf/cert/client/client.key"
	caPem  = "./conf/cert/ca.pem"
)
func main()  {

	cert, err := tls.LoadX509KeyPair(cliPem, cliKey)
	if err != nil {
		log.Fatalf("tls.LoadX509KeyPair err: %v", err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caPem)
	if err != nil {
		log.Fatalf("ioutil.ReadFile err: %v", err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("certPool.AppendCertsFromPEM err")
	}

	c := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "go-grpc-example",
		RootCAs:      certPool,
	})

	conn ,err := grpc.Dial(":"+port,grpc.WithTransportCredentials(c))
	//conn ,err := grpc.Dial(":"+port,grpc.WithInsecure())
	if err != nil{
		log.Fatalf("rpc dial 错误:%v",err.Error())
	}
	defer conn.Close()

	client := pb.NewSearchServiceClient(conn)
	res ,err := client.Search(context.Background(),&pb.SearchRequest{
		Request:"hello world",
	})
	if err != nil {
		fmt.Errorf("search Err:%v",err.Error())
	}
	fmt.Printf("res:%+v \n",res)
}