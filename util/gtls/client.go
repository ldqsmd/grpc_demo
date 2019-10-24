package gtls

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"grpc_demo/util/conf"
	"io/ioutil"
)

type Client struct {
	TLSServerName string
	CACrt         string
	ClientCrtFile string
	ClientTLSKey  string
}

func NewClientTLS() *Client {
	return &Client{
		TLSServerName: conf.Config.ClientTLS.TLSServerName,
		ClientCrtFile: conf.Config.ClientTLS.ClientCrtFile,
		ClientTLSKey:  conf.Config.ClientTLS.ClientTLSKey,
		CACrt:         conf.Config.CA.CACrt,
	}
}

//tls证书验证
func (this *Client) GetTLSCredentials() (creds credentials.TransportCredentials, err error) {
	//TLS 认证
	creds, err = credentials.NewClientTLSFromFile(conf.Config.ClientTLS.ClientCrtFile, conf.Config.ClientTLS.TLSServerName)
	if err != nil {
		err = status.Errorf(codes.NotFound, "NewClientTLSFromFile 错误 %+v", err.Error())
	}
	return
}

//基于 CA 的 TLS 证书认证
func (this *Client) GetCredentiasCA() (creds credentials.TransportCredentials, err error) {

	//从证书相关文件中读取和解析信息 得到证书公钥 秘钥对
	certificate, err := tls.LoadX509KeyPair(conf.Config.ClientTLS.ClientCrtFile, conf.Config.ClientTLS.ClientTLSKey)
	if err != nil {
		err = errors.New("LoadX509KeyPair" + err.Error())
		return
	}
	//创建一个新的、空的 CertPool
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(conf.Config.CA.CACrt)
	if err != nil {
		err = errors.New("ReadFile:" + err.Error())
		return
	}

	//尝试解析所传入的pem编码证书 如果解析成功 将其加到 cerpool 中 便于后面使用
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		err = errors.New("certPool.AppendCertsFromPEM err")
		return
	}
	creds = credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ServerName:   this.TLSServerName, // NOTE: this is required!
		RootCAs:      certPool,
	})
	return
}
