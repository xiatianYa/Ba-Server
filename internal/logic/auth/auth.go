package auth

import (
	v1 "Ba-Server/api/auth/v1"
	"Ba-Server/internal/consts"
	"Ba-Server/internal/dao"
	"Ba-Server/internal/model/entity"
	"Ba-Server/internal/model/vo"
	"Ba-Server/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
)

type (
	sAuth struct{}
)

func init() {
	service.RegisterAuth(New())
}

func New() service.IAuth {
	return &sAuth{}
}

func (s sAuth) GetUserInfo(ctx context.Context, req *v1.GetUserInfoReq) (constantRoutes *vo.SysUserInfoVo, err error) {
	userId := ctx.Value(consts.UserId)
	if userId == nil {
		return nil, gerror.New("用户ID不存在")
	}
	userModel := dao.SysUser.Ctx(ctx)
	//查询用户信息
	var sysUser entity.SysUser
	_ = userModel.Where("id=?", userId).Scan(&sysUser)
	sysUserInfoVo, err := service.User().GetUserInfoVo(ctx, sysUser)
	return sysUserInfoVo, nil
}
