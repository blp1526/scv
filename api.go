package scv

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Api struct {
	ZoneName          string
	ServerID          string
	AccessToken       string
	AccessTokenSecret string
	Logger            Logger
}

type Body struct {
	Host     string `json:"Host"`
	Password string `json:"Password"`
	Port     string `json:"Port"`
}

func (api *Api) GetServerAddress() (serverAddress string, err error) {
	scheme := "https"
	host := "secure.sakura.ad.jp"
	path := "/cloud/zone/" + api.ZoneName + "/api/cloud/1.1/server/" + api.ServerID + "/vnc/proxy"
	url := scheme + "://" + host + path
	api.Logger.Debug("URL: " + url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return serverAddress, err
	}

	req.SetBasicAuth(api.AccessToken, api.AccessTokenSecret)
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return serverAddress, err
	}

	// NOTE: not 200
	if resp.StatusCode != 201 {
		return serverAddress, fmt.Errorf("Bad response status (got %d, expected 201)", resp.StatusCode)
	}
	defer resp.Body.Close()
	body := &Body{}
	json.NewDecoder(resp.Body).Decode(body)
	serverAddress = "vnc://:" + body.Password + "@" + body.Host + ":" + body.Port
	return serverAddress, err
}
