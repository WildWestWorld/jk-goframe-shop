// 视频
// 注意这里的包已经不是v1的包了，应该换成backend
package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 这里定义的是对外的接口,和一些请求的参数，返回的结构体

// Req 发请求的payload
type VideoReq struct {
	g.Meta `path:"/backend/video/add" tags:"Video" method:"post" summary:"视频 添加接口"`
	Type   int    `json:"type"      v:"required#视频所属的类型不能为空"   dc:"视频所属的类型"`
	Title  string `json:"title"      v:"required#视频标题不能为空"   dc:"视频标题"`
	Url    string `json:"url"      v:"required#视频链接不能为空"   dc:"视频链接"`
	Cover  string `json:"cover"      v:"required#封面图片链接不能为空"   dc:"封面图片链接"`
	Sticky string `json:"sticky"       dc:"视频属于哪个topBar"`

	View     int `json:"view"       dc:"视频播放数量"`
	Favorite int `json:"favorite"       dc:"喜欢次量"`
	Duration int `json:"duration"       dc:"视频播放时长"`
}

// Req 请求完毕后的结果/响应
type VideoRes struct {
	//g.Meta `mime:"text/html" example:"string"`
	VideoId uint `json:"video_id"`
}

type VideoDeleteReq struct {
	g.Meta `path:"/backend/video/delete" method:"delete" tags:"内容" summary:"删除内容接口"`
	Id     uint `v:"min:1#请选择需要删除的内容" dc:"内容id"`
}

type VideoDeleteRes struct{}

type VideoUpdateReq struct {
	g.Meta `path:"/backend/video/update/{Id}" method:"post" tags:"视频" summary:"修改视频接口"`
	Id     uint `json:"id"      v:"min:1#请选择需要修改的视频Id" dc:"视频Id"`

	Type     int    `json:"type"         dc:"视频所属的类型"`
	Title    string `json:"title"       dc:"视频标题"`
	Url      string `json:"url"      dc:"视频链接"`
	Cover    string `json:"cover"       dc:"封面图片链接"`
	Sticky   string `json:"sticky"       dc:"视频属于哪个topBar"`
	View     int    `json:"view"       dc:"视频播放数量"`
	Favorite int    `json:"favorite"       dc:"喜欢次量"`
	Duration int    `json:"duration"       dc:"视频播放时长"`
}
type VideoUpdateRes struct {
	VideoId uint `json:"video_id"`
}

// Get列表接口
type VideoGetListCommonReq struct {
	g.Meta `path:"/backend/video/list" method:"get" tags:"视频" summary:"视频列表接口"`
	Type   int `json:"type"   in:"query" dc:"排序类型" d:"0"`
	CommonPaginationReq
}
type VideoGetListCommonRes struct {
	//todo
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`

	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
