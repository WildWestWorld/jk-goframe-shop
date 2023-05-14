package comment

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"jk-goframe-shop/internal/consts"
	"jk-goframe-shop/internal/dao"
	"jk-goframe-shop/internal/model"
	"jk-goframe-shop/internal/service"
)

type sComment struct{}

func init() {
	service.RegisterComment(New())
}

func New() *sComment {
	return &sComment{}
}

func (s *sComment) AddComment(ctx context.Context, in model.CommentAddInput) (res *model.CommentAddOutput, err error) {
	in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
	id, err := dao.CommentInfo.Ctx(ctx).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &model.CommentAddOutput{Id: gconv.Uint(id)}, nil
}

// 兼容处理优先根据收藏id删除，再根据对象id和type 删除
func (s *sComment) DeleteComment(ctx context.Context, in model.CommentDeleteInput) (res *model.CommentDeleteOutput, err error) {

	//删除评论值允许删除自己的
	condition := g.Map{
		dao.CommentInfo.Columns().Id:     in.Id,
		dao.CommentInfo.Columns().UserId: ctx.Value(consts.CtxUserId),
	}
	_, err = dao.CommentInfo.Ctx(ctx).Where(condition).Delete()
	if err != nil {
		return nil, err
	}
	return &model.CommentDeleteOutput{Id: gconv.Uint(in.Id)}, nil

}

// GetList 查询内容列表
func (s *sComment) GetList(ctx context.Context, in model.CommentGetListInput) (out *model.CommentGetListOutput, err error) {
	var m = dao.CommentInfo.Ctx(ctx)

	out = &model.CommentGetListOutput{
		Page: in.Page,
		Size: in.Size,
		//CommentGetListOutputItem
		//这样写没有值的时候就不会返回Null，而是返回【】
		List: []model.CommentGetListOutputItem{},
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)

	//条件查询
	if in.Type != 0 {
		//查询 Where(查询的字段，查询字段的内容)
		listModel.Where(dao.CommentInfo.Columns().Type, in.Type)
	}
	// 排序方式(dec 倒序)
	//listModel = listModel.OrderDesc(dao.CommentInfo.Columns().Id)

	//// 执行查询
	//var list []*entity.CommentInfo
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

	if in.Type == consts.CommentTypeGoods {
		err := listModel.With(model.GoodsItem{}).Scan(&out.List)
		if err != nil {
			return out, err
		}
	} else if in.Type == consts.CommentTypeArticle {
		err := listModel.With(model.ArticleItem{}).Scan(&out.List)
		if err != nil {
			return out, err
		}
	} else {
		if err := listModel.WithAll().Scan(&out.List); err != nil {
			return out, err
		}

	}

	// Comment
	//if err := listModel.ScanList(&out.List, "Comment"); err != nil {
	//	return out, err
	//}
	//if err := listModel.WithAll().Scan(&out.List); err != nil {
	//	return out, err
	//}

	return
}

// 抽取 获取收藏数量
func CommentCount(ctx context.Context, objectId uint, commentType uint8) (count int, err error) {
	//写查询条件
	condition := g.Map{
		dao.CommentInfo.Columns().ObjectId: objectId,
		dao.CommentInfo.Columns().Type:     commentType,
	}
	count, err = dao.CommentInfo.Ctx(ctx).Where(condition).Count()
	if err != nil {
		return count, err
	}
	return
}

// 判断用户是否收藏
func CheckIsComment(ctx context.Context, in model.CheckIsCollectInput) (bool, error) {
	//写查询条件
	condition := g.Map{
		dao.CommentInfo.Columns().UserId:   ctx.Value(consts.CtxUserId),
		dao.CommentInfo.Columns().ObjectId: in.ObjectId,
		dao.CommentInfo.Columns().Type:     in.Type,
	}
	count, err := dao.CommentInfo.Ctx(ctx).Where(condition).Count()
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
