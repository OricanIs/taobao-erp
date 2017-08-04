package service

import (
	"errors"

	"github.com/goushuyun/taobao-erp/errs"
	"github.com/goushuyun/taobao-erp/misc"

	"golang.org/x/net/context"

	"github.com/goushuyun/taobao-erp/pb"
	"github.com/goushuyun/taobao-erp/stock/db"
	"github.com/wothing/log"
)

func (s *StockServer) SaveGoods(ctx context.Context, req *pb.Goods) (*pb.GoodsResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "SaveGoods", "%#v", req))

	err := db.SaveGoods(req)
	if err != nil {
		log.Error(err)
		return nil, errs.Wrap(errors.New(err.Error()))
	}

	return &pb.GoodsResp{Code: errs.Ok, Message: "ok", Data: []*pb.Goods{req}}, nil
}

// search goods from local db
func (s *StockServer) SearchGoods(ctx context.Context, req *pb.GoodsInfo) (*pb.GoodsInfoListResp, error) {
	tid := misc.GetTidFromContext(ctx)
	defer log.TraceOut(log.TraceIn(tid, "SearchGoods", "%#v", req))
	models, err, totalCount := db.SearchGoods(req)
	if err != nil {
		log.Error(err)
		return nil, errs.Wrap(errors.New(err.Error()))
	}
	return &pb.GoodsInfoListResp{Code: errs.Ok, Message: "ok", Data: models, TotalCount: totalCount}, nil
}
