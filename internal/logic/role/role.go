package role

import (
	"Ba-Server/internal/consts"
	"Ba-Server/internal/dao"
	"Ba-Server/internal/model/vo"
	"Ba-Server/internal/service"
	"context"
)

type (
	sRole struct{}
)

func init() {
	service.RegisterRole(New())
}

func New() service.IRole {
	return &sRole{}
}

func (s sRole) GetAllRolesReq(ctx context.Context) (rolesVo []vo.SysRoleVo, err error) {
	var sysRoleVos []vo.SysRoleVo
	roleModel := dao.SysRole.Ctx(ctx)
	_ = roleModel.Where("status = ?", consts.ONE).Scan(&sysRoleVos)
	return sysRoleVos, nil
}
