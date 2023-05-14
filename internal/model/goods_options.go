package model

import "jk-goframe-shop/internal/model/entity"

// GoodsOptionsCreateUpdateBase 创建/修改内容基类
type GoodsOptionsCreateUpdateBase struct {
	GoodsId    uint
	PicUrl     string
	Name       string
	Price      uint
	Stock      int
	Sale       int
	Tags       string
	DetailInfo string
}

// GoodsOptionsCreateInput 创建内容
type GoodsOptionsCreateInput struct {
	GoodsOptionsCreateUpdateBase
}

// GoodsOptionsCreateOutput 创建内容返回结果
type GoodsOptionsCreateOutput struct {
	GoodsOptionsId int `json:"user_coupon_id"`
}

// GoodsOptionsUpdateInput 修改内容
type GoodsOptionsUpdateInput struct {
	GoodsOptionsCreateUpdateBase
	Id uint
}

// GoodsOptionsGetListInput 获取内容列表
type GoodsOptionsGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// GoodsOptionsGetListOutput 查询列表结果
type GoodsOptionsGetListOutput struct {
	List  []GoodsOptionsGetListOutputItem `json:"list" description:"列表"`
	Page  int                             `json:"page" description:"分页码"`
	Size  int                             `json:"size" description:"分页数量"`
	Total int                             `json:"total" description:"数据总数"`
}

// GoodsOptionsSearchInput 搜索列表
type GoodsOptionsSearchInput struct {
	Key            string // 关键字
	Type           string // 内容模型
	GoodsOptionsId uint   // 栏目ID
	Page           int    // 分页号码
	Size           int    // 分页数量，最大50
	Sort           int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// GoodsOptionsSearchOutput 搜索列表结果
type GoodsOptionsSearchOutput struct {
	List  []GoodsOptionsSearchOutputItem `json:"list"`  // 列表
	Stats map[string]int                 `json:"stats"` // 搜索统计
	Page  int                            `json:"page"`  // 分页码
	Size  int                            `json:"size"`  // 分页数量
	Total int                            `json:"total"` // 数据总数
}

type GoodsOptionsGetListOutputItem struct {
	entity.GoodsOptionsInfo
	//Id       uint  `json:"id"`
	//UserId   uint  `json:"user_id"`
	//CouponId uint  `json:"coupon_id"`
	//Status   uint8 `json:"status"`
	//
	//CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	//UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间

}

type GoodsOptionsSearchOutputItem struct {
	GoodsOptionsGetListOutputItem
}

// GoodsOptionsListItem 主要用于列表展示
//type GoodsOptionsListItem struct {
//	Id        uint        `json:"id"` // 自增ID
//	PicUrl    string      `json:"pic_url"`
//	Link      string      `json:"link"`
//	Sort      uint        `json:"sort"`       // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
//	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
//	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
//}
