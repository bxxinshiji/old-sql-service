package service

import (
	"strconv"
	"time"

	"github.com/bxxinshiji/old-sql-service/helper"
	"github.com/bxxinshiji/old-sql-service/models"
	ipd "github.com/bxxinshiji/old-sql-service/proto/inventory"
	"github.com/go-xorm/xorm"
)

var (
	local = time.FixedZone("CST", 8*3600)
)

//Inventory 商品仓库接口
type Inventory interface {
	Report(req *ipd.Request) (*ipd.Response, error)
}

// InventoryRepository 用户仓库
type InventoryRepository struct {
	Engine *xorm.Engine
}

// Report 获取用户信息
func (repo *InventoryRepository) Report(req *ipd.Request) (res *ipd.Response, err error) {
	goods := req.Goods

	// 默认用户ID后期识别开发
	userID := "0000"
	// 当前时间
	currentTime := time.Now().In(local)
	OrderNo := userID + currentTime.Format("0601021504")
	// 查询订单存在数量
	total, err := repo.Engine.Table("tXsPosPD").Count(&models.Inventory{
		OrderNo: OrderNo,
	})
	if err != nil {
		return nil, err
	}

	for key, good := range goods {
		g := &models.Goods{
			BarCode: good.BarCode,
		}
		err := GoodsInfo(repo.Engine, g)
		if err != nil {
			return nil, err
		}

		// 计算订单下面的序号
		LnNo := strconv.FormatInt(int64(key)+1+total, 10)
		// 库存金额计算
		YsAmt := helper.StringToFloat(g.Price) * helper.StringToFloat(good.Number)

		i := &models.Inventory{
			Date:       currentTime,
			Time:       currentTime.Format("15:04:05"),
			JzDate:     currentTime,
			OrderNo:    OrderNo,
			UserCode:   userID,
			CurrDate:   currentTime.Format("20060102"),
			BcCode:     "1",
			ClerkCode:  "",
			PageNo:     "1",
			LnNo:       LnNo,
			GoodsId:    g.Id,
			BarCode:    g.BarCode,
			GoodsName:  g.Name,
			PluAbbr:    g.PluAbbr,
			Department: g.Department,
			ClsCode:    g.ClsCode,
			SupCode:    g.SupCode,
			BrandCode:  g.BrandCode,
			Unit:       g.Unit,
			Spec:       g.Spec,
			XTaxRate:   g.TaxRate,
			SPrice:     g.Price,
			Count:      good.Number,
			YsAmt:      helper.FloatToString(YsAmt),
			IsDecimal:  "0",
			Tag:        "1",
			IsGenDno:   "1",
			BakData3:   "0",
		}
		_, err = repo.Engine.Table("tXsPosPD").Insert(i)
		if err != nil {
			return nil, err
		}
	}
	return res, err
}
