package backend

import "github.com/gogf/gf/v2/frame/g"

type PermissionCommonReq struct {
	Name string `json:"name" v:"required#权限名称必填" dc:"权限名称"`
	Path string `json:"path"  dc:"权限描述"`
}

// 请求
type PermissionReq struct {
	//	所属路径
	g.Meta `path:"/backend/permission/add" method:"post"  tag:"权限"`
	PermissionCommonReq
}

type PermissionRes struct {
	PermissionId int `json:"permission_id"`
	PermissionCommonReq
}

// 更新
type PermissionUpdateReq struct {
	g.Meta `path:"/backend/permission/update" method:"post" dc:"修改权限" tag:"permission"`
	Id     uint `json:"id" v:"required#更新的权限id必填" dc:"权限id"`
	PermissionCommonReq
}

// 更新的返回值
type PermissionUpdateRes struct {
	Id uint `json:"id"`
}

// 删除
type PermissionDeleteReq struct {
	g.Meta `path:"/backend/permission/delete" method:"delete" dc:"删除权限" tag:"permission"`
	Id     uint `json:"id" v:"required#删除的权限id必填" dc:"权限id"`
}

// 删除的返回值
type PermissionDeleteRes struct {
	Id uint `json:"id"`
}

// Get列表接口
type PermissionGetListCommonReq struct {
	g.Meta `path:"/backend/permission/list" method:"get" tags:"权限列表" summary:"权限列表"`
	CommonPaginationReq
}
type PermissionGetListCommonRes struct {
	//todo
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`

	List  interface{} `json:"list" dc:"列表"`
	Page  int         `json:"page" dc:"分页码"`
	Size  int         `json:"size" dc:"分页数量"`
	Total int         `json:"total" dc:"数据总数"`
}
