package dict

import (
	"Ba-Server/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/dict/v1"
)

func (c *ControllerV1) SaveDict(ctx context.Context, req *v1.SaveDictReq) (res *v1.SaveDictRes, err error) {
	out, err := service.Dict().SaveDict(ctx, req)
	if err != nil {
		return nil, gerror.New("新增字典失败")
	}
	return out, err
}
