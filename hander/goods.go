package hander

import (
	"context"

	gpd "github.com/bxxinshiji/old-sql-service/proto/goods"
	"github.com/bxxinshiji/old-sql-service/service"
)

// Goods 商品服务处理
type Goods struct {
	RepoGoods service.Goods
}

// Get 获取商品信息
func (srv *Goods) Get(ctx context.Context, req *gpd.Good, res *gpd.Response) (err error) {
	good, err := srv.RepoGoods.Get(req)
	if err != nil {
		return err
	}
	res.Good = good
	return err
}
