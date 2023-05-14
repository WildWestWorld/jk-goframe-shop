package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"jk-goframe-shop/internal/model/entity"
	"time"
)

type LoginIndexReq struct {
	g.Meta `path:"/login" method:"get" summary:"展示登录页面" tags:"登录"`
}
type LoginIndexRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type LoginDoReq struct {
	g.Meta   `path:"/backend/login" method:"post" summary:"执行登录请求" tags:"登录"`
	Name     string `json:"name" v:"required#请输入账号"   dc:"账号"`
	Password string `json:"password" v:"required#请输入密码"   dc:"密码(明文)"`
}

// ForJWT
type LoginDoRes struct {
	//Info interface{} `json:"info" dc:"用户信息"`

	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

// Token
type LoginRes struct {
	Type        string                  `json:"type"`
	Token       string                  `json:"token"`
	ExpireIn    int                     `json:"expire_in"`
	IsAdmin     int                     `json:"is_admin"`    //是否超管
	RoleIds     string                  `json:"role_ids"`    //角色
	Permissions []entity.PermissionInfo `json:"permissions"` //权限列表
}

type RefreshTokenReq struct {
	g.Meta `path:"/backend/refresh_token" method:"post"`
}

type RefreshTokenRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type LogoutReq struct {
	g.Meta `path:"/backend/logout" method:"post"`
}

type LogoutRes struct {
}
