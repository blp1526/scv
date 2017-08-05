package scv

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type API struct {
	ZoneName          string
	ServerID          string
	AccessToken       string
	AccessTokenSecret string
	Logger            Logger
}

func (api *API) URL() (url string) {
	scheme := "https"
	host := "secure.sakura.ad.jp"
	path := "/cloud/zone/" + api.ZoneName + "/api/cloud/1.1/server/" + api.ServerID + "/vnc/proxy"
	url = scheme + "://" + host + path
	return url
}

func (api *API) GetServerAddress() (serverAddress string, err error) {
	url := api.URL()
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
	vnc := &VNC{}
	json.NewDecoder(resp.Body).Decode(vnc)
	serverAddress = vnc.Path()
	return serverAddress, err
}
