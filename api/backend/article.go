// 商品
// 注意这里的包已经不是v1的包了，应该换成backend
package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 这里定义的是对外的接口,和一些请求的参数，返回的结构体

// Req 发请求的payload
type ArticleReq struct {
	g.Meta `path:"/backend/goods/add" tags:"Article" method:"post" summary:"商品 添加接口"`
	ArticleCommonAddUpdate
}

type ArticleCommonAddUpdate struct {
	//UserId  int    `json:"user_id"      dc:"用户Id"`
	Title  string `json:"title"           dc:"文章标题" v:"required#文章名称不能为空"`
	Desc   string `json:"desc"            dc:"文章概要" `
	PicUrl string `json:"pic_url"            dc:"图片"`
	//IsAdmin uint   `json:"is_admin" d:"1" description:"是否是管理员发布 1.后台管理员发布 2.前台用户发布"`
	IsAdmin uint   `d:"1" description:"是否是管理员发布 1.后台管理员发布 2.前台用户发布"`
	Detail  string `json:"detail" dc:"文章详情" v:"required#文章详情不能为空"`
	Praise  int    `json:"praise" dc:"点赞数量" `
}

// Req 请求完毕后的结果/响应
type ArticleRes struct {
	//g.Meta `mime:"text/html" example:"string"`
	ArticleId uint `json:"user_coupon_id"`
}

type ArticleDeleteReq struct {
	g.Meta `path:"/backend/goods/delete" method:"delete" tags:"内容" summary:"删除内容接口"`
	Id     uint `v:"min:1#请选择需要删除的内容" dc:"内容id"`
}
type ArticleDeleteRes struct{}

type ArticleUpdateReq struct {
	g.Meta `path:"/backend/goods/update/{Id}" method:"post" tags:"商品" summary:"修改商品接口"`
	Id     uint `json:"id"      v:"min:1#请选择需要修改的商品Id" dc:"商品Id"`
	ArticleCommonAddUpdate
}
type ArticleUpdateRes struct {
	ArticleId uint `json:"coupon_id"`
}

// Get列表接口
type ArticleGetListCommonReq struct {
	g.Meta `path:"/backend/goods/list" method:"get" tags:"商品" summary:"商品列表接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}
type ArticleGetListCommonRes struct {
	//todo
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`

	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

// 返回数据库中的所有内容的Req
type ArticleGetAllListCommonReq struct {
	g.Meta `path:"/backend/goods/allList" method:"get" tags:"商品" summary:"商品全部列表"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	//CommonPaginationReq
}
type ArticleGetAllListCommonRes struct {
	//todo
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`

	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
