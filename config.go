package umeng

import (
	"encoding/base64"
	"net/http"
)

type Config struct {
	HTTPClient    *http.Client
	Authorization string
}

func NewConfig(email string, password string) *Config {
	return &Config{
		HTTPClient:    http.DefaultClient,
		Authorization: base64.StdEncoding.EncodeToString([]byte(email + ":" + password)),
	}
}
