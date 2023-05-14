package goods

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"jk-goframe-shop/internal/consts"
	"jk-goframe-shop/internal/logic/collection"
	"jk-goframe-shop/internal/model/entity"

	"github.com/gogf/gf/v2/encoding/ghtml"
	"jk-goframe-shop/internal/service"

	"jk-goframe-shop/internal/dao"
	"jk-goframe-shop/internal/model"
)

type sGoods struct{}

func init() {
	service.RegisterGoods(New())
}

func New() *sGoods {
	return &sGoods{}
}

// Create 创建内容
func (s *sGoods) Create(ctx context.Context, in model.GoodsCreateInput) (out model.GoodsCreateOutput, err error) {

	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.GoodsInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.GoodsCreateOutput{GoodsId: int(lastInsertID)}, err
}

// Delete 删除
func (s *sGoods) Delete(ctx context.Context, id uint) error {
	return dao.GoodsInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		_, err := dao.GoodsInfo.Ctx(ctx).Where(g.Map{
			dao.GoodsInfo.Columns().Id: id,
		}).Delete()

		return err
	})
}

// Update 修改
func (s *sGoods) Update(ctx context.Context, in model.GoodsUpdateInput) error {
	return dao.GoodsInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		_, err := dao.GoodsInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.GoodsInfo.Columns().Id).
			Where(dao.GoodsInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}

// GetList 查询内容列表
func (s *sGoods) GetList(ctx context.Context, in model.GoodsGetListInput) (out *model.GoodsGetListOutput, err error) {
	var (
		m = dao.GoodsInfo.Ctx(ctx)
	)
	out = &model.GoodsGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)
	// 排序方式(dec 倒序)
	listModel = listModel.OrderDesc(dao.GoodsInfo.Columns().Id)

	// 执行查询
	var list []*entity.GoodsInfo
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
	// Goods
	//if err := listModel.ScanList(&out.List, "Goods"); err != nil {
	//	return out, err
	//}
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}

	return
}

// GetList 查询内容列表
// 返回所有数据
func (s *sGoods) GetAllList(ctx context.Context, in model.GoodsGetListInput) (out *model.GoodsGetListOutput, err error) {
	var (
		m = dao.GoodsInfo.Ctx(ctx)
	)
	out = &model.GoodsGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	//listModel := m.Page(in.Page, in.Size)
	//因为要返回所有的数据所以我们不需要Page分页
	listModel := m
	// 排序方式(dec 倒序)
	listModel = listModel.OrderDesc(dao.GoodsInfo.Columns().Id)

	// 执行查询
	var list []*entity.GoodsInfo
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
	// Goods
	//if err := listModel.ScanList(&out.List, "Goods"); err != nil {
	//	return out, err
	//}
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}

	return
}

// 商品详情
func (s *sGoods) GoodsDetail(ctx context.Context, in model.GoodsDetailInput) (out model.GoodsDetailOutput, err error) {
	err = dao.GoodsInfo.Ctx(ctx).WithAll().WherePri(in.Id).Scan(&out)
	if err != nil {
		return model.GoodsDetailOutput{}, err
	}

	out.IsCollect, err = collection.CheckIsCollection(ctx, model.CheckIsCollectInput{
		UserId:   gconv.Uint(ctx.Value(consts.CtxUserId)),
		ObjectId: in.Id,
		Type:     consts.CollectionTypeGoods})
	return
}
