package models

import (
	"time"
)

// Inventory 判断入库结构表
// 真他妈复杂
type Inventory struct {
	Date       time.Time `xorm:"comment('创建时间') 'XsDate'"`
	Time       string    `xorm:"comment('创建时间') 'XsTime'"`
	JzDate     time.Time `xorm:"comment('结账日期') 'JzDate'"`
	OrderNo    string    `xorm:"comment('订单编号') 'SaleItemNo'"`
	UserCode   string    `xorm:"comment('用户编号') 'UserCode'"`
	CurrDate   string    `xorm:"comment('日期') 'CurrDate'"`
	BcCode     string    `xorm:"comment('代码') 'BcCode'"`
	ClerkCode  string    `xorm:"comment('职员代码') 'ClerkCode'"`
	PageNo     string    `xorm:"comment('页面编号') 'PageNo'"`
	LnNo       string    `xorm:"comment('订单商品编号') 'LnNo'"`
	GoodsId    string    `xorm:"comment('商品ID') 'PluCode'"`
	BarCode    string    `xorm:"comment('订单商品编号') 'BarCode'"`
	GoodsName  string    `xorm:"comment('商品名称') 'PluName'"`
	PluAbbr    string    `xorm:"comment('商品别名') 'PluAbbr'"`
	Department string    `xorm:"comment('部门编码') 'DepCode'"`
	ClsCode    string    `xorm:"comment('部门编码') 'ClsCode'"`
	SupCode    string    `xorm:"comment('未知') 'SupCode'"`
	BrandCode  string    `xorm:"comment('品牌') 'BrandCode'"`
	Unit       string    `xorm:"comment('单位') 'Unit'"`
	Spec       string    `xorm:"comment('规格') 'Spec'"`
	XTaxRate   string    `xorm:"comment('销项税') 'TaxRate'"`
	SPrice     string    `xorm:"comment('单价') 'SPrice'"`
	Count      string    `xorm:"comment('盘点数量') 'XsCount'"`
	YsAmt      string    `xorm:"comment('库存总价') 'YsAmt'"`
	IsDecimal  string    `xorm:"comment('是否十进制') 'IsDecimal'"`
	Tag        string    `xorm:"comment('标签') 'Tag'"`
	IsGenDno   string    `xorm:"comment('未知') 'IsGenDno'"`
	BakData1   string    `xorm:"comment('未知') 'BakData1'"`
	BakData2   string    `xorm:"comment('未知') 'BakData2'"`
	BakData3   string    `xorm:"comment('未知') 'BakData3'"`
}
