package main

import (
	"context"

	gpd "gitee.com/xsjcs/old-sql-service/proto/goods"
	"gitee.com/xsjcs/old-sql-service/service"
)

type hander struct {
	repo service.Repository
}

func (srv *hander) Get(ctx context.Context, req *gpd.Good, res *gpd.Response) (err error) {
	good, err := srv.repo.Get(req)
	if err != nil {
		return err
	}
	res.Good = good
	return err
}
