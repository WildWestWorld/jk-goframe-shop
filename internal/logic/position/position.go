package Position

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

type sPosition struct{}

func init() {
	service.RegisterPosition(New())
}

func New() *sPosition {
	return &sPosition{}
}

// Create 创建内容
func (s *sPosition) Create(ctx context.Context, in model.PositionCreateInput) (out model.PositionCreateOutput, err error) {

	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.PositionInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.PositionCreateOutput{PositionId: int(lastInsertID)}, err
}

// Delete 删除
func (s *sPosition) Delete(ctx context.Context, id uint) error {
	return dao.PositionInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		// 删除内容(软删除)
		//_, err := dao.PositionInfo.Ctx(ctx).Where(g.Map{
		//	dao.PositionInfo.Columns().Id: id,
		//}).Delete()
		// 删除内容(硬删除)Unscoped().Delete() 就会硬删除数据
		//如果没有delete_at 字段 框架默认就会硬删除数据，反之则软删除数据

		_, err := dao.PositionInfo.Ctx(ctx).Where(g.Map{
			dao.PositionInfo.Columns().Id: id,
		}).Delete()

		return err
	})
}

// Update 修改
func (s *sPosition) Update(ctx context.Context, in model.PositionUpdateInput) error {
	return dao.PositionInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		_, err := dao.PositionInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.PositionInfo.Columns().Id).
			Where(dao.PositionInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}

// GetList 查询内容列表
func (s *sPosition) GetList(ctx context.Context, in model.PositionGetListInput) (out *model.PositionGetListOutput, err error) {
	var (
		m = dao.PositionInfo.Ctx(ctx)
	)
	out = &model.PositionGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)
	// 排序方式(dec 倒序)
	listModel = listModel.OrderDesc(dao.PositionInfo.Columns().Id)

	// 执行查询
	var list []*entity.PositionInfo
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
	// Position
	//if err := listModel.ScanList(&out.List, "Position"); err != nil {
	//	return out, err
	//}
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}

	return
}
