// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"jk-goframe-shop/internal/dao/internal"
)

// internalGoodsInfoDao is internal type for wrapping internal DAO implements.
type internalGoodsInfoDao = *internal.GoodsInfoDao

// goodsInfoDao is the data access object for table goods_info.
// You can define custom methods on it to extend its functionality as you wish.
type goodsInfoDao struct {
	internalGoodsInfoDao
}

var (
	// GoodsInfo is globally public accessible object for table goods_info operations.
	GoodsInfo = goodsInfoDao{
		internal.NewGoodsInfoDao(),
	}
)

// Fill with you ideas below.
