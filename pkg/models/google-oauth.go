package models

type GoogleOAuthToken struct {
	AccessToken string `json:"access_token"`
	IDToken string `json:"id_token"`
	ExpiresIn int `json:"expires_in"`
}

type GoogleUserResult struct {
	Id             string `json:"id"`
	Email          string `json:"email"`
	Verified_email bool `json:"verified_email"`
	Name           string `json:"name"`
	Given_name     string `json:"given_name"`
	Family_name    string `json:"family_name"`
	Picture        string `json:"picture"`
	Locale         string `json:"locale"`
}