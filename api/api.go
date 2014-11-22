package api

import "github.com/mono0926/go-qiita/client"

type Api interface {
	ItemClient() client.Item
}

func NewApi(teamKey string, accessToken string) Api {
	auth := client.NewAuth(teamKey, accessToken)
	return &api{auth: auth}
}

type api struct {
	auth *client.Auth
}

func (api api) ItemClient() client.Item {
	return client.NewNotification(api.auth)
}
