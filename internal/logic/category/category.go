package category

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"jk-goframe-shop/internal/model/entity"

	"github.com/gogf/gf/v2/encoding/ghtml"
	"jk-goframe-shop/internal/service"

	"jk-goframe-shop/internal/dao"
	"jk-goframe-shop/internal/model"
)

type sCategory struct{}

func init() {
	service.RegisterCategory(New())
}

func New() *sCategory {
	return &sCategory{}
}

// Create 创建内容
func (s *sCategory) Create(ctx context.Context, in model.CategoryCreateInput) (out model.CategoryCreateOutput, err error) {

	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.CategoryInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.CategoryCreateOutput{CategoryId: int(lastInsertID)}, err
}

// Delete 删除
func (s *sCategory) Delete(ctx context.Context, id uint) error {
	return dao.CategoryInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		_, err := dao.CategoryInfo.Ctx(ctx).Where(g.Map{
			dao.CategoryInfo.Columns().Id: id,
		}).Delete()

		return err
	})
}

// Update 修改
func (s *sCategory) Update(ctx context.Context, in model.CategoryUpdateInput) error {
	return dao.CategoryInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		_, err := dao.CategoryInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.CategoryInfo.Columns().Id).
			Where(dao.CategoryInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}

// GetList 查询内容列表
func (s *sCategory) GetList(ctx context.Context, in model.CategoryGetListInput) (out *model.CategoryGetListOutput, err error) {
	var (
		m = dao.CategoryInfo.Ctx(ctx)
	)
	out = &model.CategoryGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)
	// 排序方式(dec 倒序)
	listModel = listModel.OrderDesc(dao.CategoryInfo.Columns().Id)

	// 执行查询
	var list []*entity.CategoryInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	// Category
	//if err := listModel.ScanList(&out.List, "Category"); err != nil {
	//	return out, err
	//}
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}

	return
}

// GetList 查询内容列表
// 返回所有数据
func (s *sCategory) GetAllList(ctx context.Context, in model.CategoryGetListInput) (out *model.CategoryGetListOutput, err error) {
	var (
		m = dao.CategoryInfo.Ctx(ctx)
	)
	out = &model.CategoryGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	//listModel := m.Page(in.Page, in.Size)
	//因为要返回所有的数据所以我们不需要Page分页
	listModel := m
	// 排序方式(dec 倒序)
	listModel = listModel.OrderDesc(dao.CategoryInfo.Columns().Id)

	// 执行查询
	var list []*entity.CategoryInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	// Category
	//if err := listModel.ScanList(&out.List, "Category"); err != nil {
	//	return out, err
	//}
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}

	return
}
