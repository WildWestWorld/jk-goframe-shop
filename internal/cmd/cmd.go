package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"jk-goframe-shop/internal/controller"
	"jk-goframe-shop/internal/service"
	"net/http"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	"jk-goframe-shop/internal/controller/hello"
)

func MiddlewareCORS(r *ghttp.Request) {
	//corsOptions.AllowDomain = []string{"goframe.org", "baidu.com"}
	corsOptions := r.Response.DefaultCORSOptions()
	//corsOptions.AllowDomain = []string{"goframe.org"}
	//如果不接受跨域 接口的逻辑不会执行
	if !r.Response.CORSAllowedOrigin(corsOptions) {
		r.Response.WriteStatus(http.StatusForbidden)
		return
	}

	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}

// 会在这里做路由的管理
var (
	Main = gcmd.Command{
		Name:  "电商项目",
		Usage: "作者JK",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			// 认证接口
			//loginFunc := Login
			// 启动gtoken

			//启动管理后台gToken
			gfToken, err := StartBackendGToken()
			if err != nil {
				return err
			}

			//后台
			s.Group("/", func(group *ghttp.RouterGroup) {
				//一定要写在前面group 不然没法跨域
				group.Middleware(service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
					//service.Middleware().CORS,
					MiddlewareCORS)

				//GFtoken 中间件绑定
				err := gfToken.Middleware(ctx, group)
				if err != nil {
					panic(err)
				}

				//绑定中间件 定义响应的结构
				//官方用默认的中间件
				//group.Middleware(ghttp.MiddlewareHandlerResponse)
				//我们自定义的中间件

				group.Bind(
					hello.New(),
					controller.Rotation,     //轮播图
					controller.Position,     //手工位
					controller.Admin.Create, //管理员
					controller.Admin.Delete,
					controller.Admin.Update,
					controller.Admin.List,
					controller.Login,      //登录
					controller.Role,       //角色
					controller.Video,      //视频
					controller.Permission, //视频

				)
				//如果上面已经绑定了路由，再在下面绑定就会报错
				//所以我们会把路由拆分
				//下面的路由是需要鉴权的
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(
						//service.Middleware().Auth, 用于JWT
						MiddlewareCORS,
						//service.Middleware().CORS,
					)
					group.ALLMap(g.Map{
						"/backend/admin/info": controller.Admin.Info,
					})
					group.Bind(
						controller.File,
						controller.Upload,
						controller.Category,
						controller.Coupon,
						controller.UserCoupon,
						controller.Goods,
						controller.GoodsOptions,
						controller.Article,
					)
				})

			})

			//启动管理前台gToken
			gfFrontToken, err := StartFrontGToken()
			if err != nil {
				return err
			}

			//前台
			s.Group("/frontend", func(group *ghttp.RouterGroup) {
				//一定要写在前面group 不然没法跨域
				group.Middleware(service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
					//service.Middleware().CORS,
					MiddlewareCORS)

				//GFtoken 中间件绑定
				err := gfFrontToken.Middleware(ctx, group)
				if err != nil {
					panic(err)
				}

				//绑定中间件 定义响应的结构
				//官方用默认的中间件
				//group.Middleware(ghttp.MiddlewareHandlerResponse)
				//我们自定义的中间件

				group.Bind(
					controller.User.Register,
					controller.Goods,
				)
				//如果上面已经绑定了路由，再在下面绑定就会报错
				//所以我们会把路由拆分
				//下面的路由是需要鉴权的
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(
						//service.Middleware().Auth, 用于JWT
						MiddlewareCORS,
						//service.Middleware().CORS,
					)
					group.ALLMap(g.Map{})
					group.Bind(
						controller.User.Info,
						controller.User.UpdatePassword,
						controller.Collection,
						controller.Praise,
						controller.Comment,
					)
				})

			})

			s.Run()
			return nil
		},
	}
)
