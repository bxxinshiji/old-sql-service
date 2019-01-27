package service

import (
	"github.com/bxxinshiji/old-sql-service/models"
	"github.com/go-xorm/xorm"
)

// GoodsInfo 商品信息
func GoodsInfo(engine *xorm.Engine, goods *models.Goods) (err error) {
	res, err := engine.Table("tBmPlu").Get(goods)
	if err != nil {
		return err
	}
	if !res {
		// 多条形码时获取指定商品ID
		barCode := &models.BarCode{
			BarCode: goods.BarCode,
		}
		res, err := engine.Table("tbmMulBar").Get(barCode)
		if err != nil {
			return err
		}
		// 重新获取商品信息
		if res {
			// 缓存原始条形码
			bar := goods.BarCode
			goods.BarCode = ""
			goods.Id = barCode.GoodsId
			_, err := engine.Table("tBmPlu").Get(goods)
			if err != nil {
				return err
			}
			// 重新赋予原始条形码
			goods.BarCode = bar
		}
	}
	// 处理梳理
	goods.Hander()
	return err
}
