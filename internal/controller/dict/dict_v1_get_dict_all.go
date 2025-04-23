package dict

import (
	"Ba-Server/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/dict/v1"
)

func (c *ControllerV1) GetDictAll(ctx context.Context, req *v1.GetDictAllReq) (res *v1.GetDictAllRes, err error) {
	out, err := service.Dict().GetDictAll(ctx)
	if err != nil {
		return nil, gerror.New("获取全部字典信息失败" + err.Error())
	}
	return (*v1.GetDictAllRes)(&out), nil
}
