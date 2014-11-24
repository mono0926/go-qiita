package response

import (
	"fmt"
	"time"
)

type Item struct {
	Body      string `json:"body"`
	Coediting bool   `json:"coediting"`
	CreatedAt time.Time `json:"created_at"`
	ID        string `json:"id"`
	Private   bool   `json:"private"`
	Tags      []struct {
		IconURL  string        `json:"icon_url"`
		Name     string        `json:"name"`
		Versions []interface{} `json:"versions"`
	} `json:"tags"`
	Title     string `json:"title"`
	UpdatedAt string `json:"updated_at"`
	User      struct {
		Description       string  `json:"description"`
		FacebookID        string  `json:"facebook_id"`
		FolloweesCount    float64 `json:"followees_count"`
		FollowersCount    float64 `json:"followers_count"`
		GithubLoginName   string  `json:"github_login_name"`
		ID                string  `json:"id"`
		ItemsCount        float64 `json:"items_count"`
		LinkedinID        string  `json:"linkedin_id"`
		Location          string  `json:"location"`
		Name              string  `json:"name"`
		Organization      string  `json:"organization"`
		ProfileImageURL   string  `json:"profile_image_url"`
		TwitterScreenName string  `json:"twitter_screen_name"`
		WebsiteURL        string  `json:"website_url"`
	} `json:"user"`
}

func (item Item) Url(teamKey string) string {
	return fmt.Sprintf("https://%sqiita.com/%s/items/%s", teamKey, item.User.ID, item.ID)
}
