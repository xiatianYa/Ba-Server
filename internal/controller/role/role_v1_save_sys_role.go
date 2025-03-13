package role

import (
	"Ba-Server/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/role/v1"
)

func (c *ControllerV1) SaveSysRole(ctx context.Context, req *v1.SaveSysRoleReq) (res *v1.SaveSysRoleRes, err error) {
	out, err := service.Role().SaveSysRole(ctx, req)
	if err != nil {
		return nil, gerror.New("新增角色失败")
	}
	return out, err
}
