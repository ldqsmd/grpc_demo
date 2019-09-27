package  main

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc_demo/proto"
	"grpc_demo/server"
	"io/ioutil"
	"log"
	"net"
)

const (
	port  = "9001"
	srvPem = "./conf/cert/server/server.pem"
	srvKey = "./conf/cert/server/server.key"
	caPem  = "./conf/cert/ca.pem"
)


func main()  {

	//从证书相关文件中读取和解析信息 得到证书公钥 秘钥对
	cert, err := tls.LoadX509KeyPair(srvPem, srvKey)
	if err != nil {
		log.Fatalf("tls.LoadX509KeyPair err: %v", err)
	}
	//创建一个新的、空的 CertPool
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caPem)
	if err != nil {
		log.Fatalf("ioutil.ReadFile err: %v", err)
	}

	//尝试解析所传入的pem编码证书 如果解析成功 将其加到 cerpool 中 便于后面使用
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("certPool.AppendCertsFromPEM err")
	}
	//构建基于TLS的 TransportCredentials 选项
	c := credentials.NewTLS(&tls.Config{//Config 结构用于配置 TLS 客户端或服务器
		//设置证书链，允许包含一个或多个
		Certificates: []tls.Certificate{cert},
		//要求必须校验客户端的证书。可以根据实际情况选用以下参数：
		ClientAuth:   tls.RequireAndVerifyClientCert,
		//设置根证书的集合，校验方式使用 ClientAuth 中设定的模式
		ClientCAs:    certPool,
	})

	server := grpc.NewServer(grpc.Creds(c))
	//server := grpc.NewServer()
	pb.RegisterSearchServiceServer(server,&srv.SearchService{})

	list,err := net.Listen("tcp",":"+port)
	if err != nil {
	log.Fatalf("net.listen err:%v",err.Error())
	}
	server.Serve(list)
	}