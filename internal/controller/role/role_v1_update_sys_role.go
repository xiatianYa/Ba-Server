package role

import (
	"Ba-Server/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/role/v1"
)

func (c *ControllerV1) UpdateSysRole(ctx context.Context, req *v1.UpdateSysRoleReq) (res *v1.UpdateSysRoleRes, err error) {
	out, err := service.Role().UpdateSysRole(ctx, req)
	if err != nil {
		return nil, gerror.New("修改用户信息失败" + err.Error())
	}
	return out, err
}
