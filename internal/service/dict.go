// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "Ba-Server/api/dict/v1"
	"Ba-Server/internal/model/vo"
	"context"
)

type (
	IDict interface {
		// GetAllDict GetAllDictReq 获取全部字典Options(系统初始)
		GetAllDict(ctx context.Context) (dictMap map[string][]vo.SysDictItemOptionsVo, err error)
		// GetDictAll 获取全部字典Vo(管理页)
		GetDictAll(ctx context.Context) (dictVos []vo.SysDictVo, err error)
		// GetDictItemPage GetDictItemPageReq 分页获取子字典
		GetDictItemPage(ctx context.Context, req *v1.GetDictItemPageReq) (total int, records []vo.SysDictItemVo, err error)
		// SaveDict 添加字典
		SaveDict(ctx context.Context, req *v1.SaveDictReq) (res *v1.SaveDictRes, err error)
		// UpdateDict 修改字典
		UpdateDict(ctx context.Context, req *v1.UpdateDictReq) (res *v1.UpdateDictRes, err error)
		// RemoveDictByIds 删除多个字典
		RemoveDictByIds(ctx context.Context, req *v1.RemoveDictByIdsReq) (res *v1.RemoveDictByIdsRes, err error)
		// GetDictInfoById 获取字典详细
		GetDictInfoById(ctx context.Context, req *v1.GetDictInfoByIdReq) (res *vo.SysDictVo, err error)
		// SaveDictItem 添加子字典
		SaveDictItem(ctx context.Context, req *v1.SaveDictItemReq) (res *v1.SaveDictItemRes, err error)
		// RemoveDictItemByIds 删除多个子字段
		RemoveDictItemByIds(ctx context.Context, req *v1.RemoveDictItemByIdsReq) (res *v1.RemoveDictItemByIdsRes, err error)
		// GetDictItemInfoById 获取子字典详细
		GetDictItemInfoById(ctx context.Context, req *v1.GetDictItemInfoByIdReq) (res *vo.SysDictItemVo, err error)
		// UpdateDictItem 修改子字典
		UpdateDictItem(ctx context.Context, req *v1.UpdateDictItemReq) (res *v1.UpdateDictItemRes, err error)
	}
)

var (
	localDict IDict
)

func Dict() IDict {
	if localDict == nil {
		panic("implement not found for interface IDict, forgot register?")
	}
	return localDict
}

func RegisterDict(i IDict) {
	localDict = i
}
