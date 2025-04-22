package dict

import (
	"Ba-Server/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/dict/v1"
)

func (c *ControllerV1) GetDictInfoById(ctx context.Context, req *v1.GetDictInfoByIdReq) (res *v1.GetDictInfoByIdRes, err error) {
	out, err := service.Dict().GetDictInfoById(ctx, req)
	//发送错误
	if err != nil {
		return nil, gerror.New("获取字典信息失败" + err.Error())
	}
	return (*v1.GetDictInfoByIdRes)(out), err
}
