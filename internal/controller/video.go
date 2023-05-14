package controller

import (
	"context"
	"jk-goframe-shop/api/backend"
	"jk-goframe-shop/internal/model"
	"jk-goframe-shop/internal/service"
)

// Video 内容管理
var Video = cVideo{}

type cVideo struct{}

func (a *cVideo) Create(ctx context.Context, req *backend.VideoReq) (res *backend.VideoRes, err error) {
	out, err := service.Video().Create(ctx, model.VideoCreateInput{
		VideoCreateUpdateBase: model.VideoCreateUpdateBase{
			Type:   req.Type,
			Title:  req.Title,
			Url:    req.Url,
			Sticky: req.Sticky,
			Cover:  req.Cover,

			View:     req.View,
			Duration: req.Duration,
			Favorite: req.Favorite,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.VideoRes{VideoId: uint(out.VideoId)}, nil
}

func (a *cVideo) Delete(ctx context.Context, req *backend.VideoDeleteReq) (res *backend.VideoDeleteRes, err error) {
	err = service.Video().Delete(ctx, req.Id)
	return
}

func (a *cVideo) Update(ctx context.Context, req *backend.VideoUpdateReq) (res *backend.VideoUpdateRes, err error) {
	err = service.Video().Update(ctx, model.VideoUpdateInput{
		Id: req.Id,
		VideoCreateUpdateBase: model.VideoCreateUpdateBase{
			Type:     req.Type,
			Title:    req.Title,
			Url:      req.Url,
			Sticky:   req.Sticky,
			Cover:    req.Cover,
			Favorite: req.Favorite,
		},
	})
	return
}

// Index article list
func (a *cVideo) List(ctx context.Context, req *backend.VideoGetListCommonReq) (res *backend.VideoGetListCommonRes, err error) {
	getListRes, err := service.Video().GetList(ctx, model.VideoGetListInput{
		Page: req.Page,
		Size: req.Size,
		Type: req.Type,
	})
	if err != nil {
		return nil, err
	}
	return &backend.VideoGetListCommonRes{List: getListRes.List, Page: getListRes.Page, Size: getListRes.Size, Total: getListRes.Total}, nil

}
