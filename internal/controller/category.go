package controller

import (
	"context"
	"jk-goframe-shop/api/backend"
	"jk-goframe-shop/internal/model"
	"jk-goframe-shop/internal/service"
)

// Category 内容管理
var Category = cCategory{}

type cCategory struct{}

func (a *cCategory) Create(ctx context.Context, req *backend.CategoryReq) (res *backend.CategoryRes, err error) {
	out, err := service.Category().Create(ctx, model.CategoryCreateInput{
		CategoryCreateUpdateBase: model.CategoryCreateUpdateBase{
			ParentId: req.ParentId,
			PicUrl:   req.PicUrl,
			Name:     req.Name,
			Sort:     req.Sort,
			Level:    req.Level,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.CategoryRes{CategoryId: uint(out.CategoryId)}, nil
}

func (a *cCategory) Delete(ctx context.Context, req *backend.CategoryDeleteReq) (res *backend.CategoryDeleteRes, err error) {
	err = service.Category().Delete(ctx, req.Id)
	return
}

func (a *cCategory) Update(ctx context.Context, req *backend.CategoryUpdateReq) (res *backend.CategoryUpdateRes, err error) {
	err = service.Category().Update(ctx, model.CategoryUpdateInput{
		Id: req.Id,
		CategoryCreateUpdateBase: model.CategoryCreateUpdateBase{
			ParentId: req.ParentId,
			PicUrl:   req.PicUrl,
			Name:     req.Name,
			Sort:     req.Sort,
			Level:    req.Level,
		},
	})
	return
}

// Index article list
func (a *cCategory) List(ctx context.Context, req *backend.CategoryGetListCommonReq) (res *backend.CategoryGetListCommonRes, err error) {
	getListRes, err := service.Category().GetList(ctx, model.CategoryGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	if err != nil {
		return nil, err
	}
	return &backend.CategoryGetListCommonRes{List: getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil

}

func (a *cCategory) AllList(ctx context.Context, req *backend.CategoryGetAllListCommonReq) (res *backend.CategoryGetAllListCommonRes, err error) {
	getListRes, err := service.Category().GetAllList(ctx, model.CategoryGetListInput{
		//Page: req.Page,
		//Size: req.Size,
		//Sort: req.Sort,
	})
	if err != nil {
		return nil, err
	}
	return &backend.CategoryGetAllListCommonRes{List: getListRes.List,
		//Page: getListRes.Page,
		//Size: getListRes.Size,
		Total: getListRes.Total}, nil

}
