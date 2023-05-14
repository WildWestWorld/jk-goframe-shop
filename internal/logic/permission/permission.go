package perminssion

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/frame/g"
	"jk-goframe-shop/internal/dao"
	"jk-goframe-shop/internal/model"
	"jk-goframe-shop/internal/model/entity"
	"jk-goframe-shop/internal/service"
)

type sPermission struct{}

func init() {
	service.RegisterPermission(New())
}

func New() *sPermission {
	return &sPermission{}
}

// Create 创建内容
func (s *sPermission) Create(ctx context.Context, in model.PermissionCreateInput) (out model.PermissionCreateOutput, err error) {

	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	//InsertAndGetId会返回ID
	lastInsertID, err := dao.PermissionInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.PermissionCreateOutput{PermissionId: int(lastInsertID)}, err
}

// Delete 删除
func (s *sPermission) Delete(ctx context.Context, id uint) error {
	return dao.PermissionInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		// 删除内容(软删除)
		//_, err := dao.PermissionInfo.Ctx(ctx).Where(g.Map{
		//	dao.PermissionInfo.Columns().Id: id,
		//}).Delete()
		// 删除内容(硬删除)Unscoped().Delete() 就会硬删除数据
		//如果没有delete_at 字段 框架默认就会硬删除数据，反之则软删除数据

		_, err := dao.PermissionInfo.Ctx(ctx).Where(g.Map{
			dao.PermissionInfo.Columns().Id: id,
		}).Unscoped().Delete()

		return err
	})
}

// Update 修改
func (s *sPermission) Update(ctx context.Context, in model.PermissionUpdateInput) error {
	return dao.PermissionInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		_, err := dao.PermissionInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.PermissionInfo.Columns().Id).
			Where(dao.PermissionInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}

// GetList 查询内容列表
func (s *sPermission) GetList(ctx context.Context, in model.PermissionGetListInput) (out *model.PermissionGetListOutput, err error) {
	var (
		m = dao.PermissionInfo.Ctx(ctx)
	)
	out = &model.PermissionGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)
	// 排序方式(dec 倒序)
	//listModel = listModel.OrderDesc(dao.PermissionInfo.Columns().Id)

	// 执行查询
	var list []*entity.PermissionInfo
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
	// Permission
	//if err := listModel.ScanList(&out.List, "Permission"); err != nil {
	//	return out, err
	//}
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}

	return
}
