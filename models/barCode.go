package models

import (
	"github.com/MXi4oyu/Utils/cnencoder/gbk"
	"github.com/bxxinshiji/old-sql-service/helper"
)

// BarCode 商品
type BarCode struct {
	BarCode string `xorm:"comment('条形码') 'BarCode'"`
	GoodsId string `xorm:"comment('商品编码') 'PluCode'"`
	DepCode string `xorm:"comment('部门') 'DepCode'"`
	PluName string `xorm:"comment('商品名称') 'PluName'"`
	Spec    string `xorm:"comment('规格') 'Spec'"`
}

// Hander 数据预处理
func (b *BarCode) Hander() {
	b.GoodsId = helper.TrimSpace(b.GoodsId)
	b.BarCode = helper.TrimSpace(b.BarCode)
	b.DepCode = helper.TrimSpace(b.DepCode)
	b.PluName = helper.TrimSpace(gbk.Decode(b.PluName))
	b.Spec = helper.TrimSpace(gbk.Decode(b.Spec))
}
