// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package route

import (
	"context"

	"Ba-Server/api/route/v1"
)

type IRouteV1 interface {
	GetConstantRoutes(ctx context.Context, req *v1.GetConstantRoutesReq) (res *v1.GetConstantRoutesRes, err error)
	GetUserRoutes(ctx context.Context, req *v1.GetUserRoutesReq) (res *v1.GetUserRoutesRes, err error)
}
