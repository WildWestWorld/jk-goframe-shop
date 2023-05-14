package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

//这里定义的是对外的接口,和一些请求的参数，返回的结构体

type Req struct {
	g.Meta `path:"/hello" tags:"Hello" method:"get" summary:"You first hello api"`
}
type Res struct {
	g.Meta `mime:"text/html" example:"string"`
}
