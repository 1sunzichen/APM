package dao

import (
	"context"
	"dogapm"
	"encoding/json"
	"fmt"
	"time"
)

type userDao struct {
}

var UserDao = &userDao{}

func (u *userDao) GetUserById(ctx context.Context, id int64) map[string]interface{} {
	userCache, err := dogapm.Infra.Rdb.Get(ctx, fmt.Sprintf("%s:%s:%d", "usersrv", "uinfo", id)).Result()
	if userCache != "" {
		userinfo := make(map[string]interface{})
		err = json.Unmarshal([]byte(userCache), &userinfo)
		if err == nil {
			return userinfo
		}
	}
	userDbinfo := dogapm.DBUtil.QueryFirst(dogapm.Infra.Db.QueryContext(ctx, "select * from t_user where id=?;", id))
	if len(userDbinfo) == 0 {
		return nil
	}
	cacheDbinfo, _ := json.Marshal(userDbinfo)
	if err == nil {
		dogapm.Infra.Rdb.Set(ctx, fmt.Sprintf("%s:%s:%d", "usersrv", "uinfo", id), cacheDbinfo, 10*time.Minute)
	}

	return userDbinfo

}
