package role

import (
	"Ba-Server/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/role/v1"
)

func (c *ControllerV1) UpdatePermissionIdsByRoleId(ctx context.Context, req *v1.UpdatePermissionIdsByRoleIdReq) (res *v1.UpdatePermissionIdsByRoleIdRes, err error) {
	out, err := service.Role().UpdatePermissionIdsByRoleId(ctx, req)
	if err != nil {
		return nil, gerror.New("修改角色按钮失败" + err.Error())
	}
	return out, err
}
