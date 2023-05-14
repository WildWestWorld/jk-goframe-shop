// 手工位图
// 注意这里的包已经不是v1的包了，应该换成backend
package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 这里定义的是对外的接口,和一些请求的参数，返回的结构体

// Req 发请求的payload
type PositionReq struct {
	g.Meta    `path:"/backend/position/add" tags:"Position" method:"post" summary:"手工位图 添加接口"`
	PicUrl    string `json:"pic_url"    v:"required#图片链接不能为空" dc:"图片链接"`
	Link      string `json:"link"    v:"required#跳转链接不能为空"     dc:"跳转链接"`
	GoodsName string `json:"goods_name" v:"required#商品名字不能为空" dc:"商品名称"`
	GoodsId   uint   `json:"goods_id" v:"required#商品Id不能为空" dc:"商品Id"`

	Sort int `json:"sort"       dc:"排序"`
}

// Req 请求完毕后的结果/响应
type PositionRes struct {
	//g.Meta `mime:"text/html" example:"string"`
	PositionId uint `json:"position_id"`
}

type PositionDeleteReq struct {
	g.Meta `path:"/backend/position/delete" method:"delete" tags:"内容" summary:"删除内容接口"`
	Id     uint `v:"min:1#请选择需要删除的内容" dc:"内容id"`
}
type PositionDeleteRes struct{}

type PositionUpdateReq struct {
	g.Meta    `path:"/backend/position/update/{Id}" method:"post" tags:"手工位图" summary:"修改手工位图接口"`
	Id        uint   `json:"id"      v:"min:1#请选择需要修改的手工位图Id" dc:"手工位图Id"`
	PicUrl    string `json:"pic_url"    v:"required#手工位图图片不能为空" dc:"手工位图图片"`
	Link      string `json:"link"    v:"required#跳转链接不能为空"     dc:"跳转链接"`
	Sort      int    `json:"sort"       dc:"排序"`
	GoodsName string `json:"goods_name" v:"required#商品名字不能为空" dc:"商品名称"`
	GoodsId   uint   `json:"goods_id" v:"required#商品Id不能为空" dc:"商品Id"`
}
type PositionUpdateRes struct {
	PositionId uint `json:"position_id"`
}

// Get列表接口
type PositionGetListCommonReq struct {
	g.Meta `path:"/backend/position/list" method:"get" tags:"手工位图" summary:"手工位图列表接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}
type PositionGetListCommonRes struct {
	//todo
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`

	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
