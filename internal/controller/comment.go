package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"jk-goframe-shop/api/frontend"
	"jk-goframe-shop/internal/model"
	"jk-goframe-shop/internal/service"
)

// Comment 内容管理
var Comment = cComment{}

type cComment struct{}

func (a *cComment) Create(ctx context.Context, req *frontend.CommentAddReq) (res *frontend.CommentAddRes, err error) {
	data := model.CommentAddInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Comment().AddComment(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.CommentAddRes{Id: out.Id}, nil
}

func (a *cComment) Delete(ctx context.Context, req *frontend.CommentDeleteReq) (res *frontend.CommentDeleteRes, err error) {
	data := model.CommentDeleteInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	comment, err := service.Comment().DeleteComment(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.CommentDeleteRes{Id: comment.Id}, nil
}

// Index article list
func (a *cComment) List(ctx context.Context, req *frontend.CommentGetListCommonReq) (res *frontend.CommentGetListCommonRes, err error) {
	getListRes, err := service.Comment().GetList(ctx, model.CommentGetListInput{
		Page: req.Page,
		Size: req.Size,
		Type: req.Type,
	})
	if err != nil {
		return nil, err
	}
	return &frontend.CommentGetListCommonRes{List: getListRes.List, Page: getListRes.Page, Size: getListRes.Size, Total: getListRes.Total}, nil

}
