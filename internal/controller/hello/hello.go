package hello

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "jk-goframe-shop/api/hello/v1"
)

//Controller 用于承上启下 接受参数，调用service方法

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) Hello(ctx context.Context, req *v1.Req) (res *v1.Res, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}
