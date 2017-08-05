package scv

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type API struct {
	Logger Logger
}

func (api *API) URL(zoneName string, serverID string) (url string) {
	scheme := "https"
	host := "secure.sakura.ad.jp"
	path := "/cloud/zone/" + zoneName + "/api/cloud/1.1/server/" + serverID + "/vnc/proxy"
	url = scheme + "://" + host + path
	return url
}

func (api *API) GET(url string, accessToken string, accessTokenSecret string) (body []byte, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return body, err
	}
	req.SetBasicAuth(accessToken, accessTokenSecret)
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return body, err
	}
	// NOTE: not 200
	if resp.StatusCode != 201 {
		return body, fmt.Errorf("Bad response status (got %d, expected 201)", resp.StatusCode)
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return body, err
}
