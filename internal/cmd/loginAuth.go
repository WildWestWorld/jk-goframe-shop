package cmd

import (
	"context"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"jk-goframe-shop/api/backend"
	"jk-goframe-shop/api/frontend"
	"jk-goframe-shop/internal/consts"
	"jk-goframe-shop/internal/dao"
	"jk-goframe-shop/internal/model/entity"
	"jk-goframe-shop/utility"
	"jk-goframe-shop/utility/response"
	"strconv"
)

// 管理后台
func StartBackendGToken() (gfToken *gtoken.GfToken, err error) {
	gfToken = &gtoken.GfToken{
		ServerName:       consts.BackendServerName, //自定义的项目名称
		CacheMode:        consts.CashModeRedis,     //设置缓存的模式,1.gCash2.gRedis   gCash 每次重启服务都会导致token 失效
		LoginPath:        "/backend/login",         //登录的api路径  这里添加了后Controller层就不用写了
		LoginBeforeFunc:  loginFunc,                //登录鉴权的函数 也就是他替代了我们Controller
		LoginAfterFunc:   loginAfterFunc,           //登录后的执行的函数
		LogoutPath:       "/backend/user/logout",
		AuthAfterFunc:    authAfterFunc,
		AuthPaths:        g.SliceStr{"/backend/user/info"},              //需要鉴权的路由
		AuthExcludePaths: g.SliceStr{"/user/info", "/system/user/info"}, // 不拦截路径 /user/info,/system/user/info,/system/user,
		MultiLogin:       consts.MultiLogin,                             //多点登录
	}

	err = gfToken.Start()

	return

}

// 前台
func StartFrontGToken() (gfFrontToken *gtoken.GfToken, err error) {
	gfFrontToken = &gtoken.GfToken{
		ServerName:      consts.BackendServerName, //自定义的项目名称
		CacheMode:       consts.CashModeRedis,     //设置缓存的模式,1.gCash2.gRedis   gCash 每次重启服务都会导致token 失效
		LoginPath:       "/frontend/login",        //登录的api路径  这里添加了后Controller层就不用写了
		LoginBeforeFunc: loginFuncFrontend,        //登录鉴权的函数 也就是他替代了我们Controller
		LoginAfterFunc:  loginAfterFrontendFunc,   //登录后的执行的函数
		LogoutPath:      "/frontend/user/logout",
		AuthAfterFunc:   authAfterFrontendFunc,
		//AuthPaths:        g.SliceStr{"/backend/user/info"},              //需要鉴权的路由
		//AuthExcludePaths: g.SliceStr{"/user/info", "/system/user/info"}, // 不拦截路径 /user/info,/system/user/info,/system/user,
		MultiLogin: consts.FrontMultiLogin, //多点登录
	}

	//err = gfFrontToken.Start()

	return

}

// 登录
func loginFunc(r *ghttp.Request) (string, interface{}) {
	name := r.Get("name").String()
	password := r.Get("password").String()
	ctx := context.TODO()

	if name == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误."))
		r.ExitAll()
	}

	//定义一个用户的实体类
	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where("name", name).Scan(&adminInfo)
	if err != nil {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误"))
		r.ExitAll()
	}
	//gutil.Dump("加密后密码：", utility.EncryptPassword(in.Name, adminInfo.UserSalt))
	//如果 gmd5( 当前输入的密码 + 数据库中的加密盐 ) 不等于 数据库中的 加密后的密码   那么就是输入了错误的密码
	//因为我们 生成加密密码的方式 是   gmd5 (加密盐 + 密码) ，如果输入的密码是和数据的密码不一致 生成的加密后的密码就自然是不一致的
	if utility.EncryptPassword(password, adminInfo.UserSalt) != adminInfo.Password {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误"))
		r.ExitAll()
	}

	//if err := service.Session().SetUser(ctx, &adminInfo); err != nil {
	//	return err
	//}
	//// 自动更新上线
	//service.BizCtx().SetUser(ctx, &model.ContextUser{
	//	Id:      uint(adminInfo.Id),
	//	Name:    adminInfo.Name,
	//	IsAdmin: uint8(adminInfo.IsAdmin),
	//})
	//return nil

	// 唯一标识，扩展参数user data

	//strconv.Itoa 数字转字符串
	//前面的id是唯一标识符 后面的adminInfo是额外参数
	return consts.GTokenAdminPrefix + strconv.Itoa(adminInfo.Id), adminInfo
}

