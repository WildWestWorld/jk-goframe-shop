package backend

import "github.com/gogf/gf/v2/frame/g"

// 请求
type RoleReq struct {
	//	所属路径
	g.Meta `path:"/backend/role/add" method:"post"  tag:"角色"`
	Name   string `json:"name" v:"required#角色名称必填" dc:"角色名称"`
	Desc   string `json:"Desc"  dc:"角色描述"`
}

type RoleRes struct {
	RoleId int `json:"role_id"`
}

// 更新
type RoleUpdateReq struct {
	g.Meta `path:"/backend/role/update" method:"post" dc:"修改角色" tag:"role"`
	Id     uint   `json:"id" v:"required#更新的角色id必填" dc:"角色id"`
	Name   string `json:"name"  dc:"角色名称"`
	Desc   string `json:"desc"  dc:"角色描述"`
}

// 更新的返回值
type RoleUpdateRes struct {
	Id uint `json:"id"`
}

// 删除
type RoleDeleteReq struct {
	g.Meta `path:"/backend/role/delete" method:"delete" dc:"删除角色" tag:"role"`
	Id     uint `json:"id" v:"required#删除的角色id必填" dc:"角色id"`
}

// 删除的返回值
type RoleDeleteRes struct {
	Id uint `json:"id"`
}

// Get列表接口
type RoleGetListCommonReq struct {
	g.Meta `path:"/backend/role/list" method:"get" tags:"角色列表" summary:"角色列表"`
	CommonPaginationReq
}
type RoleGetListCommonRes struct {
	//todo
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`

	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
