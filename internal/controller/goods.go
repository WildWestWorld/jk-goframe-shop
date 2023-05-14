package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"jk-goframe-shop/api/backend"
	"jk-goframe-shop/api/frontend"
	"jk-goframe-shop/internal/model"
	"jk-goframe-shop/internal/service"
)

// Goods 内容管理
var Goods = cGoods{}

type cGoods struct{}

func (a *cGoods) Create(ctx context.Context, req *backend.GoodsReq) (res *backend.GoodsRes, err error) {
	//out, err := service.Goods().Create(ctx, model.GoodsCreateInput{
	//	GoodsCreateUpdateBase: model.GoodsCreateUpdateBase{
	//		PicUrl: req.PicUrl,
	//		Name:   req.Name,
	//		Status: req.Status,
	//	},
	//})
	data := model.GoodsCreateInput{}
	err = gconv.Scan(req, &data)
	if err != nil {
		return nil, err
	}

	out, err := service.Goods().Create(ctx, data)
	if err != nil {
		return nil, err
	}

	return &backend.GoodsRes{GoodsId: uint(out.GoodsId)}, nil
}

func (a *cGoods) Delete(ctx context.Context, req *backend.GoodsDeleteReq) (res *backend.GoodsDeleteRes, err error) {
	err = service.Goods().Delete(ctx, req.Id)
	return
}

func (a *cGoods) Update(ctx context.Context, req *backend.GoodsUpdateReq) (res *backend.GoodsUpdateRes, err error) {

	data := model.GoodsUpdateInput{}
	//Struct =性能更好的Scan
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}

	err = service.Goods().Update(ctx, data)
	return
}

// Index article list
func (a *cGoods) List(ctx context.Context, req *backend.GoodsGetListCommonReq) (res *backend.GoodsGetListCommonRes, err error) {
	getListRes, err := service.Goods().GetList(ctx, model.GoodsGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	if err != nil {
		return nil, err
	}
	return &backend.GoodsGetListCommonRes{List: getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil

}

func (a *cGoods) AllList(ctx context.Context, req *backend.GoodsGetAllListCommonReq) (res *backend.GoodsGetAllListCommonRes, err error) {
	getListRes, err := service.Goods().GetAllList(ctx, model.GoodsGetListInput{
		//Page: req.Page,
		//Size: req.Size,
		//Sort: req.Sort,
	})
	if err != nil {
		return nil, err
	}
	return &backend.GoodsGetAllListCommonRes{List: getListRes.List,
		//Page: getListRes.Page,
		//Size: getListRes.Size,
		Total: getListRes.Total}, nil

}

func (s *cGoods) Detail(ctx context.Context, req *frontend.GoodsDetailReq) (res *frontend.GoodsDetailRes, err error) {
	detail, err := service.Goods().GoodsDetail(ctx, model.GoodsDetailInput{Id: req.Id})
	if err != nil {
		return nil, err
	}
	res = &frontend.GoodsDetailRes{}
	err = gconv.Struct(detail, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
