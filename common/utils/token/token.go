package token

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/coding-yogi/go_bdd/handlers"
	"github.com/coding-yogi/go_bdd/models"
)

//GetOAuthToken ...
func GetOAuthToken() (string, error) {

	env, err := config.GetEnvDetails("qa")
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	//get access token
	req, err := http.NewRequest("POST", env.OAuthURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return "", errors.New("Error generating request for access token")
	}
	req.SetBasicAuth(env.UserName, env.Password)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", errors.New("Error getting response of access token")
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	tokenRes := models.TokenResponse{}
	err = json.Unmarshal(body, &tokenRes)
	if err != nil {
		return "", errors.New("acces token not found in response. Body --> " + string(body))
	}

	return tokenRes.AccessToken, nil
}
