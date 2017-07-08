package api

type Response struct {
	Host     string `json:Host`
	Password string `json:"Password"`
	Port     string `json:Port`
}

func Request(zoneName string, ServerId string) (*Response, err) {
	return response, err
}
