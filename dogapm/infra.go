package dogapm

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

// infrastructure 基础措施
type infra struct {
	Db  *sql.DB
	Rdb *redis.Client
}

var Infra = &infra{}

type InfraOption func(i *infra)

func InfraDbOption(connectUrl string) InfraOption {
	return func(i *infra) {
		var err error
		i.Db, err = sql.Open("mysql", connectUrl)
		if err != nil {
			panic(err)
		}
		err = i.Db.Ping()
		if err != nil {
			panic(err)
		}
	}
}

func InfraRDBOption(addr string) InfraOption {
	return func(i *infra) {
		Rdb := redis.NewClient(&redis.Options{
			Addr: addr,
			DB:   0,
		})
		res, err := Rdb.Ping(context.TODO()).Result()
		if err != nil {
			panic(err)
		}
		if res != "PONG" {
			panic("redis connect error")
		}
		i.Rdb = Rdb
	}
}

func (i *infra) Init(options ...InfraOption) {
	for _, option := range options {
		option(i)
	}
}
