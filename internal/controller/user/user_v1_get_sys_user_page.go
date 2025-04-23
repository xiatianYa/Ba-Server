package user

import (
	"Ba-Server/api/user/v1"
	"Ba-Server/internal/model/domain"
	"Ba-Server/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) GetSysUserPage(ctx context.Context, req *v1.GetSysUserPageReq) (res *v1.GetSysUserPageRes, err error) {
	total, records, err := service.User().GetSysUserPage(ctx, req)

	//发送错误
	if err != nil {
		return nil, gerror.New("获取用户分页异常" + err.Error())
	}

	//构建返回对象
	rPage := &domain.RPage{
		Current: req.Current,
		Size:    req.Size,
		Pages:   (total + req.Size - 1) / req.Size,
		Total:   total,
		Records: records,
	}
	return (*v1.GetSysUserPageRes)(&rPage), nil
}
