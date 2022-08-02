package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"

	log "github.com/sirupsen/logrus"
)

// A Config holds all configurable settings from a yml config file
type Config struct {
	ListenAddr string
	Sqlite     string
	Debug      bool

	Cookies struct {
		Key    string
		Domain string
	}
}

// New uses a file path stored in the CONFIG environment variable to populate the Config struct
func New() (*Config, error) {
	path := os.Getenv("CONFIG")

	if len(path) <= 0 {
		path = "config.yml"
	}

	data, err := ioutil.ReadFile(path)

	if err != nil {
		log.WithField("Path", path).WithError(err).Errorln("Failed to read config file")
		return nil, err
	}

	cfg := &Config{}

	err = yaml.Unmarshal(data, &cfg)

	if err != nil {
		log.WithField("Data", data).WithError(err).Errorln("Failed to unmarshal config data")
		return nil, err
	}

	//TODO: Validate Fields

	return cfg, nil
}
