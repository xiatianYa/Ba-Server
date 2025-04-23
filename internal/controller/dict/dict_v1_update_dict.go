package dict

import (
	"Ba-Server/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/dict/v1"
)

func (c *ControllerV1) UpdateDict(ctx context.Context, req *v1.UpdateDictReq) (res *v1.UpdateDictRes, err error) {
	out, err := service.Dict().UpdateDict(ctx, req)
	if err != nil {
		return nil, gerror.New("修改字典失败" + err.Error())
	}
	return out, err
}
