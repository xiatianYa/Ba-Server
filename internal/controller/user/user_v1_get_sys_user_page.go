package user

import (
	"Ba-Server/api/user/v1"
	"Ba-Server/internal/model/domain"
	"Ba-Server/internal/service"
	"context"
)

func (c *ControllerV1) GetSysUserPage(ctx context.Context, req *v1.GetSysUserPageReq) (res *v1.GetSysUserPageRes, err error) {
	pages, total, records, err := service.User().GetSysUserPage(ctx, req)
	rPage := &domain.RPage{
		Current: req.Current,
		Size:    req.Size,
		Pages:   pages,
		Total:   total,
		Records: records,
	}
	return (*v1.GetSysUserPageRes)(&rPage), nil
}
