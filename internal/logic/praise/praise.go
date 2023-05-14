package praise

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"jk-goframe-shop/internal/consts"
	"jk-goframe-shop/internal/dao"
	"jk-goframe-shop/internal/model"
	"jk-goframe-shop/internal/service"
)

type sPraise struct{}

func init() {
	service.RegisterPraise(New())
}

func New() *sPraise {
	return &sPraise{}
}

func (s *sPraise) AddPraise(ctx context.Context, in model.PraiseAddInput) (res *model.PraiseAddOutput, err error) {
	in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
	id, err := dao.PraiseInfo.Ctx(ctx).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &model.PraiseAddOutput{Id: gconv.Uint(id)}, nil
}

// 兼容处理优先根据收藏id删除，再根据对象id和type 删除
func (s *sPraise) DeletePraise(ctx context.Context, in model.PraiseDeleteInput) (res *model.PraiseDeleteOutput, err error) {
	//	如果有收藏id就根据搜藏id删除
	//如果不是就根据对象id和type删除

	if in.Id != 0 {
		_, err := dao.PraiseInfo.Ctx(ctx).WherePri(in.Id).Delete()
		if err != nil {
			return nil, err
		}
		return &model.PraiseDeleteOutput{Id: gconv.Uint(in.Id)}, nil
	} else {
		in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
		id, err := dao.PraiseInfo.Ctx(ctx).OmitEmpty().Where(in).Delete()
		if err != nil {
			return nil, err
		}
		return &model.PraiseDeleteOutput{Id: gconv.Uint(id)}, nil
	}

}

// GetList 查询内容列表
func (s *sPraise) GetList(ctx context.Context, in model.PraiseGetListInput) (out *model.PraiseGetListOutput, err error) {
	var m = dao.PraiseInfo.Ctx(ctx)

	out = &model.PraiseGetListOutput{
		Page: in.Page,
		Size: in.Size,
		//PraiseGetListOutputItem
		//这样写没有值的时候就不会返回Null，而是返回【】
		List: []model.PraiseGetListOutputItem{},
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)

	//条件查询
	if in.Type != 0 {
		//查询 Where(查询的字段，查询字段的内容)
		listModel.Where(dao.PraiseInfo.Columns().Type, in.Type)
	}
	// 排序方式(dec 倒序)
	//listModel = listModel.OrderDesc(dao.PraiseInfo.Columns().Id)

	//// 执行查询
	//var list []*entity.PraiseInfo
	////withAll 使用静态关联查询
	//if err := listModel.WithAll().Scan(&list); err != nil {
	//	return out, err
	//}
	//// 没有数据
	//if len(list) == 0 {
	//	return out, nil
	//}
	out.Total, err = listModel.Count()
	if err != nil {
		return out, err
	}
	if out.Total == 0 {
		return out, err
	}

	if in.Type == consts.PraiseTypeGoods {
		err := listModel.With(model.GoodsItem{}).Scan(&out.List)
		if err != nil {
			return out, err
		}
	} else if in.Type == consts.PraiseTypeArticle {
		err := listModel.With(model.ArticleItem{}).Scan(&out.List)
		if err != nil {
			return out, err
		}
	} else {
		if err := listModel.WithAll().Scan(&out.List); err != nil {
			return out, err
		}

	}

	// Praise
	//if err := listModel.ScanList(&out.List, "Praise"); err != nil {
	//	return out, err
	//}
	//if err := listModel.WithAll().Scan(&out.List); err != nil {
	//	return out, err
	//}

	return
}

// 抽取 获取收藏数量
func PraiseCount(ctx context.Context, objectId uint, praiseType uint8) (count int, err error) {
	//写查询条件
	condition := g.Map{
		dao.PraiseInfo.Columns().ObjectId: objectId,
		dao.PraiseInfo.Columns().Type:     praiseType,
	}
	count, err = dao.PraiseInfo.Ctx(ctx).Where(condition).Count()
	if err != nil {
		return count, err
	}
	return
}

// 判断用户是否收藏
func CheckIsPraise(ctx context.Context, in model.CheckIsCollectInput) (bool, error) {
	//写查询条件
	condition := g.Map{
		dao.PraiseInfo.Columns().UserId:   ctx.Value(consts.CtxUserId),
		dao.PraiseInfo.Columns().ObjectId: in.ObjectId,
		dao.PraiseInfo.Columns().Type:     in.Type,
	}
	count, err := dao.PraiseInfo.Ctx(ctx).Where(condition).Count()
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
