package hander

import (
	"context"

	ipd "github.com/bxxinshiji/old-sql-service/proto/inventory"
	"github.com/bxxinshiji/old-sql-service/service"
)

// Inventory 商品服务处理
type Inventory struct {
	RepoInventory service.Inventory
}

// Report 获取商品信息
func (srv *Inventory) Report(ctx context.Context, req *ipd.Request, res *ipd.Response) (err error) {
	_, err = srv.RepoInventory.Report(req)
	if err != nil {
		return err
	}
	res.Valid = true
	return err
}
