// 商品
// 注意这里的包已经不是v1的包了，应该换成backend
package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// 这里定义的是对外的接口,和一些请求的参数，返回的结构体

// Get列表接口
type GoodsGetListCommonReq struct {
	g.Meta `path:"/goods/list" method:"get" tags:"商品" summary:"商品列表接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}
type GoodsGetListCommonRes struct {
	//todo
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`

	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

// 返回数据库中的所有内容的Req
type GoodsGetAllListCommonReq struct {
	g.Meta `path:"/goods/allList" method:"get" tags:"商品" summary:"商品全部列表"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	//CommonPaginationReq
}
type GoodsGetAllListCommonRes struct {
	//todo
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`

	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

// 商品详情
type GoodsDetailReq struct {
	g.Meta `path:"/goods/detail" method:"post" tags:"前台商品" summary:"商品详情"`
	Id     uint `json:"id"`
}
type GoodsDetailRes struct {
	GoodsInfoBase
	Options  []GoodsOptionsBase `json:"options"`
	Comments []CommentBase      `json:"Comments"`
	//IsCollect bool               `json:"IsCollect"`
}

type GoodsInfoBase struct {
	Id               int         `json:"id"               description:""`
	PicUrl           string      `json:"picUrl"           description:"图片"`
	Name             string      `json:"name"             description:"商品名称"`
	Price            int         `json:"price"            description:"价格 单位分"`
	Level1CategoryId int         `json:"level1CategoryId" description:"1级分类id"`
	Level2CategoryId int         `json:"level2CategoryId" description:"2级分类id"`
	Level3CategoryId int         `json:"level3CategoryId" description:"3级分类id"`
	Brand            string      `json:"brand"            description:"品牌"`
	Stock            int         `json:"stock"            description:"库存"`
	Sale             int         `json:"sale"             description:"销量"`
	Tags             string      `json:"tags"             description:"标签"`
	DetailInfo       string      `json:"detailInfo"       description:"商品详情"`
	CreatedAt        *gtime.Time `json:"createdAt"        description:""`
	UpdatedAt        *gtime.Time `json:"updatedAt"        description:""`
	DeletedAt        *gtime.Time `json:"deletedAt"        description:""`
}

type GoodsOptionsBase struct {
	Id      int    `json:"id"        description:""`
	GoodsId int    `json:"goodsId"   description:"商品id"`
	PicUrl  string `json:"picUrl"    description:"图片"`
	Name    string `json:"name"      description:"商品名称"`
	Price   int    `json:"price"     description:"价格 单位分"`
	Stock   int    `json:"stock"     description:"库存"`
}
type BaseGoodsColumns struct {
	gmeta.Meta `orm:"table:goods_info"`
	Id         string `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Brand      string `json:"brand"`
	Tags       string `json:"tags"`
	PicUrl     string `json:"pic_url"`
	DetailInfo string `json:"detail_info"`
}
