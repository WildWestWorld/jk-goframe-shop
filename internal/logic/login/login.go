package login

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"jk-goframe-shop/internal/dao"
	"jk-goframe-shop/internal/model"
	"jk-goframe-shop/internal/model/entity"
	"jk-goframe-shop/internal/service"
	"jk-goframe-shop/utility"
)

type sLogin struct{}

func init() {
	service.RegisterLogin(New())
}

func New() *sLogin {
	return &sLogin{}
}

// 执行登录
func (s *sLogin) Login(ctx context.Context, in model.UserLoginInput) error {
	//userEntity, err := s.GetLoginByPassportAndPassword(
	//	ctx,
	//	in.Passport,
	//	s.EncryptPassword(in.Passport, in.Password),
	//)
	//if err != nil {
	//	return err
	//}
	//if userEntity == nil {
	//	return gerror.New(`账号或密码错误`)
	//}

	//定义一个用户的实体类
	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where("name", in.Name).Scan(&adminInfo)
	if err != nil {
		return err
	}
	//gutil.Dump("加密后密码：", utility.EncryptPassword(in.Name, adminInfo.UserSalt))
	//如果 gmd5( 当前输入的密码 + 数据库中的加密盐 ) 不等于 数据库中的 加密后的密码   那么就是输入了错误的密码
	//因为我们 生成加密密码的方式 是   gmd5 (加密盐 + 密码) ，如果输入的密码是和数据的密码不一致 生成的加密后的密码就自然是不一致的
	if utility.EncryptPassword(in.Password, adminInfo.UserSalt) != adminInfo.Password {
		return gerror.New("账号或者密码不正确")
	}

	if err := service.Session().SetUser(ctx, &adminInfo); err != nil {
		return err
	}
	// 自动更新上线
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:      uint(adminInfo.Id),
		Name:    adminInfo.Name,
		IsAdmin: uint8(adminInfo.IsAdmin),
	})
	return nil
}
