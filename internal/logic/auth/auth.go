package auth

import (
	"Ba-Server/internal/service"
	"context"
)

type (
	sAuth struct{}
)

func init() {
	service.RegisterAuth(New())
}

func New() service.IAuth {
	return &sAuth{}
}

func (s sAuth) UserNameLogin(ctx context.Context) (constantRoutes string, err error) {
	//TODO implement me
	panic("implement me")
}
