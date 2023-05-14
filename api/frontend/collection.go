package frontend

import "github.com/gogf/gf/v2/frame/g"

type CollectionAddReq struct {
	g.Meta   `path:"/collection/add" method:"post" tags:"前台收藏" summary:"添加收藏"`
	ObjectId uint  `json:"objectId"  description:"对象id" v:"required#收藏对象id必填"`
	Type     uint8 `json:"type"      description:"收藏类型：1商品 2文章" v:"in:1,2"`
}
type CollectionAddRes struct {
	Id uint `json:"Id"`
}

type CollectionDeleteReq struct {
	g.Meta   `path:"/collection/delete" method:"post" tags:"前台收藏" summary:"添加收藏"`
	Id       uint  `json:"Id"    description:"用户id"`
	ObjectId uint  `json:"objectId"  description:"对象id"`
	Type     uint8 `json:"type"      description:"收藏类型：1商品 2文章" v:"in:1,2"`
}
type CollectionDeleteRes struct {
	Id uint `json:"Id"`
}

// Get列表接口
type CollectionGetListCommonReq struct {
	g.Meta `path:"/collection/list" method:"get" tags:"前台收藏" summary:"收藏列表"`
	Type   uint8 `json:"type" v:"in:0,1,2" dc:"搜藏类型"`
	CommonPaginationReq
}
type CollectionGetListCommonRes struct {
	//todo
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`

	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type CollectionListItem struct {
	Id       int `json:"id"        description:""`
	UserId   int `json:"userId"    description:"用户id"`
	ObjectId int `json:"objectId"  description:"对象id"`
	Type     int `json:"type"      description:"收藏类型：1商品 2文章"`
	//关联查询
	//orm:"with:副type的字段=主type的字段(主type:使用with的type)"
	Goods    interface{} `json:"goods" `
	Articles interface{} `json:"articles" `
}

// 一对一查询
// 使用with 必须使用 orm关联表
// orm:"table:关联表名"
//type GoodsItem struct {
//	g.Meta `orm:"table:good_info"`
//	Id     uint   `json:"id"`
//	Name   string `json:"name"`
//	PicUrl string `json:"pic_url"`
//	Price  int    `json:"price"`
//}
//
//type ArticleItem struct {
//	g.Meta `orm:"table:article_info"`
//	Id     uint   `json:"id"`
//	Title  string `json:"title"`
//	Desc   string `json:"desc"`
//	PicUrl int    `json:"pic_url"`
//}
