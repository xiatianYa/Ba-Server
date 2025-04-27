package v1

import (
	"Ba-Server/internal/model/domain"
	"Ba-Server/internal/model/vo"
	"github.com/gogf/gf/v2/frame/g"
)

type GetAllDictReq struct {
	g.Meta `path:"/allDict" tags:"getAllDict" method:"get" summary:"获取全部字典(初始化)"`
}

type GetAllDictRes map[string][]vo.SysDictItemOptionsVo

type GetDictAllReq struct {
	g.Meta `path:"/dictAll" tags:"getAllDict" method:"get" summary:"获取全部字典(管理页)"`
}

type GetDictAllRes []vo.SysDictVo

type GetDictItemPageReq struct {
	g.Meta      `path:"/itemPage" tags:"GetDictItemPageReq" method:"get" summary:"获取子字典(分页)"`
	Current     int    `v:"required" json:"current"`
	Size        int    `v:"required" json:"size"`
	DictId      int64  `json:"dictId"`
	Value       string `json:"value"`
	Description string `json:"description"`
	ZhCn        string `json:"zhCn"`
	EnUs        string `json:"enUs"`
}

type GetDictItemPageRes *domain.RPage

type SaveDictReq struct {
	g.Meta      `path:"/save" method:"post" summary:"新增字典"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Type        string `json:"type"`
	Status      string `json:"status"`
	Sort        int    `json:"sort"`
	Description string `json:"description"`
}

type SaveDictRes bool

type UpdateDictReq struct {
	g.Meta      `path:"/update" method:"put" summary:"修改字典"`
	Id          int64  `v:"required" json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Type        string `json:"type"`
	Status      string `json:"status"`
	Sort        int    `json:"sort"`
	Description string `json:"description"`
}

type UpdateDictRes bool

type UpdateDictItemReq struct {
	g.Meta      `path:"/updateItem" method:"put" summary:"修改字典"`
	Id          int64  `v:"required" json:"id"`
	Value       string `json:"value"`
	Description string `json:"description"`
	ZhCn        string `json:"zhCn"`
	EnUs        string `json:"enUs"`
	Status      string `json:"status"`
	Sort        int    `json:"sort"`
	Type        string `json:"type"`
}

type UpdateDictItemRes bool

type RemoveDictByIdsReq struct {
	g.Meta `path:"/removeByIds" tags:"removeByIds" method:"delete" summary:"删除多个字典"`
	Ids    []int64 `v:"required" json:"ids"`
}

type RemoveDictByIdsRes bool

type GetDictInfoByIdReq struct {
	g.Meta `path:"/getInfo/{id}" tags:"getInfo" method:"get" summary:"根据字典Id获取字典详细"`
	Id     int64 `v:"required" dc:"dict id"`
}

type GetDictInfoByIdRes vo.SysDictVo

type GetDictItemInfoByIdReq struct {
	g.Meta `path:"/getItemInfo/{id}" tags:"getItemInfo" method:"get" summary:"根据字典Id获取子字典详细"`
	Id     int64 `v:"required" dc:"dict id"`
}

type GetDictItemInfoByIdRes vo.SysDictItemVo

type SaveDictItemReq struct {
	g.Meta      `path:"/saveItem" method:"post" summary:"新增子字典"`
	DictId      int64  `v:"required" json:"dictId"`
	DictCode    string `json:"dictCode"`
	Value       string `json:"value"`
	ZhCn        string `json:"zhCn"`
	EnUs        string `json:"enUs"`
	Status      string `json:"status"`
	Sort        int    `json:"sort"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type SaveDictItemRes bool

type RemoveDictItemByIdsReq struct {
	g.Meta `path:"/removeItemByIds" tags:"removeItemByIds" method:"delete" summary:"删除多个子字典"`
	Ids    []int64 `v:"required" json:"ids"`
}

type RemoveDictItemByIdsRes bool
