package dict

import (
	"Ba-Server/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/dict/v1"
)

func (c *ControllerV1) GetDictItemInfoById(ctx context.Context, req *v1.GetDictItemInfoByIdReq) (res *v1.GetDictItemInfoByIdRes, err error) {
	out, err := service.Dict().GetDictItemInfoById(ctx, req)
	//发送错误
	if err != nil {
		return nil, gerror.New("获取子字典信息失败" + err.Error())
	}
	return (*v1.GetDictItemInfoByIdRes)(out), err
}
