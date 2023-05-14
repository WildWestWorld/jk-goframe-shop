// 商品
// 注意这里的包已经不是v1的包了，应该换成backend
package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 这里定义的是对外的接口,和一些请求的参数，返回的结构体

// Req 发请求的payload
type GoodsReq struct {
	g.Meta `path:"/backend/goods/add" tags:"Goods" method:"post" summary:"商品 添加接口"`
	GoodsCommonAddUpdate
}

type GoodsCommonAddUpdate struct {
	Id               int    `json:"id"               description:""`
	PicUrl           string `json:"picUrl"           description:"图片"`
	Name             string `json:"name"            description:"商品名称"  v:"required#名称不能为空"`
	Price            int    `json:"price"          description:"价格 单位分"   v:"required#价格不能为空"  `
	Level1CategoryId int    `json:"level1_category_id" description:"1级分类id"`
	Level2CategoryId int    `json:"level2_category_id" description:"2级分类id"`
	Level3CategoryId int    `json:"level3_category_id" description:"3级分类id"`
	Brand            string `json:"brand"            description:"品牌"  v:"max-length:30#品牌名称最大30个字"`
	Stock            int    `json:"stock"            description:"库存"`
	Sale             int    `json:"sale"             description:"销量"`
	Tags             string `json:"tags"             description:"标签"`
	DetailInfo       string `json:"detail_info"       description:"商品详情"`
}

// Req 请求完毕后的结果/响应
type GoodsRes struct {
	//g.Meta `mime:"text/html" example:"string"`
	GoodsId uint `json:"user_coupon_id"`
}

type GoodsDeleteReq struct {
	g.Meta `path:"/backend/goods/delete" method:"delete" tags:"内容" summary:"删除内容接口"`
	Id     uint `v:"min:1#请选择需要删除的内容" dc:"内容id"`
}
type GoodsDeleteRes struct{}

type GoodsUpdateReq struct {
	g.Meta `path:"/backend/goods/update/{Id}" method:"post" tags:"商品" summary:"修改商品接口"`
	Id     uint `json:"id"      v:"min:1#请选择需要修改的商品Id" dc:"商品Id"`
	GoodsCommonAddUpdate
}
type GoodsUpdateRes struct {
	GoodsId uint `json:"coupon_id"`
}

// Get列表接口
type GoodsGetListCommonReq struct {
	g.Meta `path:"/backend/goods/list" method:"get" tags:"商品" summary:"商品列表接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}
type GoodsGetListCommonRes struct {
	//todo
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`

	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

// 返回数据库中的所有内容的Req
type GoodsGetAllListCommonReq struct {
	g.Meta `path:"/backend/goods/allList" method:"get" tags:"商品" summary:"商品全部列表"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	//CommonPaginationReq
}
type GoodsGetAllListCommonRes struct {
	//todo
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`

	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
