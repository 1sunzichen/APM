package api

import (
	"context"
	"dogapm"
	"net/http"
	"ordersvc/grpclient"
	protos "proto"
	"strconv"

	"github.com/google/uuid"
)

type order struct {
}

var Order = &order{}

func (o *order) Add(w http.ResponseWriter, r *http.Request) {
	value := r.URL.Query()
	var (
		uid, _   = strconv.Atoi(value.Get("uid"))
		skuid, _ = strconv.Atoi(value.Get("skuid"))
		num, _   = strconv.Atoi(value.Get("num"))
	)
	//检查用户信息
	_, err := grpclient.UserClient.GetUser(context.TODO(), &protos.User{Id: int64(uid)})
	if err != nil {
		dogapm.Logger.Error(context.TODO(), "add order", map[string]interface{}{"err": err.Error(), "uid": uid, "skuid": skuid, "num": num})
		dogapm.HttpStatus.Error(w, err.Error(), nil)
		return
	}
	//对库存进行扣减
	skuMsg, err := grpclient.SkuClient.DecreaseStock(context.TODO(), &protos.Sku{Id: int64(skuid), Num: int32(num)})
	if err != nil {
		dogapm.Logger.Error(context.TODO(), "add order", map[string]interface{}{"err": err.Error(), "uid": uid, "skuid": skuid, "num": num})
		dogapm.HttpStatus.Error(w, err.Error(), nil)
		return
	}
	//创建订单
	_, err = dogapm.Infra.Db.ExecContext(context.TODO(), "insert into t_order(order_id,sku_id,num,price,uid) values(?,?,?,?,?)", uuid.New().String(), skuid, num, int(skuMsg.Price)*num, uid)
	//do something
	if err != nil {
		dogapm.Logger.Error(context.TODO(), "add order", map[string]interface{}{"err": err.Error(), "uid": uid, "skuid": skuid, "num": num})
		dogapm.HttpStatus.Error(w, err.Error(), nil)
		return
	}

	dogapm.HttpStatus.Ok(w)
}
