package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/user"
	"path/filepath"
	"time"

	"github.com/blp1526/scv/config"
)

type Body struct {
	Host     string `json:"Host"`
	Password string `json:"Password"`
	Port     string `json:"Port"`
}

func Request(body *Body, zoneName string, serverName string) error {
	scv := &config.Scv{}
	current, _ := user.Current()
	dir := filepath.Join(current.HomeDir, "scv.json")
	config.Load(scv, dir)

	if scv.AccessToken == "" || scv.AccessTokenSecret == "" {
		return fmt.Errorf("Check scv.json, AccessToken is %s, AccessTokenSecret is %s", scv.AccessToken, scv.AccessTokenSecret)
	}

	serverId := ""
	for i := 0; i < len(scv.Servers); i++ {
		if scv.Servers[i].ZoneName == zoneName && scv.Servers[i].Name == serverName {
			serverId = scv.Servers[i].ID
		}
	}

	if serverId == "" {
		return fmt.Errorf("ServerID is not found by ZoneName %s and ServerName %s", zoneName, serverName)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	scheme := "https"
	host := "secure.sakura.ad.jp"
	path := "/cloud/zone/" + zoneName + "/api/cloud/1.1/server/" + serverId + "/vnc/proxy"
	url := scheme + "://" + host + path

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(scv.AccessToken, scv.AccessTokenSecret)
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
