package menu

import (
	"Ba-Server/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/menu/v1"
)

func (c *ControllerV1) GetSysMenuIdsByRoleId(ctx context.Context, req *v1.GetSysMenuIdsByRoleIdReq) (res *v1.GetSysMenuIdsByRoleIdRes, err error) {
	out, err := service.Menu().GetSysMenuIdsByRoleId(ctx, req)
	if err != nil {
		return nil, gerror.New("获取角色菜单列表失败")
	}
	return (*v1.GetSysMenuIdsByRoleIdRes)(&out), nil
}
