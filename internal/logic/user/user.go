package user

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"jk-goframe-shop/internal/consts"
	"jk-goframe-shop/internal/dao"
	"jk-goframe-shop/internal/model"
	"jk-goframe-shop/internal/model/do"
	"jk-goframe-shop/internal/service"
	"jk-goframe-shop/utility"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

// Register 创建内容
func (s *sUser) Register(ctx context.Context, in model.RegisterInput) (out model.RegisterOutput, err error) {

	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	//加盐
	//利用grand.S 方法随机生成一个10位数的字符串
	UserSalt := grand.S(10)
	//利用 EncryptPassword 方法对用户的密码进行加密
	//in.Password = utility.EncryptPassword(in.Password, UserSalt)
	//利用 框架给我们的 gmd5 方法对用户的密码进行加密
	//md5加密(md5加密密码+md5加密随机数)
	in.Password = gmd5.MustEncryptString(gmd5.MustEncryptString(in.Password) + gmd5.MustEncryptString(UserSalt))
	in.UserSalt = UserSalt

	lastInsertID, err := dao.UserInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RegisterOutput{UserId: uint(lastInsertID)}, err
}

// 修改密码
func (s *sUser) UpdatePassword(ctx context.Context, in model.UpdatePasswordInput) (out model.UpdatePasswordOutput, err error) {
	//	验证密保问题
	userInfo := do.UserInfo{}
	userId := consts.CtxAdminId
	userIdFormat := gconv.Uint(ctx.Value(userId))
	err = dao.UserInfo.Ctx(ctx).WherePri(ctx.Value(userId)).Scan(&userInfo)
	if err != nil {
		return model.UpdatePasswordOutput{}, err
	}
	if gconv.String(userInfo.SecretAnswer) != in.SecretAnswer {
		return out, errors.New("密保问题答案不一致")
	}

	userSalt := grand.S(10)
	in.UserSalt = userSalt
	in.Password = utility.EncryptPassword(in.Password, userSalt)

	_, err = dao.UserInfo.Ctx(ctx).WherePri(userId).Update(in)
	if err != nil {
		return model.UpdatePasswordOutput{}, err
	}
	return model.UpdatePasswordOutput{UserId: userIdFormat}, err
}
