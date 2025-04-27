package dict

import (
	v1 "Ba-Server/api/dict/v1"
	"Ba-Server/internal/dao"
	"Ba-Server/internal/model/entity"
	"Ba-Server/internal/model/vo"
	"Ba-Server/internal/service"
	"context"
)

type (
	sDict struct{}
)

func init() {
	service.RegisterDict(New())
}

func New() service.IDict {
	return &sDict{}
}

// GetAllDict GetAllDictReq 获取全部字典Options(系统初始)
func (s sDict) GetAllDict(ctx context.Context) (dictMap map[string][]vo.SysDictItemOptionsVo, err error) {
	result := make(map[string][]vo.SysDictItemOptionsVo)
	var SysDicts []entity.SysDict
	var SysDictItems []entity.SysDictItem

	DictModel := dao.SysDict.Ctx(ctx)
	DictItemModel := dao.SysDictItem.Ctx(ctx)

	//查询所有的字典 子字典
	_ = DictModel.Scan(&SysDicts)
	_ = DictItemModel.Scan(&SysDictItems)

	result = getDictItemOptions(SysDicts, SysDictItems)
	return result, nil
}

// GetDictAll 获取全部字典Vo(管理页)
func (s sDict) GetDictAll(ctx context.Context) (dictVos []vo.SysDictVo, err error) {
	var SysDicts []entity.SysDict

	DictModel := dao.SysDict.Ctx(ctx)

	//查询所有的字典
	_ = DictModel.Scan(&SysDicts)

	//拷贝属性
	sysDictVo := make([]vo.SysDictVo, len(SysDicts))
	for i, item := range SysDicts {
		sysDictVo[i] = vo.SysDictVo{
			Id:          item.Id,
			Name:        item.Name,
			Code:        item.Code,
			Type:        item.Type,
			Sort:        item.Sort,
			Description: item.Description,
			Status:      item.Status,
			CreateTime:  item.CreateTime,
			UpdateTime:  item.UpdateTime,
		}
	}
	return sysDictVo, nil
}

// GetDictItemPage GetDictItemPageReq 分页获取子字典
func (s sDict) GetDictItemPage(ctx context.Context, req *v1.GetDictItemPageReq) (total int, records []vo.SysDictItemVo, err error) {
	// 创建指针切片
	var result []vo.SysDictItemVo

	DictItemModel := dao.SysDictItem.Ctx(ctx)
	pageQuery := DictItemModel.OmitEmpty().Page(req.Current, req.Size)
	pageQuery = pageQuery.Where("dict_id", req.DictId)
	pageQuery = pageQuery.Where("value", req.Value)
	pageQuery = pageQuery.WhereLike("description", "%"+req.Description+"%")
	pageQuery = pageQuery.WhereLike("zh_cn", "%"+req.ZhCn+"%")
	pageQuery = pageQuery.WhereLike("en_us", "%"+req.EnUs+"%")

	err = pageQuery.ScanAndCount(&result, &total, true)

	if err != nil {
		return 0, nil, err
	}

	// 如果列表为空
	if result == nil {
		records = []vo.SysDictItemVo{}
		return
	}

	records = result
	return
}

