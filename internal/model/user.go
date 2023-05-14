package model

import "github.com/gogf/gf/v2/frame/g"

type RegisterInput struct {
	Name         string
	Avatar       string
	Password     string
	UserSalt     string
	Sex          int
	Status       int
	Sign         string
	SecretAnswer string
}

type RegisterOutput struct {
	UserId uint
}
type LoginInput struct {
	Name     string
	Password string
}
type UpdatePasswordInput struct {
	Password     string
	UserSalt     string
	SecretAnswer string
}

type UpdatePasswordOutput struct {
	UserId uint
}

type UserInfoBase struct {
	g.Meta `orm:"table:user_info""`
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Sex    uint8  `json:"sex"`
	Sign   string `json:"sign"`
	Status uint8  `json:"status"`
}
