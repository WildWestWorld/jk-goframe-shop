package model

import "github.com/gogf/gf/v2/os/gtime"

//Permission Model用于Service

// PermissionCreateUpdateBase 创建/修改内容基类
type PermissionCreateUpdateBase struct {
	Name string
	Path string
}

// PermissionCreateInput 创建内容
type PermissionCreateInput struct {
	PermissionCreateUpdateBase
}

// PermissionCreateOutput 创建内容返回结果
type PermissionCreateOutput struct {
	PermissionId int `json:"role_id"`
}

// PermissionUpdateInput 修改内容
type PermissionUpdateInput struct {
	PermissionCreateUpdateBase
	Id uint
}

// PermissionGetListInput 获取内容列表
type PermissionGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
}

// PermissionGetListOutput 查询列表结果
type PermissionGetListOutput struct {
	List  []PermissionGetListOutputItem `json:"list" description:"列表"`
	Page  int                           `json:"page" description:"分页码"`
	Size  int                           `json:"size" description:"分页数量"`
	Total int                           `json:"total" description:"数据总数"`
}

//// PermissionSearchInput 搜索列表
//type PermissionSearchInput struct {
//	Key        string // 关键字
//	Type       string // 内容模型
//	CategoryId uint   // 栏目ID
//	Page       int    // 分页号码
//	Size       int    // 分页数量，最大50
//	Sort       int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
//}
//
//// PermissionSearchOutput 搜索列表结果
//type PermissionSearchOutput struct {
//	List  []PermissionSearchOutputItem `json:"list"`  // 列表
//	Stats map[string]int          `json:"stats"` // 搜索统计
//	Page  int                     `json:"page"`  // 分页码
//	Size  int                     `json:"size"`  // 分页数量
//	Total int                     `json:"total"` // 数据总数
//}

type PermissionGetListOutputItem struct {
	Id   uint   `json:"id"` // 自增ID
	Name string `json:"name"`
	Path string `json:"desc"`

	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}

type PermissionSearchOutputItem struct {
	PermissionGetListOutputItem
}

// PermissionListItem 主要用于列表展示
//type PermissionListItem struct {
//	Id        uint        `json:"id"` // 自增ID
//	PicUrl    string      `json:"pic_url"`
//	Link      string      `json:"link"`
//	Sort      uint        `json:"sort"`       // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
//	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
//	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
//}
