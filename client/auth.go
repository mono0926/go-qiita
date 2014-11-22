package client

type Auth struct {
	TeamKey     string
	AccessToken string
}

func NewAuth(teamKey string, accessToken string) *Auth {
	return &Auth{TeamKey: teamKey, AccessToken: accessToken}
}
