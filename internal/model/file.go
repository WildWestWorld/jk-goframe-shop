package model

import "github.com/gogf/gf/v2/net/ghttp"

type FileUploadInput struct {
	File       *ghttp.UploadFile //上传文件对象
	Name       string
	RandomName bool //是否是随机命名
}
type FileUploadOutput struct {
	Id   uint
	Name string
	Src  string
	Url  string
}
