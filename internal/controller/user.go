package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"jk-goframe-shop/api/frontend"
	"jk-goframe-shop/internal/consts"
	"jk-goframe-shop/internal/model"
	"jk-goframe-shop/internal/service"
)

// User 内容管理
var User = cUser{}

type cUser struct{}

func (a *cUser) Register(ctx context.Context, req *frontend.RegisterReq) (res *frontend.RegisterRes, err error) {
	data := model.RegisterInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.User().Register(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.RegisterRes{UserId: uint(out.UserId)}, nil
}

func (c *cUser) UpdatePassword(ctx context.Context, req *frontend.UpdatePasswordReq) (res *frontend.UpdatePasswordRes, err error) {
	data := model.UpdatePasswordInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.User().UpdatePassword(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.UpdatePasswordRes{UserId: out.UserId}, nil

}

func (c *cUser) Info(ctx context.Context, req *frontend.UserInfoReq) (res *frontend.UserInfoRes, err error) {
	res = &frontend.UserInfoRes{}
	res.Id = gconv.Uint(ctx.Value(consts.CtxUserId))
	res.Name = gconv.String(ctx.Value(consts.CtxUserName))
	res.Avatar = gconv.String(ctx.Value(consts.CtxUserAvatar))
	res.Sex = gconv.Uint8(ctx.Value(consts.CtxUserSex))
	res.Sign = gconv.String(ctx.Value(consts.CtxUserSign))
	res.Status = gconv.Uint8(ctx.Value(consts.CtxUserId))

	return res, nil
}
