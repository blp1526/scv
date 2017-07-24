package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func (config *Config) GetServerID(zoneName string, serverName string) (id string, err error) {
	for i := 0; i < len(config.Servers); i++ {
		if config.Servers[i].ZoneName == zoneName && config.Servers[i].Name == serverName {
			id = config.Servers[i].ID
		}
	}
	if id == "" {
		return id, fmt.Errorf("ServerID is not found by zoneName: %s, serverName: %s", zoneName, serverName)
	} else {
		return id, nil
	}
}
