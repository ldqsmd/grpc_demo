package mid

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"grpc_demo/util/conf"
	"log"
)

type Authentication struct {
	AppId  string
	AppKey string
}

func (a *Authentication) Auth(ctx context.Context) error {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("panic 异常 %+v", p)
		}
	}()
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing credentials")
	}
	var appid string
	var appkey string

	if val, ok := md["appid"]; ok {
		appid = val[0]
	}
	if val, ok := md["appkey"]; ok {
		appkey = val[0]
	}
	fmt.Printf("MD信息 %#v \n", md)
	if appid != a.AppId || appkey != a.AppKey {
		return status.Errorf(codes.Unauthenticated, "token 不正确")
	}
	return nil
}

func AuthRequestToken(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println("fileter:", info)
	auth := new(Authentication)
	auth.AppKey = conf.Config.Server.AppKey
	auth.AppId = conf.Config.Server.AppId
	err = auth.Auth(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return handler(ctx, req)
}
