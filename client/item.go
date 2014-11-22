package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"github.com/mono0926/go-qiita/response"
)

type Item interface {
	GetItems() []response.Item
}

func NewNotification(auth *Auth) Item {
	return &item{Client: NewClient(auth)}
}

type item struct {
	*Client
}

func (n item) GetItems() []response.Item {
	r := n.Get("items")
	var items []response.Item
	bytes, _ := ioutil.ReadAll(r.Body)
	error := json.Unmarshal(bytes, &items)
	if error != nil {
		log.Fatalln(error)
	}
	return items
}