// 前台登录
func loginFuncFrontend(r *ghttp.Request) (string, interface{}) {
	name := r.Get("name").String()
	password := r.Get("password").String()
	ctx := context.TODO()

	if name == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误."))
		r.ExitAll()
	}

	//定义一个用户的实体类
	userInfo := entity.UserInfo{}
	err := dao.UserInfo.Ctx(ctx).Where("name", name).Scan(&userInfo)
	if err != nil {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误"))
		r.ExitAll()
	}
	//gutil.Dump("加密后密码：", utility.EncryptPassword(in.Name, adminInfo.UserSalt))
	//如果 gmd5( 当前输入的密码 + 数据库中的加密盐 ) 不等于 数据库中的 加密后的密码   那么就是输入了错误的密码
	//因为我们 生成加密密码的方式 是   gmd5 (加密盐 + 密码) ，如果输入的密码是和数据的密码不一致 生成的加密后的密码就自然是不一致的
	if utility.EncryptPassword(password, userInfo.UserSalt) != userInfo.Password {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误"))
		r.ExitAll()
	}

	//if err := service.Session().SetUser(ctx, &adminInfo); err != nil {
	//	return err
	//}
	//// 自动更新上线
	//service.BizCtx().SetUser(ctx, &model.ContextUser{
	//	Id:      uint(adminInfo.Id),
	//	Name:    adminInfo.Name,
	//	IsAdmin: uint8(adminInfo.IsAdmin),
	//})
	//return nil

	// 唯一标识，扩展参数user data

	//strconv.Itoa 数字转字符串
	//前面的id是唯一标识符 后面的adminInfo是额外参数
	return consts.GTokenFrontPrefix + strconv.Itoa(userInfo.Id), userInfo
}

// 自定义的登录之后的函数
func loginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		respData.Code = 0
		r.Response.WriteJson(respData)
		return
	} else {
		respData.Code = 1
		//获得登录用户id
		userKey := respData.GetString("userKey")
		//gstr.StrEx(字符串，要被裁剪的字符串) :作用裁剪字符串
		//为什么什么要裁剪，因为我们的标识符 是带自定义的前缀的，现在我们要裁剪后用与裁剪

		adminId := gstr.StrEx(userKey, consts.GTokenAdminPrefix)
		//根据id获得登录用户其他信息
		adminInfo := entity.AdminInfo{}
		err := dao.AdminInfo.Ctx(context.TODO()).WherePri(adminId).Scan(&adminInfo)
		if err != nil {
			return
		}
		//通过角色查询权限
		//先通过角色查询权限id
		var rolePermissionInfos []entity.RolePermissionInfo
		err = dao.RolePermissionInfo.Ctx(context.TODO()).WhereIn(dao.RolePermissionInfo.Columns().RoleId, g.Slice{adminInfo.RoleIds}).Scan(&rolePermissionInfos)
		if err != nil {
			return
		}
		permissionIds := g.Slice{}
		for _, info := range rolePermissionInfos {
			permissionIds = append(permissionIds, info.PermissionId)
		}

		var permissions []entity.PermissionInfo
		err = dao.PermissionInfo.Ctx(context.TODO()).WhereIn(dao.PermissionInfo.Columns().Id, permissionIds).Scan(&permissions)
		if err != nil {
			return
		}
		data := &backend.LoginRes{
			Type:        "Bearer",
			Token:       respData.GetString("token"),
			ExpireIn:    consts.GTokenExpireIn, //单位秒,
			IsAdmin:     adminInfo.IsAdmin,
			RoleIds:     adminInfo.RoleIds,
			Permissions: permissions,
		}
		response.JsonExit(r, 0, "", data)
	}
	return
}

