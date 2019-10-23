package conf

import (
	"fmt"
	"github.com/jinzhu/configor"
)

var Config = struct {

	//GRPC 服务配置
	Server struct {
		ServerName string `default:"test"`
		Port       string `default:"9001"`
		AppId      string `default:"appid"`
		AppKey     string `default:"appkey"`
	}

	//TLS配置
	TLS struct {
		ServerName string `default:"grpc_demo"`
		TLSKey     string `default:"conf/cert/tls/tls.key"`
		CerFile    string `default:"conf/cert/tls/tls.pem"`
	}

	//接口token认证
	AuthToken struct {
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
