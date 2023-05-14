// 轮播图
// 注意这里的包已经不是v1的包了，应该换成backend
package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 这里定义的是对外的接口,和一些请求的参数，返回的结构体

// Req 发请求的payload
type RotationReq struct {
	g.Meta `path:"/rotation/add" tags:"Rotation" method:"post" summary:"轮播图 添加接口"`
	PicUrl string `json:"pic_url"    v:"required#图片链接不能为空" dc:"图片链接"`
	Link   string `json:"link"    v:"required#跳转链接不能为空"     dc:"跳转链接"`
	Sort   int    `json:"sort"       dc:"排序"`
}

// Req 请求完毕后的结果/响应
type RotationRes struct {
	//g.Meta `mime:"text/html" example:"string"`
	RotationId uint `json:"rotation_id"`
}

type RotationDeleteReq struct {
	g.Meta `path:"/rotation/delete" method:"delete" tags:"内容" summary:"删除内容接口"`
	Id     uint `v:"min:1#请选择需要删除的内容" dc:"内容id"`
}
type RotationDeleteRes struct{}

type RotationUpdateReq struct {
	g.Meta `path:"/rotation/update/{Id}" method:"post" tags:"轮播图" summary:"修改轮播图接口"`
	Id     uint   `json:"id"      v:"min:1#请选择需要修改的轮播图Id" dc:"轮播图Id"`
	PicUrl string `json:"pic_url"    v:"required#轮播图图片不能为空" dc:"轮播图图片"`
	Link   string `json:"link"    v:"required#跳转链接不能为空"     dc:"跳转链接"`
	Sort   int    `json:"sort"       dc:"排序"`
}
type RotationUpdateRes struct {
	RotationId uint `json:"rotation_id"`
}

// Get列表接口
type RotationGetListCommonReq struct {
	g.Meta `path:"/rotation/list" method:"get" tags:"轮播图" summary:"轮播图列表接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}
type RotationGetListCommonRes struct {
	//todo
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`

	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
