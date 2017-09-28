package token

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.wdf.sap.corp/ml-base/lr-bdd-tests/common/utils/rest"
	"github.wdf.sap.corp/ml-base/lr-bdd-tests/handlers"
	"github.wdf.sap.corp/ml-base/lr-bdd-tests/models/api/response"
)

//GetOAuthToken ...
func GetOAuthToken() (string, error) {

	env, err := config.GetEnvDetails("qa")
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	headers := []rest.Header{
		rest.Header{Key: "content-type", Value: "application/x-www-form-urlencoded"},
	}

	req := rest.GenerateRequest("POST", env.OAuthURL, []byte(data.Encode()), headers)
	req.SetBasicAuth(env.UserName, env.Password)

	client := &http.Client{}
	res, err := rest.ExecuteRequestAndGetResponse(req, client)
	if err != nil {
		return "", errors.New("Error executing recommendation request")
	}

	body := rest.GetResponseBody(res)
	tokenRes := responsemodels.TokenResponse{}

	err = json.Unmarshal(body, &tokenRes)
	if err != nil {
		return "", errors.New("acces token not found in response. Body --> " + string(body))
	}

	return tokenRes.AccessToken, nil
}
