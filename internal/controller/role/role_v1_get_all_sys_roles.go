package role

import (
	"Ba-Server/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/role/v1"
)

func (c *ControllerV1) GetAllSysRoles(ctx context.Context, req *v1.GetAllSysRolesReq) (res *v1.GetAllSysRolesRes, err error) {
	out, err := service.Role().GetAllSysRolesReq(ctx)
	if err != nil {
		return nil, gerror.New("获取全部角色失败")
	}
	return (*v1.GetAllSysRolesRes)(&out), err
}
