package caches

import (
	"auth-service/api/models"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

const (
	ClientInfoCacheKey = "auth_service" + "_client_info_"
	CacheMinExpire     = 60 * 60
)

var _ClientCacheFunc = clientCacheFunc{}

type clientCacheFunc struct {
	cli redis.ClusterClient
}

func GenCacheById(id int64) string {
	return ClientInfoCacheKey + strconv.FormatInt(id, 10)
}

func (c clientCacheFunc) GetOneById(ctx context.Context, id int64) (res models.ClientInfo, putCb func(res models.ClientInfo), err error) {
	putCb = func(res models.ClientInfo) {
		c.cli.Set(ctx, GenCacheById(id), "ok", time.Second*CacheMinExpire)
	}
	cmd := c.cli.Get(ctx, GenCacheById(id))
	if cmd.Err() != nil {
		return res, nil, cmd.Err()
	}
	err = json.Unmarshal([]byte(cmd.String()), &res)
	return
}

func ClientCacheFunc() clientCacheFunc {
	return _ClientCacheFunc
}
