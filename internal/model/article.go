package model

import "jk-goframe-shop/internal/model/entity"

// ArticleCreateUpdateBase 创建/修改内容基类
type ArticleCreateUpdateBase struct {
	UserId  int
	Title   string
	Desc    string
	PicUrl  string
	IsAdmin int
	Praise  int
	Detail  string
}

// ArticleCreateInput 创建内容
type ArticleCreateInput struct {
	ArticleCreateUpdateBase
}

// ArticleCreateOutput 创建内容返回结果
type ArticleCreateOutput struct {
	ArticleId int `json:"user_coupon_id"`
}

// ArticleUpdateInput 修改内容
type ArticleUpdateInput struct {
	ArticleCreateUpdateBase
	Id uint
}

// ArticleGetListInput 获取内容列表
type ArticleGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// ArticleGetListOutput 查询列表结果
type ArticleGetListOutput struct {
	List  []ArticleGetListOutputItem `json:"list" description:"列表"`
	Page  int                        `json:"page" description:"分页码"`
	Size  int                        `json:"size" description:"分页数量"`
	Total int                        `json:"total" description:"数据总数"`
}

// ArticleSearchInput 搜索列表
type ArticleSearchInput struct {
	Key       string // 关键字
	Type      string // 内容模型
	ArticleId uint   // 栏目ID
	Page      int    // 分页号码
	Size      int    // 分页数量，最大50
	Sort      int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// ArticleSearchOutput 搜索列表结果
type ArticleSearchOutput struct {
	List  []ArticleSearchOutputItem `json:"list"`  // 列表
	Stats map[string]int            `json:"stats"` // 搜索统计
	Page  int                       `json:"page"`  // 分页码
	Size  int                       `json:"size"`  // 分页数量
	Total int                       `json:"total"` // 数据总数
}

type ArticleGetListOutputItem struct {
	entity.ArticleInfo
	//Id       uint  `json:"id"`
	//UserId   uint  `json:"user_id"`
	//CouponId uint  `json:"coupon_id"`
	//Status   uint8 `json:"status"`
	//
	//CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	//UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间

}

type ArticleSearchOutputItem struct {
	ArticleGetListOutputItem
}

// ArticleListItem 主要用于列表展示
//type ArticleListItem struct {
//	Id        uint        `json:"id"` // 自增ID
//	PicUrl    string      `json:"pic_url"`
//	Link      string      `json:"link"`
//	Sort      uint        `json:"sort"`       // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
//	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
//	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
//}