// 自定义的前台登录之后的函数
func loginAfterFrontendFunc(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		respData.Code = 0
		r.Response.WriteJson(respData)
		return
	} else {
		respData.Code = 1
		//获得登录用户id
		userKey := respData.GetString("userKey")
		//gstr.StrEx(字符串，要被裁剪的字符串) :作用裁剪字符串
		//为什么什么要裁剪，因为我们的标识符 是带自定义的前缀的，现在我们要裁剪后用与裁剪

		userId := gstr.StrEx(userKey, consts.GTokenFrontPrefix)
		//根据id获得登录用户其他信息
		userInfo := entity.UserInfo{}
		err := dao.UserInfo.Ctx(context.TODO()).WherePri(userId).Scan(&userInfo)
		if err != nil {
			return
		}
		////通过角色查询权限
		////先通过角色查询权限id
		//var rolePermissionInfos []entity.RolePermissionInfo
		//err = dao.RolePermissionInfo.Ctx(context.TODO()).WhereIn(dao.RolePermissionInfo.Columns().RoleId, g.Slice{userId.RoleIds}).Scan(&rolePermissionInfos)
		//if err != nil {
		//	return
		//}
		//permissionIds := g.Slice{}
		//for _, info := range rolePermissionInfos {
		//	permissionIds = append(permissionIds, info.PermissionId)
		//}
		//
		//var permissions []entity.PermissionInfo
		//err = dao.PermissionInfo.Ctx(context.TODO()).WhereIn(dao.PermissionInfo.Columns().Id, permissionIds).Scan(&permissions)
		//if err != nil {
		//	return
		//}
		data := &frontend.LoginRes{
			Type:     "Bearer",
			Token:    respData.GetString("token"),
			ExpireIn: consts.GTokenExpireIn, //单位秒,
			Name:     userInfo.Name,
			Avatar:   userInfo.Avatar,
			Sign:     userInfo.Sign,
			Status:   uint8(userInfo.Status),
		}
		response.JsonExit(r, 0, "", data)
	}
	return
}

func authAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	var adminInfo entity.AdminInfo
	err := gconv.Struct(respData.GetString("data"), &adminInfo)
	if err != nil {
		response.Auth(r)
		return
	}
	//todo 这里可以写账号前置校验、是否被拉黑、有无权限等逻辑
	//放到全局变量中
	r.SetCtxVar(consts.CtxAdminId, adminInfo.Id)
	r.SetCtxVar(consts.CtxAdminName, adminInfo.Name)
	r.SetCtxVar(consts.CtxAdminIsAdmin, adminInfo.IsAdmin)
	r.SetCtxVar(consts.CtxAdminRoleIds, adminInfo.RoleIds)
	r.Middleware.Next()
}

// 登录鉴权中间件for后台
func authAfterFrontendFunc(r *ghttp.Request, respData gtoken.Resp) {
	var userInfo entity.UserInfo
	err := gconv.Struct(respData.GetString("data"), &userInfo)
	if err != nil {
		response.Auth(r)
		return
	}
	//todo 这里可以写账号前置校验、是否被拉黑、有无权限等逻辑
	//放到全局变量中
	r.SetCtxVar(consts.CtxUserId, userInfo.Id)
	r.SetCtxVar(consts.CtxUserName, userInfo.Name)
	r.SetCtxVar(consts.CtxUserStatus, userInfo.Status)
	r.SetCtxVar(consts.CtxUserAvatar, userInfo.Avatar)
	r.SetCtxVar(consts.CtxUserSign, userInfo.Sign)
	r.SetCtxVar(consts.CtxUserSex, userInfo.Sex)
	r.Middleware.Next()
}
