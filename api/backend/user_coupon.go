// 用户优惠券
// 注意这里的包已经不是v1的包了，应该换成backend
package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 这里定义的是对外的接口,和一些请求的参数，返回的结构体

// Req 发请求的payload
type UserCouponReq struct {
	g.Meta `path:"/backend/user/coupon/add" tags:"UserCoupon" method:"post" summary:"用户优惠券 添加接口"`
	UserCouponCommonAddUpdate
}

type UserCouponCommonAddUpdate struct {
	UserId   uint  `json:"user_id"  v:"required#用户id为必填"    dc:"用户id"`
	CouponId uint  `json:"coupon_id"   dc:"可用的商品分类,只能指定一类"`
	Status   uint8 `json:"status"   dc:"状态"`
}

// Req 请求完毕后的结果/响应
type UserCouponRes struct {
	//g.Meta `mime:"text/html" example:"string"`
	UserCouponId uint `json:"user_coupon_id"`
}

type UserCouponDeleteReq struct {
	g.Meta `path:"/backend/user/coupon/delete" method:"delete" tags:"内容" summary:"删除内容接口"`
	Id     uint `v:"min:1#请选择需要删除的内容" dc:"内容id"`
}
type UserCouponDeleteRes struct{}

type UserCouponUpdateReq struct {
	g.Meta `path:"/backend/user/coupon/update/{Id}" method:"post" tags:"用户优惠券" summary:"修改用户优惠券接口"`
	Id     uint `json:"id"      v:"min:1#请选择需要修改的用户优惠券Id" dc:"用户优惠券Id"`
	UserCouponCommonAddUpdate
}
type UserCouponUpdateRes struct {
	UserCouponId uint `json:"coupon_id"`
}

// Get列表接口
type UserCouponGetListCommonReq struct {
	g.Meta `path:"/backend/user/coupon/list" method:"get" tags:"用户优惠券" summary:"用户优惠券列表接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}
type UserCouponGetListCommonRes struct {
	//todo
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`

	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

// 返回数据库中的所有内容的Req
type UserCouponGetAllListCommonReq struct {
	g.Meta `path:"/backend/user/coupon/allList" method:"get" tags:"用户优惠券" summary:"用户优惠券全部列表"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	//CommonPaginationReq
}
type UserCouponGetAllListCommonRes struct {
	//todo
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`

	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
