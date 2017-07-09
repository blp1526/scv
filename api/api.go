package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os/user"
	"path/filepath"
	"time"

	"github.com/blp1526/scv/config"
	"github.com/blp1526/scv/logger"
)

type Body struct {
	Host     string `json:"Host"`
	Password string `json:"Password"`
	Port     string `json:"Port"`
}

func Request(body interface{}, zoneName string, serverName string) error {
	scv := &config.Scv{}
	current, _ := user.Current()
	dir := filepath.Join(current.HomeDir, "scv.json")
	config.Load(scv, dir)

	if scv.AccessToken == "" || scv.AccessTokenSecret == "" {
		message := fmt.Sprintf("Check scv.json, AccessToken is %s, AccessTokenSecret is %s", scv.AccessToken, scv.AccessTokenSecret)
		return errors.New(message)
	}

	serverId := ""
	for i := 0; i < len(scv.Servers); i++ {
		if scv.Servers[i].ZoneName == zoneName && scv.Servers[i].Name == serverName {
			serverId = scv.Servers[i].ID
		}
	}
	logger.Debug(fmt.Sprintf("Found ServerID is %s", serverId))

	if serverId == "" {
		return errors.New(fmt.Sprintf("ServerID is not found by ZoneName %s and ServerName %s", zoneName, serverName))
	}

	client := &http.Client{Timeout: 10 * time.Second}
	scheme := "https"
	host := "secure.sakura.ad.jp"
	path := "/cloud/zone/" + zoneName + "/api/cloud/1.1/server/" + serverId + "/vnc/proxy"
	url := scheme + "://" + host + path
	logger.Debug("URL is " + url)

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
		message := fmt.Sprintf("Bad response status (got %d, expected 201)", resp.StatusCode)
		return errors.New(message)
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(body)
}
