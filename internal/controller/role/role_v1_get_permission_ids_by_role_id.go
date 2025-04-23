package role

import (
	"Ba-Server/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/role/v1"
)

func (c *ControllerV1) GetPermissionIdsByRoleId(ctx context.Context, req *v1.GetPermissionIdsByRoleIdReq) (res *v1.GetPermissionIdsByRoleIdRes, err error) {
	out, err := service.Role().GetPermissionIdsByRoleId(ctx, req)
	if err != nil {
		return nil, gerror.New("获取角色按钮列表失败" + err.Error())
	}
	return (*v1.GetPermissionIdsByRoleIdRes)(&out), nil
}
