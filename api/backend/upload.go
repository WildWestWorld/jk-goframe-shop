package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadImgToCloudReq struct {
	g.Meta `path:"/backend/upload/cloud" method:"post" mine:"multipart/form-data" tags:"工具" dc:"上传文件到云"`
	File   *ghttp.UploadFile `json:"file" dc:"选择上传的文件"`
}
type UploadImgToCloudRes struct {
	Url string `json:"url" dc:"文件地址"`
}
