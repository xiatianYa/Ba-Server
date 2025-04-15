package dict

import (
	"Ba-Server/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/dict/v1"
)

func (c *ControllerV1) RemoveDictByIds(ctx context.Context, req *v1.RemoveDictByIdsReq) (res *v1.RemoveDictByIdsRes, err error) {
	out, err := service.Dict().RemoveDictByIds(ctx, req)
	if err != nil {
		return nil, gerror.New("删除字典失败")
	}
	return out, err
}
