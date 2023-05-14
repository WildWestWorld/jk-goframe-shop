package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"jk-goframe-shop/api/backend"
	"jk-goframe-shop/internal/consts"
	"jk-goframe-shop/internal/model"
	"jk-goframe-shop/internal/service"
)

// Article 内容管理
var Article = cArticle{}

type cArticle struct{}

func (a *cArticle) Create(ctx context.Context, req *backend.ArticleReq) (res *backend.ArticleRes, err error) {
	//out, err := service.Article().Create(ctx, model.ArticleCreateInput{
	//	ArticleCreateUpdateBase: model.ArticleCreateUpdateBase{
	//		PicUrl: req.PicUrl,
	//		Name:   req.Name,
	//		Status: req.Status,
	//	},
	//})
	data := model.ArticleCreateInput{}
	err = gconv.Scan(req, &data)
	if err != nil {
		return nil, err
	}

	out, err := service.Article().Create(ctx, data)
	if err != nil {
		return nil, err
	}

	return &backend.ArticleRes{ArticleId: uint(out.ArticleId)}, nil
}

func (a *cArticle) Delete(ctx context.Context, req *backend.ArticleDeleteReq) (res *backend.ArticleDeleteRes, err error) {
	err = service.Article().Delete(ctx, req.Id)
	return
}

func (a *cArticle) Update(ctx context.Context, req *backend.ArticleUpdateReq) (res *backend.ArticleUpdateRes, err error) {
	data := model.ArticleUpdateInput{}
	//Struct =性能更好的Scan
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	data.UserId = gconv.Int(ctx.Value(consts.CtxAdminId))
	err = service.Article().Update(ctx, data)
	return
}

// Index article list
func (a *cArticle) List(ctx context.Context, req *backend.ArticleGetListCommonReq) (res *backend.ArticleGetListCommonRes, err error) {
	getListRes, err := service.Article().GetList(ctx, model.ArticleGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	if err != nil {
		return nil, err
	}
	return &backend.ArticleGetListCommonRes{List: getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil

}

func (a *cArticle) AllList(ctx context.Context, req *backend.ArticleGetAllListCommonReq) (res *backend.ArticleGetAllListCommonRes, err error) {
	getListRes, err := service.Article().GetAllList(ctx, model.ArticleGetListInput{
		//Page: req.Page,
		//Size: req.Size,
		//Sort: req.Sort,
	})
	if err != nil {
		return nil, err
	}
	return &backend.ArticleGetAllListCommonRes{List: getListRes.List,
		//Page: getListRes.Page,
		//Size: getListRes.Size,
		Total: getListRes.Total}, nil

}
