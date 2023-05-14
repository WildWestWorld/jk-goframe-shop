package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"jk-goframe-shop/api/frontend"
	"jk-goframe-shop/internal/model"
	"jk-goframe-shop/internal/service"
)

// Praise 内容管理
var Praise = cPraise{}

type cPraise struct{}

func (a *cPraise) Create(ctx context.Context, req *frontend.PraiseAddReq) (res *frontend.PraiseAddRes, err error) {
	data := model.PraiseAddInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Praise().AddPraise(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.PraiseAddRes{Id: out.Id}, nil
}

func (a *cPraise) Delete(ctx context.Context, req *frontend.PraiseDeleteReq) (res *frontend.PraiseDeleteRes, err error) {
	data := model.PraiseDeleteInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	praise, err := service.Praise().DeletePraise(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.PraiseDeleteRes{Id: praise.Id}, nil
}

// Index article list
func (a *cPraise) List(ctx context.Context, req *frontend.PraiseGetListCommonReq) (res *frontend.PraiseGetListCommonRes, err error) {
	getListRes, err := service.Praise().GetList(ctx, model.PraiseGetListInput{
		Page: req.Page,
		Size: req.Size,
		Type: req.Type,
	})
	if err != nil {
		return nil, err
	}
	return &frontend.PraiseGetListCommonRes{List: getListRes.List, Page: getListRes.Page, Size: getListRes.Size, Total: getListRes.Total}, nil

}
