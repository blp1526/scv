package config

import (
	"encoding/json"
	"io/ioutil"
)

type Scv struct {
	AccessToken       string   `json:"access_token"`
	AccessTokenSecret string   `json:"access_token_secret"`
	Servers           []Server `json:"servers"`
}

type Server struct {
	Name string `json:"name"`
	Zone string `json:"zone"`
	ID   string `json:"id"`
}

func Load(scv *Scv, filePath string) error {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, scv)
	return err
}
