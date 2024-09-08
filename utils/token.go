package utils

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/go-mpesa-integration/config"
	"github.com/sirupsen/logrus"
)

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expiry_in"`
}

func GenerateAccessToken(consumerKey string, consumerSecret string) (string, error) {
	config := config.NewEnvConfig()
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method,config.MPesaTokenUrl, nil)

	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(consumerKey+":"+consumerSecret)))
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logrus.Errorf("error closing body: %v", err)
		}
	}(res.Body)

	var accessTokenResponse AccessTokenResponse
	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	// Unmarshal the JSON response into the struct
	err = json.Unmarshal(body, &accessTokenResponse)
	if err != nil {
		return "", err
	}
	return accessTokenResponse.AccessToken, nil
}


func  GenerateTimeStamp() string {
	return time.Now().Format("20060102150405")
}

func GeneratePassword(businessshortcode string , passKey string, timeStamp string) string {
	return base64.StdEncoding.EncodeToString([]byte(businessshortcode + passKey + timeStamp))
}