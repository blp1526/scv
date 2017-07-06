package config

import (
	"encoding/json"
	"io/ioutil"
)

type Scv struct {
	AccessToken       string   `json:"access_token"`
	AccessTokenSecret string   `json:access_token_secret`
	Servers           []Server `json:servers`
}

type Server struct {
	Name string `json:"name"`
	Zone string `json:"zone"`
	ID   string `json:"id"`
}

func Load(filePath string) (*Scv, error) {
	var scv Scv

	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return &scv, err
	}

	err = json.Unmarshal(bytes, &scv)
	if err != nil {
		return &scv, err
	}
	return &scv, err
}
