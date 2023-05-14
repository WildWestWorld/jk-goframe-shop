package frontend

import "github.com/gogf/gf/v2/frame/g"

type CommentAddReq struct {
	g.Meta   `path:"/comment/add" method:"post" tags:"前台收藏" summary:"添加收藏"`
	UserId   uint   `json:"userId"    description:"用户id"`
	ObjectId uint   `json:"objectId"  description:"对象id" v:"required#评论对象id必填"`
	Type     uint8  `json:"type"      description:"收藏类型：1商品 2文章" v:"in:1,2"`
	ParentId uint   `json:"parentId" description:"父级评论id"`
	Content  string `json:"content" v:"required#评论必填"`
}

type CommentAddRes struct {
	Id uint `json:"Id"`
}

type CommentDeleteReq struct {
	g.Meta `path:"/comment/delete" method:"post" tags:"前台收藏" summary:"添加收藏"`
	Id     uint `json:"Id"    description:"用户id"`
}
type CommentDeleteRes struct {
	Id uint `json:"Id"`
}

// Get列表接口
type CommentGetListCommonReq struct {
	g.Meta `path:"/comment/list" method:"get" tags:"前台收藏" summary:"收藏列表"`
	Type   uint8 `json:"type" v:"in:0,1,2" dc:"搜藏类型"`
	CommonPaginationReq
}
type CommentGetListCommonRes struct {
	//todo
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`

	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type CommentListItem struct {
	Id       int `json:"id"        description:""`
	UserId   int `json:"userId"    description:"用户id"`
	ObjectId int `json:"objectId"  description:"对象id"`
	Type     int `json:"type"      description:"收藏类型：1商品 2文章"`
	//关联查询
	//orm:"with:副type的字段=主type的字段(主type:使用with的type)"
	Goods    interface{} `json:"goods" `
	Articles interface{} `json:"articles" `
}

type CommentBase struct {
	Id       int          `json:"id"        description:""`
	ParentId int          `json:"parentId"  description:"父级评论id"`
	UserId   int          `json:"userId"    description:""`
	User     UserInfoBase `json:"user" dc:"用户信息"`
	ObjectId int          `json:"objectId"  description:""`
	Type     int          `json:"type"      description:"评论类型：1商品 2文章"`
	Content  string       `json:"content"   description:"评论内容"`
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
