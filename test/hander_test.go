package test

import (
	"context"
	"testing"

	"github.com/bxxinshiji/old-sql-service/hander"
	gpd "github.com/bxxinshiji/old-sql-service/proto/goods"
	ipd "github.com/bxxinshiji/old-sql-service/proto/inventory"
	db "github.com/bxxinshiji/old-sql-service/providers/database"
	"github.com/bxxinshiji/old-sql-service/service"
)

func TestGoodsGet(t *testing.T) {
	// 商品仓库 db 接口实现
	repo := &service.GoodsRepository{db.Engine}
	h := hander.Goods{repo}
	req := &gpd.Good{
		BarCode: "3693189830097",
	}
	res := &gpd.Response{}
	err := h.Get(context.TODO(), req, res)
	if err != nil {
		t.Errorf("Query goods failed, %v!", err)
	}
}

func TestInventoryReport(t *testing.T) {
	// 商品仓库 db 接口实现
	repo := &service.InventoryRepository{db.Engine}
	h := hander.Inventory{repo}

	req := &ipd.Request{}
	var goods []*ipd.Good
	goods = append(
		goods,
		&ipd.Good{
			BarCode: "2000019810223",
			Number:  "2",
		},
		&ipd.Good{
			BarCode: "3693189830097",
			Number:  "3",
		},
	)
	req.Goods = goods
	res := &ipd.Response{}
	err := h.Report(context.TODO(), req, res)
	if err != nil {
		t.Errorf("Inventory count data failed, %v!", err)
	}
}
