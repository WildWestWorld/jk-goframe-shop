// 商品分类
// 注意这里的包已经不是v1的包了，应该换成backend
package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 这里定义的是对外的接口,和一些请求的参数，返回的结构体

// Req 发请求的payload
type CategoryReq struct {
	g.Meta `path:"/backend/category/add" tags:"Category" method:"post" summary:"商品分类 添加接口"`
	CategoryCommonAddUpdate
}

type CategoryCommonAddUpdate struct {
	ParentId uint   `json:"parent_id"     dc:"父级id"`
	Name     string `json:"name"  v:"required#名称必填"    dc:"父级id"`

	PicUrl string `json:"pic_url"    v:"required#图片链接不能为空" dc:"图片链接"`
	Level  uint8  `json:"level" dc:"等级 默认一级分类"`
	Sort   uint8  `json:"sort" dc:"排序"`
}

// Req 请求完毕后的结果/响应
type CategoryRes struct {
	//g.Meta `mime:"text/html" example:"string"`
	CategoryId uint `json:"category_id"`
}

type CategoryDeleteReq struct {
	g.Meta `path:"/backend/category/delete" method:"delete" tags:"内容" summary:"删除内容接口"`
	Id     uint `v:"min:1#请选择需要删除的内容" dc:"内容id"`
}
type CategoryDeleteRes struct{}

type CategoryUpdateReq struct {
	g.Meta `path:"/backend/category/update/{Id}" method:"post" tags:"商品分类" summary:"修改商品分类接口"`
	Id     uint `json:"id"      v:"min:1#请选择需要修改的商品分类Id" dc:"商品分类Id"`
	CategoryCommonAddUpdate
}
type CategoryUpdateRes struct {
	CategoryId uint `json:"category_id"`
}

// Get列表接口
type CategoryGetListCommonReq struct {
	g.Meta `path:"/backend/category/list" method:"get" tags:"商品分类" summary:"商品分类列表接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}
type CategoryGetListCommonRes struct {
	//todo
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`

	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

// 返回数据库中的所有内容的Req
type CategoryGetAllListCommonReq struct {
	g.Meta `path:"/backend/category/allList" method:"get" tags:"商品分类" summary:"商品分类全部列表"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	//CommonPaginationReq
}
type CategoryGetAllListCommonRes struct {
	//todo
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`

	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