// SaveDict 添加字典
func (s sDict) SaveDict(ctx context.Context, req *v1.SaveDictReq) (res *v1.SaveDictRes, err error) {
	dictModel := dao.SysDict.Ctx(ctx)
	dict := entity.SysDict{
		Name:        req.Name,
		Code:        req.Code,
		Type:        req.Type,
		Sort:        req.Sort,
		Description: req.Description,
		Status:      req.Status,
	}
	_, err = dictModel.Insert(dict)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// UpdateDict 修改字典
func (s sDict) UpdateDict(ctx context.Context, req *v1.UpdateDictReq) (res *v1.UpdateDictRes, err error) {
	dictModel := dao.SysDict.Ctx(ctx)
	dict := entity.SysDict{
		Id:          req.Id,
		Name:        req.Name,
		Code:        req.Code,
		Type:        req.Type,
		Sort:        req.Sort,
		Description: req.Description,
		Status:      req.Status,
	}
	_, err = dictModel.OmitEmpty().Data(dict).Where("id", req.Id).Update()
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// RemoveDictByIds 删除多个字典
func (s sDict) RemoveDictByIds(ctx context.Context, req *v1.RemoveDictByIdsReq) (res *v1.RemoveDictByIdsRes, err error) {
	DictModel := dao.SysDict.Ctx(ctx)
	_, err = DictModel.WhereIn("id", req.Ids).Delete()
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// GetDictInfoById 获取字典详细
func (s sDict) GetDictInfoById(ctx context.Context, req *v1.GetDictInfoByIdReq) (res *vo.SysDictVo, err error) {
	var dictVo *vo.SysDictVo
	dictModel := dao.SysDict.Ctx(ctx)
	err = dictModel.Where("id", req.Id).Scan(&dictVo)
	if err != nil {
		return nil, err
	}
	return dictVo, nil
}

// SaveDictItem 添加子字典
func (s sDict) SaveDictItem(ctx context.Context, req *v1.SaveDictItemReq) (res *v1.SaveDictItemRes, err error) {
	dictItemModel := dao.SysDictItem.Ctx(ctx)
	dict := entity.SysDictItem{
		DictId:      req.DictId,
		DictCode:    req.DictCode,
		Value:       req.Value,
		ZhCn:        req.ZhCn,
		EnUs:        req.EnUs,
		Type:        req.Type,
		Sort:        req.Sort,
		Description: req.Description,
		Status:      req.Status,
	}
	_, err = dictItemModel.Insert(dict)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// RemoveDictItemByIds 删除多个子字段
func (s sDict) RemoveDictItemByIds(ctx context.Context, req *v1.RemoveDictItemByIdsReq) (res *v1.RemoveDictItemByIdsRes, err error) {
	DictItemModel := dao.SysDictItem.Ctx(ctx)
	_, err = DictItemModel.WhereIn("id", req.Ids).Delete()
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// getDictItemOptions 构建字典 MapVo
func getDictItemOptions(sysDicts []entity.SysDict, sysDictItems []entity.SysDictItem) map[string][]vo.SysDictItemOptionsVo {
	result := make(map[string][]vo.SysDictItemOptionsVo)
	for _, dict := range sysDicts {
		var options []vo.SysDictItemOptionsVo
		for _, item := range sysDictItems {
			if item.DictId == dict.Id {
				vo := vo.SysDictItemOptionsVo{
					Label: item.ZhCn,
					Value: item.Value,
					Type:  item.Type,
					Sort:  item.Sort,
				}
				options = append(options, vo)
			}
		}
		result[dict.Code] = options
	}
	return result
}

// GetDictItemInfoById 获取子字典详细
func (s sDict) GetDictItemInfoById(ctx context.Context, req *v1.GetDictItemInfoByIdReq) (res *vo.SysDictItemVo, err error) {
	var dictItemVo *vo.SysDictItemVo
	dictItemModel := dao.SysDictItem.Ctx(ctx)
	err = dictItemModel.Where("id", req.Id).Scan(&dictItemVo)
	if err != nil {
		return nil, err
	}
	return dictItemVo, nil
}

// UpdateDictItem 修改子字典
func (s sDict) UpdateDictItem(ctx context.Context, req *v1.UpdateDictItemReq) (res *v1.UpdateDictItemRes, err error) {
	dictItemModel := dao.SysDictItem.Ctx(ctx)
	dict := entity.SysDictItem{
		Id:          req.Id,
		Value:       req.Value,
		ZhCn:        req.ZhCn,
		EnUs:        req.EnUs,
		Status:      req.Status,
		Sort:        req.Sort,
		Type:        req.Type,
		Description: req.Description,
	}
	_, err = dictItemModel.OmitEmpty().Data(dict).Where("id", req.Id).Update()
	if err != nil {
		return nil, err
	}
	return nil, nil
}
