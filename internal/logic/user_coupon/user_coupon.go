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

type sUserCoupon struct{}

func init() {
	service.RegisterUserCoupon(New())
}

func New() *sUserCoupon {
	return &sUserCoupon{}
}

// Create 创建内容
func (s *sUserCoupon) Create(ctx context.Context, in model.UserCouponCreateInput) (out model.UserCouponCreateOutput, err error) {

	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.UserCouponInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.UserCouponCreateOutput{UserCouponId: int(lastInsertID)}, err
}

// Delete 删除
func (s *sUserCoupon) Delete(ctx context.Context, id uint) error {
	return dao.UserCouponInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		_, err := dao.UserCouponInfo.Ctx(ctx).Where(g.Map{
			dao.UserCouponInfo.Columns().Id: id,
		}).Delete()

		return err
	})
}

// Update 修改
func (s *sUserCoupon) Update(ctx context.Context, in model.UserCouponUpdateInput) error {
	return dao.UserCouponInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		_, err := dao.UserCouponInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.UserCouponInfo.Columns().Id).
			Where(dao.UserCouponInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}

// GetList 查询内容列表
func (s *sUserCoupon) GetList(ctx context.Context, in model.UserCouponGetListInput) (out *model.UserCouponGetListOutput, err error) {
	var (
		m = dao.UserCouponInfo.Ctx(ctx)
	)
	out = &model.UserCouponGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)
	// 排序方式(dec 倒序)
	listModel = listModel.OrderDesc(dao.UserCouponInfo.Columns().Id)

	// 执行查询
	var list []*entity.UserCouponInfo
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
	// UserCoupon
	//if err := listModel.ScanList(&out.List, "UserCoupon"); err != nil {
	//	return out, err
	//}
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}

	return
}

// GetList 查询内容列表
// 返回所有数据
func (s *sUserCoupon) GetAllList(ctx context.Context, in model.UserCouponGetListInput) (out *model.UserCouponGetListOutput, err error) {
	var (
		m = dao.UserCouponInfo.Ctx(ctx)
	)
	out = &model.UserCouponGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	//listModel := m.Page(in.Page, in.Size)
	//因为要返回所有的数据所以我们不需要Page分页
	listModel := m
	// 排序方式(dec 倒序)
	listModel = listModel.OrderDesc(dao.UserCouponInfo.Columns().Id)

	// 执行查询
	var list []*entity.UserCouponInfo
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
	// UserCoupon
	//if err := listModel.ScanList(&out.List, "UserCoupon"); err != nil {
	//	return out, err
	//}
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}

	return
}
