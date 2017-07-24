package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/user"
	"path/filepath"
	"time"

	"github.com/blp1526/scv/conf"
)

type Body struct {
	Host     string `json:"Host"`
	Password string `json:"Password"`
	Port     string `json:"Port"`
}

func Request(body *Body, zoneName string, serverName string) (err error) {
	current, _ := user.Current()
	dir := filepath.Join(current.HomeDir, "scv.json")

	config := &conf.Config{}
	err = config.LoadFile(dir)
	if err != nil {
		return err
	}

	serverID, err := config.GetServerID(zoneName, serverName)
	if err != nil {
		return err
	}

	client := &http.Client{Timeout: 10 * time.Second}
	scheme := "https"
	host := "secure.sakura.ad.jp"
	path := "/cloud/zone/" + zoneName + "/api/cloud/1.1/server/" + serverID + "/vnc/proxy"
	url := scheme + "://" + host + path

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(config.AccessToken, config.AccessTokenSecret)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	// NOTE: not 200
	if resp.StatusCode != 201 {
		return fmt.Errorf("Bad response status (got %d, expected 201)", resp.StatusCode)
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(body)
}
