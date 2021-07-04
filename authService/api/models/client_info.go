package models

import (
	"auth-service/api/caches"
	"auth-service/services"
	"context"
)

type ClientInfo struct {
	ClientId string
	Callback string
}

type AuthInfoFunc struct {
	table ClientInfo
}

func (c AuthInfoFunc) GetOneById(ctx context.Context, id int64) (ClientInfo, error) {
	clientInfo, putCb, err := caches.ClientCacheFunc().GetOneById(ctx, id)
	if err == nil {
		return clientInfo, err
	}
	err = services.Mysql().Session(nil).First(&clientInfo, ClientInfo{}).Error
	if err != nil {
		return clientInfo, err
	}
	putCb(clientInfo)
	return clientInfo, err
}
