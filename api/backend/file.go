package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 请求
type FileUploadReq struct {
	//	所属路径
	//mine设置返回的数据类型 是json还是form-data
	// *ghttp.UploadFile 是框架上传文件默认的格式手机固定的
	g.Meta `path:"/backend/file" method:"post"  mine:"multipart/form-data" tags:"工具" dc:"上传文件"`
	File   *ghttp.UploadFile `json:"file" type:"file" dc:"选择上传文件"`
}

type FileUploadRes struct {
	Name string `json:"name" dc:"文件名称"`
	Url  string `json:"url" dc:"图片地址"`
}
