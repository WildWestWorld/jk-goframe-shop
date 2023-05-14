package Admin

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"jk-goframe-shop/internal/model/entity"
	"jk-goframe-shop/internal/service"
	"jk-goframe-shop/utility"

	"jk-goframe-shop/internal/dao"
	"jk-goframe-shop/internal/model"
)

type sAdmin struct{}

func init() {
	service.RegisterAdmin(New())
}

func New() *sAdmin {
	return &sAdmin{}
}

// Create 创建内容
func (s *sAdmin) Create(ctx context.Context, in model.AdminCreateInput) (out model.AdminCreateOutput, err error) {

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

	lastInsertID, err := dao.AdminInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.AdminCreateOutput{AdminId: int(lastInsertID)}, err
}

// Delete 删除
func (s *sAdmin) Delete(ctx context.Context, id uint) error {
	return dao.AdminInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		// 删除内容(软删除)
		//_, err := dao.AdminInfo.Ctx(ctx).Where(g.Map{
		//	dao.AdminInfo.Columns().Id: id,
		//}).Delete()
		// 删除内容(硬删除)Unscoped().Delete() 就会硬删除数据
		//如果没有delete_at 字段 框架默认就会硬删除数据，反之则软删除数据

		_, err := dao.AdminInfo.Ctx(ctx).Where(g.Map{
			dao.AdminInfo.Columns().Id: id,
		}).Delete()

		return err
	})
}

// Update 修改
func (s *sAdmin) Update(ctx context.Context, in model.AdminUpdateInput) error {
	return dao.AdminInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		//如果有用户修改了密码，那么我们的加密盐的数据也要跟着变
		//要是用户没有修改密码，我们加密盐的数据是不能变的
		//如何判定用户有没有修改了密码？
		//只要看传入进来的payload的password是否有值就能判别
		if in.Password != "" {
			//加盐
			//利用grand.S 方法随机生成一个10位数的字符串
			UserSalt := grand.S(10)

			//利用 框架给我们的 gmd5 方法对用户的密码进行加密
			//md5加密(md5加密密码+md5加密随机数)
			in.Password = gmd5.MustEncryptString(gmd5.MustEncryptString(in.Password) + gmd5.MustEncryptString(UserSalt))
			in.UserSalt = UserSalt
		}

		_, err := dao.AdminInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.AdminInfo.Columns().Id).
			Where(dao.AdminInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}

// GetList 查询内容列表
func (s *sAdmin) GetList(ctx context.Context, in model.AdminGetListInput) (out *model.AdminGetListOutput, err error) {
	var (
		m = dao.AdminInfo.Ctx(ctx)
	)
	out = &model.AdminGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)
	// 排序方式(dec 倒序)
	//listModel = listModel.OrderDesc(dao.AdminInfo.Columns().Id)

	// 执行查询
	var list []*entity.AdminInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	// Admin
	//if err := listModel.ScanList(&out.List, "Admin"); err != nil {
	//	return out, err
	//}
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}

	return
}

//登录

func (s *sAdmin) GetAdminByNamePassword(ctx context.Context, in model.UserLoginInput) map[string]interface{} {
	//if in.Name == "admin" && in.Password == "admin" {
	//	return g.Map{
	//		"id":       1,
	//		"username": "admin",
	//	}
	//}
	//return nil

	//定义一个用户的实体类
	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where("name", in.Name).Scan(&adminInfo)
	if err != nil {
		return nil
	}
	//gutil.Dump("加密后密码：", utility.EncryptPassword(in.Name, adminInfo.UserSalt))
	//如果 gmd5( 当前输入的密码 + 数据库中的加密盐 ) 不等于 数据库中的 加密后的密码   那么就是输入了错误的密码
	//因为我们 生成加密密码的方式 是   gmd5 (加密盐 + 密码) ，如果输入的密码是和数据的密码不一致 生成的加密后的密码就自然是不一致的
	if utility.EncryptPassword(in.Password, adminInfo.UserSalt) != adminInfo.Password {
		return nil
	} else {
		return g.Map{
			"id":       adminInfo.Id,
			"userName": adminInfo.Name,
		}
	}
	return nil
}
