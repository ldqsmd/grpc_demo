package auth

import (
	"context"
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
