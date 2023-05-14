package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"jk-goframe-shop/api/frontend"
	"jk-goframe-shop/internal/model"
	"jk-goframe-shop/internal/service"
)

// Collection 内容管理
var Collection = cCollection{}

type cCollection struct{}

func (a *cCollection) Create(ctx context.Context, req *frontend.CollectionAddReq) (res *frontend.CollectionAddRes, err error) {
	data := model.CollectionAddInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Collection().AddCollection(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.CollectionAddRes{Id: out.Id}, nil
}

func (a *cCollection) Delete(ctx context.Context, req *frontend.CollectionDeleteReq) (res *frontend.CollectionDeleteRes, err error) {
	data := model.CollectionDeleteInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	collection, err := service.Collection().DeleteCollection(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.CollectionDeleteRes{Id: collection.Id}, nil
}

// Index article list
func (a *cCollection) List(ctx context.Context, req *frontend.CollectionGetListCommonReq) (res *frontend.CollectionGetListCommonRes, err error) {
	getListRes, err := service.Collection().GetList(ctx, model.CollectionGetListInput{
		Page: req.Page,
		Size: req.Size,
		Type: req.Type,
	})
	if err != nil {
		return nil, err
	}
	return &frontend.CollectionGetListCommonRes{List: getListRes.List, Page: getListRes.Page, Size: getListRes.Size, Total: getListRes.Total}, nil

}
