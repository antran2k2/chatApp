package config

import (
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log"
	"os"
)

var GoogleOAuthConfig *oauth2.Config

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}
	GoogleOAuthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		Endpoint:     google.Endpoint,
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8080/auth/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	}
}
