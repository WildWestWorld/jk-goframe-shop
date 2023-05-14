package user_coupon

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

type sGoodsOptions struct{}

func init() {
	service.RegisterGoodsOptions(New())
}

func New() *sGoodsOptions {
	return &sGoodsOptions{}
}

// Create 创建内容
func (s *sGoodsOptions) Create(ctx context.Context, in model.GoodsOptionsCreateInput) (out model.GoodsOptionsCreateOutput, err error) {

	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.GoodsOptionsInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.GoodsOptionsCreateOutput{GoodsOptionsId: int(lastInsertID)}, err
}

// Delete 删除
func (s *sGoodsOptions) Delete(ctx context.Context, id uint) error {
	return dao.GoodsOptionsInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		_, err := dao.GoodsOptionsInfo.Ctx(ctx).Where(g.Map{
			dao.GoodsOptionsInfo.Columns().Id: id,
		}).Delete()

		return err
	})
}

// Update 修改
func (s *sGoodsOptions) Update(ctx context.Context, in model.GoodsOptionsUpdateInput) error {
	return dao.GoodsOptionsInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		_, err := dao.GoodsOptionsInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.GoodsOptionsInfo.Columns().Id).
			Where(dao.GoodsOptionsInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}

// GetList 查询内容列表
func (s *sGoodsOptions) GetList(ctx context.Context, in model.GoodsOptionsGetListInput) (out *model.GoodsOptionsGetListOutput, err error) {
	var (
		m = dao.GoodsOptionsInfo.Ctx(ctx)
	)
	out = &model.GoodsOptionsGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)
	// 排序方式(dec 倒序)
	listModel = listModel.OrderDesc(dao.GoodsOptionsInfo.Columns().Id)

	// 执行查询
	var list []*entity.GoodsOptionsInfo
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
	// GoodsOptions
	//if err := listModel.ScanList(&out.List, "GoodsOptions"); err != nil {
	//	return out, err
	//}
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}

	return
}

// GetList 查询内容列表
// 返回所有数据
func (s *sGoodsOptions) GetAllList(ctx context.Context, in model.GoodsOptionsGetListInput) (out *model.GoodsOptionsGetListOutput, err error) {
	var (
		m = dao.GoodsOptionsInfo.Ctx(ctx)
	)
	out = &model.GoodsOptionsGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	//listModel := m.Page(in.Page, in.Size)
	//因为要返回所有的数据所以我们不需要Page分页
	listModel := m
	// 排序方式(dec 倒序)
	listModel = listModel.OrderDesc(dao.GoodsOptionsInfo.Columns().Id)

	// 执行查询
	var list []*entity.GoodsOptionsInfo
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
	// GoodsOptions
	//if err := listModel.ScanList(&out.List, "GoodsOptions"); err != nil {
	//	return out, err
	//}
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}

	return
}
