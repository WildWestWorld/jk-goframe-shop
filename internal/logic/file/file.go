package file

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"jk-goframe-shop/internal/consts"
	"jk-goframe-shop/internal/dao"
	"jk-goframe-shop/internal/model"
	"jk-goframe-shop/internal/model/entity"
	"jk-goframe-shop/internal/service"
	"time"
)

type sFile struct {
}

func init() {
	service.RegisterFile(New())
}

func New() *sFile {
	return &sFile{}
}

// 上传
// 1.定义图片上传位置
// 2.校验上传位置是否正确
// 3.定义年月日的目录
// 4.入库
// 5.返回数据
func (s *sFile) Upload(ctx context.Context, in model.FileUploadInput) (out *model.FileUploadOutput, err error) {
	//利用g.Cfg().MustGet(ctx,"在config文件中的路径")获取到我们存储的数据
	uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
	//如果上传路径格式空的,就报错
	if uploadPath == "" {
		return nil, gerror.New("读取配置文件失败,上传路径不存在")
	}
	//如果Controller层传入的路径不是空的就
	if in.Name != "" {
		//就把文件的名字改为我们Controller传入的名字
		in.File.Filename = in.Name
	}
	//安全性校验:1分钟只能上传10次
	count, err := dao.FileInfo.Ctx(ctx).
		Where(dao.FileInfo.Columns().UserId, gconv.Int(ctx.Value(consts.CtxAdminId))).
		WhereGTE(dao.FileInfo.Columns().CreatedAt, gtime.Now().Add(time.Minute)).Count()
	if err != nil {
		return nil, err
	}
	if count >= consts.FileMaxUploadCountMinute {
		return nil, gerror.New("上传频繁,一分钟只能上传10次")
	}

	//定义年月日YMD
	dateDirName := gtime.Now().Format("Ymd")
	//根据RandomName控制是否是随机名称
	//gfile.Join 用"/"拼接
	fileName, err := in.File.Save(gfile.Join(uploadPath, dateDirName), in.RandomName)
	if err != nil {
		return nil, err
	}
	//入库
	data := entity.FileInfo{
		Name:   fileName,
		Src:    gfile.Join(uploadPath, dateDirName, fileName),
		Url:    "/upload/" + dateDirName + "/" + fileName,
		UserId: gconv.Int(ctx.Value(consts.CtxAdminId)),
	}
	id, err := dao.FileInfo.Ctx(ctx).Data(data).OmitEmpty().InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &model.FileUploadOutput{Id: uint(id), Name: data.Name, Src: data.Src, Url: data.Url}, nil

}
