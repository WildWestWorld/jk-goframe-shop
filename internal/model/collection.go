package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type CollectionAddInput struct {
	UserId   uint  `json:"userId"    description:"用户id"`
	ObjectId uint  `json:"objectId"  description:"对象id"`
	Type     uint8 `json:"type"      description:"收藏类型：1商品 2文章"`
}

type CollectionAddOutput struct {
	Id uint `json:"Id"    description:"用户id"`
}

type CollectionDeleteInput struct {
	Id       uint  `json:"Id"    description:"用户id"`
	UserId   uint  `json:"userId"    description:"用户id"`
	ObjectId uint  `json:"objectId"  description:"对象id"`
	Type     uint8 `json:"type"      description:"收藏类型：1商品 2文章"`
}

type CollectionDeleteOutput struct {
	Id uint `json:"Id"    description:"用户id"`
}

// CollectionGetListInput 获取内容列表
type CollectionGetListInput struct {
	Page int   // 分页号码
	Size int   // 分页数量，最大50
	Type uint8 // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// CollectionGetListOutput 查询列表结果
type CollectionGetListOutput struct {
	List  []CollectionGetListOutputItem `json:"list" description:"列表"`
	Page  int                           `json:"page" description:"分页码"`
	Size  int                           `json:"size" description:"分页数量"`
	Total int                           `json:"total" description:"数据总数"`
}

type CollectionGetListOutputItem struct {
	Id       int `json:"id"        description:""`
	UserId   int `json:"userId"    description:"用户id"`
	ObjectId int `json:"objectId"  description:"对象id"`
	Type     int `json:"type"      description:"收藏类型：1商品 2文章"`

	//关联查询
	//orm:"with:副type的字段=主type的字段(主type:使用with的type)"
	Goods    GoodsItem   `json:"goods" orm:"with:id=object_id"`
	Articles ArticleItem `json:"articles" orm:"with:id=object_id"`

	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
}

// 一对一查询
// 使用with 必须使用 orm关联表
// orm:"table:关联表名"
type GoodsItem struct {
	g.Meta `orm:"table:good_info"`
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	PicUrl string `json:"pic_url"`
	Price  int    `json:"price"`
}

type ArticleItem struct {
	g.Meta `orm:"table:article_info"`
	Id     uint   `json:"id"`
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	PicUrl int    `json:"pic_url"`
}

type CheckIsCollectInput struct {
	UserId   uint  `json:"userId"    description:"用户id"`
	ObjectId uint  `json:"objectId"  description:"对象id"`
	Type     uint8 `json:"type"      description:"收藏类型：1商品 2文章"`
}
