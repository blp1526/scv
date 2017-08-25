package scv

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	AccessToken       string   `json:"access_token"`
	AccessTokenSecret string   `json:"access_token_secret"`
	Servers           []Server `json:"servers"`
}

type Server struct {
	Name     string `json:"name"`
	ZoneName string `json:"zone_name"`
	ID       string `json:"id"`
}

func (config *Config) LoadFile(dir string) (err error) {
	bytes, err := ioutil.ReadFile(dir)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, config)
	return err
}

func (config *Config) ServerID(zoneName string, serverName string) (id string, err error) {
	for i := 0; i < len(config.Servers); i++ {
		if config.Servers[i].ZoneName == zoneName && config.Servers[i].Name == serverName {
			id = config.Servers[i].ID
		}
	}
	if id == "" {
		return id, fmt.Errorf("ServerID is not found by zoneName: %s, serverName: %s", zoneName, serverName)
	} else {
		return id, err
	}
}

func (config *Config) CreateFile(dir string) (result string, err error) {
	_, err = os.Stat(dir)
	if err == nil {
		return result, fmt.Errorf("Already you have %s", dir)
	}

	j, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return result, err
	}

	err = ioutil.WriteFile(dir, j, 0600)
	if err != nil {
		return result, err
	}

	return fmt.Sprintf("%s created", dir), err
}
