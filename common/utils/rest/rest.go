package rest

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Header ..
type Header struct {
	Key   string
	Value string
}

// GenerateRequest ...
func GenerateRequest(method, url string, body []byte, headers []Header) *http.Request {
	req, _ := http.NewRequest(method, url, bytes.NewReader(body))
	//Add headers
	for _, header := range headers {
		req.Header.Add(header.Key, header.Value)
	}

	return req
}

// ExecuteRequestAndGetResponse ...
func ExecuteRequestAndGetResponse(req *http.Request, client *http.Client) (*http.Response, error) {
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("Error getting response of access token")
	}

	return res, nil
}

// GetResponseBody ...
func GetResponseBody(res *http.Response) []byte {
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return body
}
