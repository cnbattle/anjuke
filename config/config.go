package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	IsAll bool
	Sleep int
	Cites []City
}
type City struct {
	Name string
	Url  string
}

var V *Config

func init() {
	if _, err := toml.DecodeFile("./config.toml", &V); err != nil {
		log.Fatal(err)
	}
}
