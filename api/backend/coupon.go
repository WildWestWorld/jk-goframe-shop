// 优惠券
// 注意这里的包已经不是v1的包了，应该换成backend
package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 这里定义的是对外的接口,和一些请求的参数，返回的结构体

// Req 发请求的payload
type CouponReq struct {
	g.Meta `path:"/backend/coupon/add" tags:"Coupon" method:"post" summary:"优惠券 添加接口"`
	CouponCommonAddUpdate
}

type CouponCommonAddUpdate struct {
	//ParentId uint   `json:"parent_id"     dc:"父级id"`
	Name  string `json:"name"  v:"required#名称必填"    dc:"父级id"`
	Price int    `json:"price"  v:"required#优惠券价格为必填"    dc:"优惠券价格"`

	GoodsIds   string `json:"goods_ids"   dc:"可用的商品id 可用多个,使用逗号分割"`
	CategoryId int    `json:"category_id"   dc:"可用的商品分类,只能指定一类"`
}

// Req 请求完毕后的结果/响应
type CouponRes struct {
	//g.Meta `mime:"text/html" example:"string"`
	CouponId uint `json:"coupon_id"`
}

type CouponDeleteReq struct {
	g.Meta `path:"/backend/coupon/delete" method:"delete" tags:"内容" summary:"删除内容接口"`
	Id     uint `v:"min:1#请选择需要删除的内容" dc:"内容id"`
}
type CouponDeleteRes struct{}

type CouponUpdateReq struct {
	g.Meta `path:"/backend/coupon/update/{Id}" method:"post" tags:"优惠券" summary:"修改优惠券接口"`
	Id     uint `json:"id"      v:"min:1#请选择需要修改的优惠券Id" dc:"优惠券Id"`
	CouponCommonAddUpdate
}
type CouponUpdateRes struct {
	CouponId uint `json:"coupon_id"`
}

// Get列表接口
type CouponGetListCommonReq struct {
	g.Meta `path:"/backend/coupon/list" method:"get" tags:"优惠券" summary:"优惠券列表接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}
type CouponGetListCommonRes struct {
	//todo
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`

	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

// 返回数据库中的所有内容的Req
type CouponGetAllListCommonReq struct {
	g.Meta `path:"/backend/coupon/allList" method:"get" tags:"优惠券" summary:"优惠券全部列表"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	//CommonPaginationReq
}
type CouponGetAllListCommonRes struct {
	//todo
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`

	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
