package model

import (
	"github.com/gogf/gf/v2/util/gmeta"
	"jk-goframe-shop/internal/model/do"
	"jk-goframe-shop/internal/model/entity"
)

// GoodsCreateUpdateBase 创建/修改内容基类
type GoodsCreateUpdateBase struct {
	PicUrl           string
	Name             string
	Price            int
	Level1CategoryId int
	Level2CategoryId int
	Level3CategoryId int
	Brand            string
	Stock            int
	Sale             int
	Tags             string
	DetailInfo       string
}

// GoodsCreateInput 创建内容
type GoodsCreateInput struct {
	GoodsCreateUpdateBase
}

// GoodsCreateOutput 创建内容返回结果
type GoodsCreateOutput struct {
	GoodsId int `json:"user_coupon_id"`
}

// GoodsUpdateInput 修改内容
type GoodsUpdateInput struct {
	GoodsCreateUpdateBase
	Id uint
}

// GoodsGetListInput 获取内容列表
type GoodsGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// GoodsGetListOutput 查询列表结果
type GoodsGetListOutput struct {
	List  []GoodsGetListOutputItem `json:"list" description:"列表"`
	Page  int                      `json:"page" description:"分页码"`
	Size  int                      `json:"size" description:"分页数量"`
	Total int                      `json:"total" description:"数据总数"`
}

// GoodsSearchInput 搜索列表
type GoodsSearchInput struct {
	Key     string // 关键字
	Type    string // 内容模型
	GoodsId uint   // 栏目ID
	Page    int    // 分页号码
	Size    int    // 分页数量，最大50
	Sort    int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// GoodsSearchOutput 搜索列表结果
type GoodsSearchOutput struct {
	List  []GoodsSearchOutputItem `json:"list"`  // 列表
	Stats map[string]int          `json:"stats"` // 搜索统计
	Page  int                     `json:"page"`  // 分页码
	Size  int                     `json:"size"`  // 分页数量
	Total int                     `json:"total"` // 数据总数
}

type GoodsGetListOutputItem struct {
	entity.GoodsInfo
	//Id       uint  `json:"id"`
	//UserId   uint  `json:"user_id"`
	//CouponId uint  `json:"coupon_id"`
	//Status   uint8 `json:"status"`
	//
	//CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	//UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间

}

type GoodsSearchOutputItem struct {
	GoodsGetListOutputItem
}

// GoodsListItem 主要用于列表展示
//
//	type GoodsListItem struct {
//		Id        uint        `json:"id"` // 自增ID
//		PicUrl    string      `json:"pic_url"`
//		Link      string      `json:"link"`
//		Sort      uint        `json:"sort"`       // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
//		CreatedAt *gtime.Time `json:"created_at"` // 创建时间
//		UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
//	}
type GoodsDetailInput struct {
	Id uint
}
type GoodsDetailOutput struct {
	do.GoodsInfo
	Options []*do.GoodsOptionsInfo `orm:"with:goods_id=id"`
	//with 搜索的后面 的空格必加 不然无法正确查询数据
	Comments []*CommentBase `orm:"with:object_id=id, where:type=1"`
	//IsCollect bool
}

type BaseGoodsColumns struct {
	gmeta.Meta `orm:"table:goods_info"`
	Id         string `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Brand      string `json:"brand"`
	Tags       string `json:"tags"`
	PicUrl     string `json:"pic_url"`
	DetailInfo string `json:"detail_info"`
}
