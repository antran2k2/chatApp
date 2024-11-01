package auth

import (
	"context"
	"encoding/json"
	"golang.org/x/oauth2"
	_ "net/http"
)

type UserInfo struct {
	Email string `json:"email"`
}

func GetGoogleUserInfo(ctx context.Context, token *oauth2.Token) (*UserInfo, error) {
	client := oauth2.NewClient(ctx, oauth2.StaticTokenSource(token))
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	userInfo := &UserInfo{}
	if err := json.NewDecoder(resp.Body).Decode(userInfo); err != nil {
		return nil, err
	}
	return userInfo, nil
}
