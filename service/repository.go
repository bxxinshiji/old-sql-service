package service

import (
	"log"

	gpd "gitee.com/xsjcs/old-sql-service/proto/goods"
	"github.com/go-xorm/xorm"
)

//Repository 商品仓库接口
type Repository interface {
	Get(good *gpd.Good) (*gpd.Good, error)
}

// GoodsRepository 用户仓库
type GoodsRepository struct {
	Engine *xorm.Engine
}

// Get 获取用户信息
func (repo *GoodsRepository) Get(good *gpd.Good) (*gpd.Good, error) {
	log.Println(repo)
	return good, nil
}
