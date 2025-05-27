package service

import (
	confid "auth-service/internal/config"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleOAuthConfig *oauth2.Config

func InitGoogleOAuth(cfg confid.GoogleOAuth) {
	GoogleOAuthConfig = &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RedirectURL:  cfg.RedirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}
