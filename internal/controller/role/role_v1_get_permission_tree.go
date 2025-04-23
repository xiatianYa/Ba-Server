package role

import (
	"Ba-Server/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/role/v1"
)

func (c *ControllerV1) GetPermissionTree(ctx context.Context, req *v1.GetPermissionTreeReq) (res *v1.GetPermissionTreeRes, err error) {
	out, err := service.Role().GetPermissionTree(ctx, req)
	if err != nil {
		return nil, gerror.New("获取按钮树失败" + err.Error())
	}
	return (*v1.GetPermissionTreeRes)(&out), nil
}
