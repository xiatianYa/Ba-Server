package dict

import (
	"Ba-Server/internal/service"
	"context"

	"Ba-Server/api/dict/v1"
)

func (c *ControllerV1) GetDictAll(ctx context.Context, req *v1.GetDictAllReq) (res *v1.GetDictAllRes, err error) {
	out, err := service.Dict().GetDictAll(ctx)
	if err != nil {
		return nil, err
	}
	return (*v1.GetDictAllRes)(&out), nil
}
