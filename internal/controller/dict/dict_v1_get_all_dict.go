package dict

import (
	"Ba-Server/internal/service"
	"context"

	"Ba-Server/api/dict/v1"
)

func (c *ControllerV1) GetAllDict(ctx context.Context, req *v1.GetAllDictReq) (res *v1.GetAllDictRes, err error) {
	out, err := service.Dict().GetAllDict(ctx)
	if err != nil {
		return nil, err
	}
	return (*v1.GetAllDictRes)(&out), nil
}
