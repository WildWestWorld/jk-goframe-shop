package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type PraiseAddInput struct {
	UserId   uint  `json:"userId"    description:"用户id"`
	ObjectId uint  `json:"objectId"  description:"对象id"`
	Type     uint8 `json:"type"      description:"收藏类型：1商品 2文章"`
}

type PraiseAddOutput struct {
	Id uint `json:"Id"    description:"用户id"`
}

type PraiseDeleteInput struct {
	Id       uint  `json:"Id"    description:"用户id"`
	UserId   uint  `json:"userId"    description:"用户id"`
	ObjectId uint  `json:"objectId"  description:"对象id"`
	Type     uint8 `json:"type"      description:"收藏类型：1商品 2文章"`
}

type PraiseDeleteOutput struct {
	Id uint `json:"Id"    description:"用户id"`
}

// PraiseGetListInput 获取内容列表
type PraiseGetListInput struct {
	Page int   // 分页号码
	Size int   // 分页数量，最大50
	Type uint8 // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// PraiseGetListOutput 查询列表结果
type PraiseGetListOutput struct {
	List  []PraiseGetListOutputItem `json:"list" description:"列表"`
	Page  int                       `json:"page" description:"分页码"`
	Size  int                       `json:"size" description:"分页数量"`
	Total int                       `json:"total" description:"数据总数"`
}

type PraiseGetListOutputItem struct {
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
