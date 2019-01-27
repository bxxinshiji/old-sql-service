package main

import (
	"fmt"
	"log"
	"os"

	db "github.com/bxxinshiji/old-sql-service/providers/database"

	"github.com/bxxinshiji/old-sql-service/hander"
	gpb "github.com/bxxinshiji/old-sql-service/proto/goods"
	ipb "github.com/bxxinshiji/old-sql-service/proto/inventory"
	"github.com/bxxinshiji/old-sql-service/service"
	micro "github.com/micro/go-micro"
)

var (
	// serviceName 服务名称
	serviceName = os.Getenv("SERVICE_NAME")
)

func main() {

	srv := micro.NewService(
		micro.Name(serviceName),
		micro.Version("latest"),
	)
	srv.Init()

	// 商品仓库 db 接口实现
	repoGoods := &service.GoodsRepository{db.Engine}
	// Register handler
	gpb.RegisterGoodsHandler(srv.Server(), &hander.Goods{repoGoods})

	// 盘点仓库 db 接口实现
	repoInventory := &service.InventoryRepository{db.Engine}
	// Register handler
	ipb.RegisterInventoryHandler(srv.Server(), &hander.Inventory{repoInventory})
	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
	log.Println("serviser run ...")
}
