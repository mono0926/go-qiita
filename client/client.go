package client

import (
	"fmt"
	"net/http"
	"log"
)

type Client struct {
	Auth *Auth
}

func (c Client) Get(resource string) *http.Response {
	url := fmt.Sprintf("http://%sqiita.com/api/v2/%s", c.Auth.TeamKey, resource)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Auth.AccessToken))
	client := http.Client{}
	r, e := client.Do(req)
	if e != nil {
		log.Fatalln(e)
	}
	return r
}

func NewClient(auth *Auth) *Client {
	return &Client{Auth: auth}
}
