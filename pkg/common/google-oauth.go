package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/salmanfr/catatann-api/pkg/models"
)

func GetGoogleOAuthToken(code string) (*models.GoogleOAuthToken, error) {
	const root_url = "https://oauth2.googleapis.com/token"

	values := url.Values{}

	values.Add("grant_type", "authorization_code")
	values.Add("code", code)
	values.Add("client_id", os.Getenv("GOOGLE_OAUTH_CLIENT_ID"))
	values.Add("client_secret", os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"))
	values.Add("redirect_uri", os.Getenv("GOOGLE_OAUTH_REDIRECT_URL"))

	fmt.Println("grant_type", values.Get("grant_type"))
	fmt.Println("code", values.Get("code"))
	fmt.Println("client_id", values.Get("client_id"))
	fmt.Println("client_secret", values.Get("client_secret"))
	fmt.Println("redirect_uri", values.Get("redirect_uri"))
	
	query := values.Encode()

	req, err := http.NewRequest("POST", root_url, bytes.NewBufferString(query))
	
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	
	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve token")
	}

	var res_body bytes.Buffer

	_, err = io.Copy(&res_body, res.Body)

	if err != nil {
		return nil, err
	}

	var parsed models.GoogleOAuthToken

	if err := json.Unmarshal(res_body.Bytes(), &parsed); err != nil {
		return nil, err
	}

	return &parsed, nil
}

func GetGoogleUser(access_token string, id_token string) (*models.GoogleUserResult, error) {
	root_url := fmt.Sprintf("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=%s", access_token)

	req, err := http.NewRequest("GET", root_url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", id_token))

	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve user")
	}

	var res_body bytes.Buffer
	
	_, err = io.Copy(&res_body, res.Body)
	
	if err != nil {
		return nil, err
	}
	
	var parsed models.GoogleUserResult

	err = json.Unmarshal(res_body.Bytes(), &parsed)

	if err != nil {
		return nil, err
	}

	return &parsed, nil
}