package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"jk-goframe-shop/api/backend"
	"jk-goframe-shop/internal/model"
	"jk-goframe-shop/internal/service"
)

// GoodsOptions 内容管理
var GoodsOptions = cGoodsOptions{}

type cGoodsOptions struct{}

func (a *cGoodsOptions) Create(ctx context.Context, req *backend.GoodsOptionsReq) (res *backend.GoodsOptionsRes, err error) {
	//out, err := service.GoodsOptions().Create(ctx, model.GoodsOptionsCreateInput{
	//	GoodsOptionsCreateUpdateBase: model.GoodsOptionsCreateUpdateBase{
	//		PicUrl: req.PicUrl,
	//		Name:   req.Name,
	//		Status: req.Status,
	//	},
	//})
	data := model.GoodsOptionsCreateInput{}
	err = gconv.Scan(req, &data)
	if err != nil {
		return nil, err
	}

	out, err := service.GoodsOptions().Create(ctx, data)
	if err != nil {
		return nil, err
	}

	return &backend.GoodsOptionsRes{GoodsOptionsId: uint(out.GoodsOptionsId)}, nil
}

func (a *cGoodsOptions) Delete(ctx context.Context, req *backend.GoodsOptionsDeleteReq) (res *backend.GoodsOptionsDeleteRes, err error) {
	err = service.GoodsOptions().Delete(ctx, req.Id)
	return
}

func (a *cGoodsOptions) Update(ctx context.Context, req *backend.GoodsOptionsUpdateReq) (res *backend.GoodsOptionsUpdateRes, err error) {

	data := model.GoodsOptionsUpdateInput{}
	//Struct =性能更好的Scan
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}

	err = service.GoodsOptions().Update(ctx, data)
	return
}

// Index article list
func (a *cGoodsOptions) List(ctx context.Context, req *backend.GoodsOptionsGetListCommonReq) (res *backend.GoodsOptionsGetListCommonRes, err error) {
	getListRes, err := service.GoodsOptions().GetList(ctx, model.GoodsOptionsGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	if err != nil {
		return nil, err
	}
	return &backend.GoodsOptionsGetListCommonRes{List: getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil

}

func (a *cGoodsOptions) AllList(ctx context.Context, req *backend.GoodsOptionsGetAllListCommonReq) (res *backend.GoodsOptionsGetAllListCommonRes, err error) {
	getListRes, err := service.GoodsOptions().GetAllList(ctx, model.GoodsOptionsGetListInput{
		//Page: req.Page,
		//Size: req.Size,
		//Sort: req.Sort,
	})
	if err != nil {
		return nil, err
	}
	return &backend.GoodsOptionsGetAllListCommonRes{List: getListRes.List,
		//Page: getListRes.Page,
		//Size: getListRes.Size,
		Total: getListRes.Total}, nil

}
