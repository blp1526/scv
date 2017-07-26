package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Vnc struct {
	ZoneName          string
	ServerID          string
	AccessToken       string
	AccessTokenSecret string
}

type Body struct {
	Host     string `json:"Host"`
	Password string `json:"Password"`
	Port     string `json:"Port"`
}

func (vnc *Vnc) GetServerAddress() (serverAddress string, err error) {
	scheme := "https"
	host := "secure.sakura.ad.jp"
	path := "/cloud/zone/" + vnc.ZoneName + "/api/cloud/1.1/server/" + vnc.ServerID + "/vnc/proxy"
	url := scheme + "://" + host + path

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return serverAddress, err
	}

	req.SetBasicAuth(vnc.AccessToken, vnc.AccessTokenSecret)
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
