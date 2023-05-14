package Video

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

type sVideo struct{}

func init() {
	service.RegisterVideo(New())
}

func New() *sVideo {
	return &sVideo{}
}

// Create 创建内容
func (s *sVideo) Create(ctx context.Context, in model.VideoCreateInput) (out model.VideoCreateOutput, err error) {

	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}

	//InsertAndGetId 方法就是 插入后返回对应的id
	lastInsertID, err := dao.VideoInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.VideoCreateOutput{VideoId: int(lastInsertID)}, err
}

// Delete 删除
func (s *sVideo) Delete(ctx context.Context, id uint) error {
	return dao.VideoInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		// 删除内容(软删除)
		//_, err := dao.VideoInfo.Ctx(ctx).Where(g.Map{
		//	dao.VideoInfo.Columns().Id: id,
		//}).Delete()
		// 删除内容(硬删除)Unscoped().Delete() 就会硬删除数据
		//如果没有delete_at 字段 框架默认就会硬删除数据，反之则软删除数据

		_, err := dao.VideoInfo.Ctx(ctx).Where(g.Map{
			dao.VideoInfo.Columns().Id: id,
		}).Unscoped().Delete()

		return err
	})
}

// Update 修改
func (s *sVideo) Update(ctx context.Context, in model.VideoUpdateInput) error {
	return dao.VideoInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		_, err := dao.VideoInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.VideoInfo.Columns().Id).
			Where(dao.VideoInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}

// GetList 查询内容列表
func (s *sVideo) GetList(ctx context.Context, in model.VideoGetListInput) (out *model.VideoGetListOutput, err error) {
	var (
		m = dao.VideoInfo.Ctx(ctx)
	)
	out = &model.VideoGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)
	g.Dump(in.Type)
	// 排序方式(dec 倒序)
	listModel = listModel.WhereIn("type", in.Type).OrderDesc(dao.VideoInfo.Columns().Id)

	// 执行查询
	var list []*entity.VideoInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count("type", in.Type)
	//out.Total, err = m.Where("type", in.Type).Count()

	//out.Total = len(list)

	if err != nil {
		return out, err
	}
	// Video
	//if err := listModel.ScanList(&out.List, "Video"); err != nil {
	//	return out, err
	//}
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}

	return
}
