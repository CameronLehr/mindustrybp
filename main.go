package main

import (
	"mindustrybp/config"
	"mindustrybp/server"

	log "github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.New()

	if err != nil {
		log.WithError(err).
			Fatalln("Failed to parse config")
	}

	s, err := server.New(cfg)

	if err != nil {
		log.WithError(err).
			Fatalln("Failed to create server")
	}

	if err = s.Listen(); err != nil {
		log.WithError(err).
			Infoln("Server shutting down")
	}
}
