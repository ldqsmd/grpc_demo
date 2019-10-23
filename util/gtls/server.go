package gtls

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"google.golang.org/grpc/credentials"
	"grpc_demo/util/conf"
	"io/ioutil"
)

type Server struct {
	CaFile  string
	CerFile string
	KeyFile string
}

func NewServerTLS() *Server {

	return &Server{
		CerFile: conf.Config.TLS.CerFile,
		KeyFile: conf.Config.TLS.TLSKey,
	}

}

//tls证书验证
func (this *Server) GetTLSCredentials() (creds credentials.TransportCredentials, err error) {
	//TLS 认证
	creds, err = credentials.NewServerTLSFromFile(this.CerFile, this.KeyFile)
	return
}

//基于 CA 的 TLS 证书认证
func (this *Server) GetCredentiasCA() (creds credentials.TransportCredentials, err error) {

	//从证书相关文件中读取和解析信息 得到证书公钥 秘钥对
	cert, err := tls.LoadX509KeyPair(this.CerFile, this.KeyFile)
	if err != nil {
		return
	}
	//创建一个新的、空的 CertPool
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(this.CaFile)
	if err != nil {
		return
	}
	//尝试解析所传入的pem编码证书 如果解析成功 将其加到 cerpool 中 便于后面使用
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		err = errors.New("certPool.AppendCertsFromPEM err")
		return
	}
	//构建基于TLS的 TransportCredentials 选项
	creds = credentials.NewTLS(&tls.Config{ //Config 结构用于配置 TLS 客户端或服务器
		//设置证书链，允许包含一个或多个
		Certificates: []tls.Certificate{cert},
		//要求必须校验客户端的证书。可以根据实际情况选用以下参数：
		ClientAuth: tls.RequireAndVerifyClientCert,
		//设置根证书的集合，校验方式使用 ClientAuth 中设定的模式
		ClientCAs: certPool,
	})
	return
}
