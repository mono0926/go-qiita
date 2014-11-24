package main

import (
	"fmt"
	"github.com/mono0926/go-qiita/api"
)

func main() {
	teamKey := "TEAM_KEY"
	accessToken := "YOUR_ACCESS_TOKEN"
	a := api.NewTeamApi(teamKey, accessToken)
	nc := a.ItemClient()
	items := nc.GetItems()
	if len(items) == 0 {
		return
	}
	for _, n := range items {
		url := n.Url("")
		fmt.Printf("%s さんが %s を投稿しました。\n%s\n", n.User.Name, n.Title, url)
	}
}
