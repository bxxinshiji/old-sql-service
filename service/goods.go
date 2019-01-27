package service

import (
	"github.com/bxxinshiji/old-sql-service/models"
	gpd "github.com/bxxinshiji/old-sql-service/proto/goods"
	"github.com/go-xorm/xorm"
)

//Goods 商品仓库接口
type Goods interface {
	Get(good *gpd.Good) (*gpd.Good, error)
}

// GoodsRepository 用户仓库
type GoodsRepository struct {
	Engine *xorm.Engine
}

// Get 获取用户信息
func (repo *GoodsRepository) Get(good *gpd.Good) (*gpd.Good, error) {
	goods := &models.Goods{
		BarCode: good.BarCode,
	}
	err := GoodsInfo(repo.Engine, goods)
	if err != nil {
		return nil, err
	}
	g := &gpd.Good{
		Id:       goods.Id,
		Name:     goods.Name,
		BarCode:  goods.BarCode,
		Price:    goods.Price,
		Standard: goods.Standard,
	}
	return g, err
}
