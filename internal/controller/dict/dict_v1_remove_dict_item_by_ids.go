package dict

import (
	"Ba-Server/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/dict/v1"
)

func (c *ControllerV1) RemoveDictItemByIds(ctx context.Context, req *v1.RemoveDictItemByIdsReq) (res *v1.RemoveDictItemByIdsRes, err error) {
	out, err := service.Dict().RemoveDictItemByIds(ctx, req)
	if err != nil {
		return nil, gerror.New("删除子字典失败" + err.Error())
	}
	return out, err
}
