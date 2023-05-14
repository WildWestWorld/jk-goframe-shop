// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"jk-goframe-shop/internal/model"
)

type (
	IVideo interface {
		Create(ctx context.Context, in model.VideoCreateInput) (out model.VideoCreateOutput, err error)
		Delete(ctx context.Context, id uint) error
		Update(ctx context.Context, in model.VideoUpdateInput) error
		GetList(ctx context.Context, in model.VideoGetListInput) (out *model.VideoGetListOutput, err error)
	}
)

var (
	localVideo IVideo
)

func Video() IVideo {
	if localVideo == nil {
		panic("implement not found for interface IVideo, forgot register?")
	}
	return localVideo
}

func RegisterVideo(i IVideo) {
	localVideo = i
}
