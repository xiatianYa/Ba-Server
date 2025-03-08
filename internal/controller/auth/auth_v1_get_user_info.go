package auth

import (
	"Ba-Server/internal/model/api"
	"Ba-Server/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"Ba-Server/api/auth/v1"
)

func (c *ControllerV1) GetUserInfo(ctx context.Context, req *v1.GetUserInfoReq) (res *v1.GetUserInfoRes, err error) {
	token, err := service.Auth().GetUserInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	g.RequestFromCtx(ctx).Response.WriteJson(api.Result{Code: 200, Data: token, Msg: "success"})
	return
}
