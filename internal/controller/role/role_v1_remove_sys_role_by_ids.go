package role

import (
	"Ba-Server/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/role/v1"
)

func (c *ControllerV1) RemoveSysRoleByIds(ctx context.Context, req *v1.RemoveSysRoleByIdsReq) (res *v1.RemoveSysRoleByIdsRes, err error) {
	out, err := service.Role().RemoveSysRoleByIds(ctx, req)
	if err != nil {
		return nil, gerror.New("删除角色失败" + err.Error())
	}
	return out, err
}
