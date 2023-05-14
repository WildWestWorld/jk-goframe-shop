// 管理员
// 注意这里的包已经不是v1的包了，应该换成backend
package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 这里定义的是对外的接口,和一些请求的参数，返回的结构体

// Req 发请求的payload
type AdminReq struct {
	g.Meta   `path:"/backend/admin/add" tags:"Admin" method:"post" summary:"管理员 添加接口"`
	Name     string `json:"name"    v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password"    v:"required#密码不能为空"     dc:"密码"`
	RoleIds  string `json:"role_ids"        dc:"角色id"`
	IsAdmin  string `json:"is_admin"    d:"0"      dc:"是否为超级管理员"`
}

// Req 请求完毕后的结果/响应
type AdminRes struct {
	//g.Meta `mime:"text/html" example:"string"`
	AdminId uint `json:"admin_id"`
}

type AdminDeleteReq struct {
	g.Meta `path:"/backend/admin/delete" method:"delete" tags:"内容" summary:"删除内容接口"`
	Id     uint `v:"min:1#请选择需要删除的内容" dc:"内容id"`
}
type AdminDeleteRes struct{}

type AdminUpdateReq struct {
	g.Meta   `path:"/backend/admin/update/{Id}" method:"post" tags:"管理员" summary:"修改管理员接口"`
	Id       uint   `json:"id"      v:"min:1#请选择需要修改的管理员Id" dc:"管理员Id"`
	Name     string `json:"name"    v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password"    v:"required#密码不能为空"     dc:"密码"`
	RoleIds  string `json:"role_ids"        dc:"角色id"`
	IsAdmin  string `json:"is_admin"         dc:"是否为超级管理员"`
}
type AdminUpdateRes struct {
	AdminId uint `json:"admin_id"`
}

// Get列表接口
type AdminGetListCommonReq struct {
	g.Meta `path:"/backend/admin/list" method:"get" tags:"管理员" summary:"管理员列表接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}
type AdminGetListCommonRes struct {
	//todo
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`

	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type AdminGetInfoReq struct {
	g.Meta `path:"/backend/admin/info" method:"get"`
}

// For JWT
//type AdminGetInfoRes struct {
//	Id          int    `json:"id"`
//	IdentityKey string `json:"identity_key"`
//	Payload     string `json:"payload"`
//}

// For token
type AdminGetInfoRes struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	RoleIds string `json:"role_ids"`
	IsAdmin int    `json:"is_admin"`
}
