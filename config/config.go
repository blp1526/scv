package config

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

func Load() {
}
