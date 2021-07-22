package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	DBUri  string `json:"db_uri"`
	DBName string `json:"db_name"`
	DBUser string `json:"db_user"`
	DBPass string `json:"db_pass"`
}

// ParseConfig parses the JSON config
// and returns pointer to the struct
func ParseConfig(path string) (*Config, error) {
	var config Config

	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
