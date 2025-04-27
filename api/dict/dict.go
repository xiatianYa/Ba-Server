// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dict

import (
	"context"

	"Ba-Server/api/dict/v1"
)

type IDictV1 interface {
	GetAllDict(ctx context.Context, req *v1.GetAllDictReq) (res *v1.GetAllDictRes, err error)
	GetDictAll(ctx context.Context, req *v1.GetDictAllReq) (res *v1.GetDictAllRes, err error)
	GetDictItemPage(ctx context.Context, req *v1.GetDictItemPageReq) (res *v1.GetDictItemPageRes, err error)
	SaveDict(ctx context.Context, req *v1.SaveDictReq) (res *v1.SaveDictRes, err error)
	UpdateDict(ctx context.Context, req *v1.UpdateDictReq) (res *v1.UpdateDictRes, err error)
	UpdateDictItem(ctx context.Context, req *v1.UpdateDictItemReq) (res *v1.UpdateDictItemRes, err error)
	RemoveDictByIds(ctx context.Context, req *v1.RemoveDictByIdsReq) (res *v1.RemoveDictByIdsRes, err error)
	GetDictInfoById(ctx context.Context, req *v1.GetDictInfoByIdReq) (res *v1.GetDictInfoByIdRes, err error)
	GetDictItemInfoById(ctx context.Context, req *v1.GetDictItemInfoByIdReq) (res *v1.GetDictItemInfoByIdRes, err error)
	SaveDictItem(ctx context.Context, req *v1.SaveDictItemReq) (res *v1.SaveDictItemRes, err error)
	RemoveDictItemByIds(ctx context.Context, req *v1.RemoveDictItemByIdsReq) (res *v1.RemoveDictItemByIdsRes, err error)
}
