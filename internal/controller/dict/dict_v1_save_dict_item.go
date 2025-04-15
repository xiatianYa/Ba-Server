package dict

import (
	"Ba-Server/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/dict/v1"
)

func (c *ControllerV1) SaveDictItem(ctx context.Context, req *v1.SaveDictItemReq) (res *v1.SaveDictItemRes, err error) {
	out, err := service.Dict().SaveDictItem(ctx, req)
	if err != nil {
		return nil, gerror.New("新增子字典失败")
	}
	return out, err
}
