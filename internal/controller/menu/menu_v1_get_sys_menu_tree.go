package menu

import (
	"Ba-Server/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/menu/v1"
)

func (c *ControllerV1) GetSysMenuTree(ctx context.Context, req *v1.GetSysMenuTreeReq) (res *v1.GetSysMenuTreeRes, err error) {
	out, err := service.Menu().GetSysMenuTree(ctx, req)
	if err != nil {
		return nil, gerror.New("获取菜单树失败")
	}
	return (*v1.GetSysMenuTreeRes)(&out), nil
}
