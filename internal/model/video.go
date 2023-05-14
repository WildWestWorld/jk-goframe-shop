package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// VideoCreateUpdateBase 创建/修改内容基类
type VideoCreateUpdateBase struct {
	Type   int
	Title  string
	Url    string
	Cover  string
	Sticky string

	View     int
	Duration int
	Favorite int
}

// VideoCreateInput 创建内容
type VideoCreateInput struct {
	VideoCreateUpdateBase
}

// VideoCreateOutput 创建内容返回结果
type VideoCreateOutput struct {
	VideoId int `json:"video_id"`
}

// VideoUpdateInput 修改内容
type VideoUpdateInput struct {
	VideoCreateUpdateBase
	Id uint
}

// VideoGetListInput 获取内容列表
type VideoGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Type int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// VideoGetListOutput 查询列表结果
type VideoGetListOutput struct {
	List  []VideoGetListOutputItem `json:"list" description:"列表"`
	Page  int                      `json:"page" description:"分页码"`
	Size  int                      `json:"size" description:"分页数量"`
	Total int                      `json:"total" description:"数据总数"`
}

// VideoSearchInput 搜索列表
type VideoSearchInput struct {
	Key        string // 关键字
	Type       string // 内容模型
	CategoryId uint   // 栏目ID
	Page       int    // 分页号码
	Size       int    // 分页数量，最大50
	Sort       int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// VideoSearchOutput 搜索列表结果
type VideoSearchOutput struct {
	List  []VideoSearchOutputItem `json:"list"`  // 列表
	Stats map[string]int          `json:"stats"` // 搜索统计
	Page  int                     `json:"page"`  // 分页码
	Size  int                     `json:"size"`  // 分页数量
	Total int                     `json:"total"` // 数据总数
}

type VideoGetListOutputItem struct {
	Id       uint   `json:"id"` // 自增ID
	Type     int    `json:"type"`
	Title    string `json:"title"`
	Url      string `json:"url"`
	Cover    string `json:"cover"`
	Sticky   string `json:"sticky"`
	View     int    `json:"view"`
	Duration int    `json:"duration"`
	Favorite int    `json:"favorite"`

	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}

type VideoSearchOutputItem struct {
	VideoGetListOutputItem
}

// VideoListItem 主要用于列表展示
//type VideoListItem struct {
//	Id        uint        `json:"id"` // 自增ID
//	PicUrl    string      `json:"pic_url"`
//	Link      string      `json:"link"`
//	Sort      uint        `json:"sort"`       // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
//	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
//	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
//}
