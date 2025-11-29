package client

import (
	"golang.org/x/oauth2"
	"go_data_gouv_client/config"
)
// a check selon les scopes nécessaires pour ce que je veux
// manière d'automatiser cela ???

func GetOauth2Config(url:) *oauth2.Config {
	cfg := config.GetOauthConfig()
}
