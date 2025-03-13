package role

import (
	"Ba-Server/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/role/v1"
)

func (c *ControllerV1) UpdateRoleMenu(ctx context.Context, req *v1.UpdateRoleMenuReq) (res *v1.UpdateRoleMenuRes, err error) {
	out, err := service.Role().UpdateRoleMenu(ctx, req)
	if err != nil {
		return nil, gerror.New("修改角色菜单失败")
	}
	return out, err
}
