package config

import (
	"log"
	"testing"

	"github.com/11me/wb_service/config"
)

func TestParseConfig(t *testing.T) {
	path := "./config.json"
	config, err := config.ParseConfig(path)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(config)
}
