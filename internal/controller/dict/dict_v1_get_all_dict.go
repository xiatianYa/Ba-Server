package dict

import (
	"Ba-Server/api/dict/v1"
	"Ba-Server/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) GetAllDict(ctx context.Context, req *v1.GetAllDictReq) (res *v1.GetAllDictRes, err error) {
	out, err := service.Dict().GetAllDict(ctx)
	if err != nil {
		return nil, gerror.New("获取全部字典信息失败" + err.Error())
	}
	return (*v1.GetAllDictRes)(&out), nil
}
