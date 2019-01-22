package main

import (
	"fmt"
	"log"
	"os"

	db "gitee.com/xsjcs/old-sql-service/providers/database"

	// 执行数据迁移
	pb "gitee.com/xsjcs/old-sql-service/proto/goods"
	"gitee.com/xsjcs/old-sql-service/service"
	micro "github.com/micro/go-micro"
)

var (
	// serviceName 服务名称
	serviceName = os.Getenv("SERVICE_NAME")
)

func main() {
	// 用户仓库 db 接口实现
	repo := &service.Repository{db.Engine}

	srv := micro.NewService(
		micro.Name(serviceName),
		micro.Version("latest"),
	)
	srv.Init()
	// Register handler
	pb.RegisterGoodsHandler(srv.Server(), &hander{repo})
	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
	log.Println("serviser run ...")
}
