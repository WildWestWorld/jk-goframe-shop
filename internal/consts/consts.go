// 定义一些常量的地方
package consts

const (
	Version                  = "v0.2.0"             // 当前服务版本(用于模板展示)
	CaptchaDefaultName       = "CaptchaDefaultName" // 验证码默认存储空间名称
	ContextKey               = "ContextKey"         // 上下文变量存储键名，前后端系统共享
	FileMaxUploadCountMinute = 10                   // 同一用户1分钟之内最大上传数量

	GTokenAdminPrefix = "Admin:" //gtoken管理后台前缀区分，用于区分是后台用户还是前台用户
	GTokenFrontPrefix = "User:"  //gtoken管理后台前缀区分，用于区分是前台用户还是前台用户

	//for admin
	CtxAdminId      = "CtxAdminId"
	CtxAdminName    = "CtxAdminName"
	CtxAdminIsAdmin = "CtxAdminIsAdmin"
	CtxAdminRoleIds = "CtxAdminRoleIds"

	//for User
	CtxUserId     = "CtxUserId"
	CtxUserName   = "CtxUserName"
	CtxUserAvatar = "CtxUserAvatar"
	CtxUserSign   = "CtxUserSign"
	CtxUserStatus = "CtxUserStatus"
	CtxUserSex    = "CtxUserSex"

	//collection
	CollectionTypeGoods   = 1
	CollectionTypeArticle = 2

	//praise
	PraiseTypeGoods   = 1
	PraiseTypeArticle = 2
	//comment
	CommentTypeGoods   = 1
	CommentTypeArticle = 2

	CashModeRedis     = 2
	BackendServerName = "shop"
	MultiLogin        = true
	FrontMultiLogin   = false
	GTokenExpireIn    = 10 * 24 * 60 * 60
	//	统一管理错误提示
	CodeMissingParameterMsg = "请检查是否缺少参数"
)
