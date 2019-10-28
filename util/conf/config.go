package conf

import (
	"fmt"
	"github.com/jinzhu/configor"
)

var Config = struct {

	/**RPC server 配置**/
	//GRPC 服务配置
	Server struct {
		ServerName string `default:"test"`
		Port       string `default:"9001"`
		//服务端token校验 appId appkey
		AppId  string `default:"appid"`
		AppKey string `default:"appkey"`
	}

	//TLS配置
	ServerTLS struct {
		ServerName    string `default:"grpc-server"`
		ServerTLSKey  string `default:"conf/cert/tls.key"`
		ServerCrtFile string `default:"conf/cert/tls.crt"`
	}

	CA struct {
		CAKey string `default:"conf/cert/ca.key"`
		CACrt string `default:"conf/cert/ca.crt"`
	}
	Trace struct {
		ServerName         string `default:"grpc-server"`
		LocalAgentHostPort string `defautl:"127.0.0.1:6831"`
		LogSpans           bool   `defautl:"false"`
	}

	/**RPC CLIENT 配置**/
	Client struct {
		//gRPC的服务名
		ServerName string `default:"grpc-server"`
		ServerPort string `default:"9001"`
		//服务端预留token校验appId/appkey
		AppId  string `default:"hello"`
		AppKey string `default:"world"`
	}

	ClientTLS struct {
		//生成TLS证书时serverName用于证书验证
		TLSServerName string `default:"grpc-server"`
		ClientTLSKey  string `default:"conf/cert/tls.key"`
		ClientCrtFile string `default:"conf/cert/tls.crt"`
	}
}{}

func InitConfig() {

	//初始化配置文件
	//配置加载顺序1.是否设置了变量conf，设置了第一个加载，如果文件不存在，加载默认配置文件
	//如果设置了环境变量 CONFIGOR_ENV = test等，那么加载config_test.yml的配置文件
	//最后加载环境变量,是否设置环境变量前缀,如果设置了CONFIGOR_ENV_PREFIX=WEB,设置环境变量为WEB_DB_NAME=root,否则为DB_NAME=root
	//DEBUG 是否开启调试模式
	configor.New(&configor.Config{Debug: false}).Load(&Config, "conf/config.yml")
	fmt.Printf("config: %+v \n", Config)
}
