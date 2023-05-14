package role

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

type sRole struct{}

func init() {
	service.RegisterRole(New())
}

func New() *sRole {
	return &sRole{}
}

// Create 创建内容
func (s *sRole) Create(ctx context.Context, in model.RoleCreateInput) (out model.RoleCreateOutput, err error) {

	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	//InsertAndGetId会返回ID
	lastInsertID, err := dao.RoleInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RoleCreateOutput{RoleId: int(lastInsertID)}, err
}

// Delete 删除
func (s *sRole) Delete(ctx context.Context, id uint) error {
	return dao.RoleInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		// 删除内容(软删除)
		//_, err := dao.RoleInfo.Ctx(ctx).Where(g.Map{
		//	dao.RoleInfo.Columns().Id: id,
		//}).Delete()
		// 删除内容(硬删除)Unscoped().Delete() 就会硬删除数据
		//如果没有delete_at 字段 框架默认就会硬删除数据，反之则软删除数据

		_, err := dao.RoleInfo.Ctx(ctx).Where(g.Map{
			dao.RoleInfo.Columns().Id: id,
		}).Unscoped().Delete()

		return err
	})
}

// Update 修改
func (s *sRole) Update(ctx context.Context, in model.RoleUpdateInput) error {
	return dao.RoleInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		_, err := dao.RoleInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.RoleInfo.Columns().Id).
			Where(dao.RoleInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}

// GetList 查询内容列表
func (s *sRole) GetList(ctx context.Context, in model.RoleGetListInput) (out *model.RoleGetListOutput, err error) {
	var (
		m = dao.RoleInfo.Ctx(ctx)
	)
	out = &model.RoleGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)
	// 排序方式(dec 倒序)
	//listModel = listModel.OrderDesc(dao.RoleInfo.Columns().Id)

	// 执行查询
	var list []*entity.RoleInfo
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
	// Role
	//if err := listModel.ScanList(&out.List, "Role"); err != nil {
	//	return out, err
	//}
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}

	return
}
