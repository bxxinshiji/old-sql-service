package models

import (
	"github.com/MXi4oyu/Utils/cnencoder/gbk"
	"github.com/bxxinshiji/old-sql-service/helper"
)

// Goods 商品
type Goods struct {
	Id         string `xorm:"comment('商品ID') 'PluCode'"`
	Name       string `xorm:"comment('商品名称') 'PluName'"`
	PluAbbr    string `xorm:"comment('商品名称') 'PluAbbrName'"`
	BarCode    string `xorm:"comment('条形码') 'BarCode'"`
	Standard   string `xorm:"comment('商品规格') 'Spec'"`
	Price      string `xorm:"comment('商品规格') 'SPrice'"`
	Unit       string `xorm:"comment('计量单位') 'Unit'"`
	Department string `xorm:"comment('部门编码') 'DepCode'"`
	ClsCode    string `xorm:"comment('部门编码') 'ClsCode'"`
	SupCode    string `xorm:"comment('未知') 'SupCode'"`
	BrandCode  string `xorm:"comment('品牌') 'BrandCode'"`
	Spec       string `xorm:"comment('规格') 'Spec'"`
	TaxRate    string `xorm:"comment('税率') 'XTaxRate'"`
	TaxNo      string `xorm:"comment('税收分类编码') 'Produce'"`
	TaxName    string `xorm:"comment('税收分类编码简称') 'Grade'"`
}

// Hander 数据预处理
func (g *Goods) Hander() {
	g.Id = helper.TrimSpace(g.Id)
	g.Name = helper.TrimSpace(gbk.Decode(g.Name))
	g.PluAbbr = helper.TrimSpace(gbk.Decode(g.PluAbbr))
	g.BarCode = helper.TrimSpace(g.BarCode)
	g.Standard = helper.TrimSpace(gbk.Decode(g.Standard))
	g.Price = helper.TrimSpace(g.Price)
	g.Unit = helper.TrimSpace(g.Unit)
	g.Department = helper.TrimSpace(g.Department)
	g.ClsCode = helper.TrimSpace(g.ClsCode)
	g.SupCode = helper.TrimSpace(g.SupCode)
	g.BrandCode = helper.TrimSpace(g.BrandCode)
	g.Spec = helper.TrimSpace(g.Spec)
	g.TaxRate = helper.TrimSpace(g.TaxRate)
	g.TaxNo = helper.TrimSpace(g.TaxNo)
	g.TaxName = helper.TrimSpace(gbk.Decode(g.TaxName))
}
