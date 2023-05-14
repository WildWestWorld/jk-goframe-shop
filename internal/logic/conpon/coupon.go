package coupon

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

type sCoupon struct{}

func init() {
	service.RegisterCoupon(New())
}

func New() *sCoupon {
	return &sCoupon{}
}

// Create 创建内容
func (s *sCoupon) Create(ctx context.Context, in model.CouponCreateInput) (out model.CouponCreateOutput, err error) {

	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.CouponInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.CouponCreateOutput{CouponId: int(lastInsertID)}, err
}

// Delete 删除
func (s *sCoupon) Delete(ctx context.Context, id uint) error {
	return dao.CouponInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		_, err := dao.CouponInfo.Ctx(ctx).Where(g.Map{
			dao.CouponInfo.Columns().Id: id,
		}).Delete()

		return err
	})
}

// Update 修改
func (s *sCoupon) Update(ctx context.Context, in model.CouponUpdateInput) error {
	return dao.CouponInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		_, err := dao.CouponInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.CouponInfo.Columns().Id).
			Where(dao.CouponInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}

// GetList 查询内容列表
func (s *sCoupon) GetList(ctx context.Context, in model.CouponGetListInput) (out *model.CouponGetListOutput, err error) {
	var (
		m = dao.CouponInfo.Ctx(ctx)
	)
	out = &model.CouponGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)
	// 排序方式(dec 倒序)
	listModel = listModel.OrderDesc(dao.CouponInfo.Columns().Price)

	// 执行查询
	var list []*entity.CouponInfo
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
	// Coupon
	//if err := listModel.ScanList(&out.List, "Coupon"); err != nil {
	//	return out, err
	//}
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}

	return
}

// GetList 查询内容列表
// 返回所有数据
func (s *sCoupon) GetAllList(ctx context.Context, in model.CouponGetListInput) (out *model.CouponGetListOutput, err error) {
	var (
		m = dao.CouponInfo.Ctx(ctx)
	)
	out = &model.CouponGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	//listModel := m.Page(in.Page, in.Size)
	//因为要返回所有的数据所以我们不需要Page分页
	listModel := m
	// 排序方式(dec 倒序)
	listModel = listModel.OrderDesc(dao.CouponInfo.Columns().Price)

	// 执行查询
	var list []*entity.CouponInfo
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
	// Coupon
	//if err := listModel.ScanList(&out.List, "Coupon"); err != nil {
	//	return out, err
	//}
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}

	return
}
