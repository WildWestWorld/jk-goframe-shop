package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"jk-goframe-shop/internal/model/do"
)

type CommentAddInput struct {
	UserId   uint
	ObjectId uint
	Type     uint8
	ParentId uint
	Content  string
}

type CommentAddOutput struct {
	Id uint `json:"Id"    description:"用户id"`
}

type CommentDeleteInput struct {
	Id       uint
	UserId   uint
	ObjectId uint
	Type     uint8
}

type CommentDeleteOutput struct {
	Id uint
}

// CommentGetListInput 获取内容列表
type CommentGetListInput struct {
	Page int   // 分页号码
	Size int   // 分页数量，最大50
	Type uint8 // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// CommentGetListOutput 查询列表结果
type CommentGetListOutput struct {
	List  []CommentGetListOutputItem `json:"list" description:"列表"`
	Page  int                        `json:"page" description:"分页码"`
	Size  int                        `json:"size" description:"分页数量"`
	Total int                        `json:"total" description:"数据总数"`
}

type CommentGetListOutputItem struct {
	Id       int    `json:"id"        description:""`
	UserId   int    `json:"userId"    description:"用户id"`
	ObjectId int    `json:"objectId"  description:"对象id"`
	Type     int    `json:"type"      description:"收藏类型：1商品 2文章"`
	ParentId uint   `json:"parentId" description:"父级评论id"`
	Content  string `json:"content" description:"评论内容"`
	//关联查询
	//orm:"with:副type的字段=主type的字段(主type:使用with的type)"
	Goods    GoodsItem   `json:"goods" orm:"with:id=object_id"`
	Articles ArticleItem `json:"articles" orm:"with:id=object_id"`

	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
}
type CommentBase struct {
	do.CommentInfo
	User UserInfoBase `json:"user" orm:"with:id=user_id"`
}
