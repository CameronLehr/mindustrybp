package server

import (
	"database/sql"
	"net/http"

	log "github.com/sirupsen/logrus"

	_ "modernc.org/sqlite"

	"mindustrybp/config"
	"mindustrybp/routes"
)

type Server struct {
	cfg *config.Config
	r   *routes.Routes
}

func New(cfg *config.Config) (*Server, error) {
	if cfg.Debug {
		log.SetLevel(log.DebugLevel)
	}
	db, err := sql.Open("sqlite", cfg.Sqlite)

	if err != nil {
		return nil, err
	}

	r, err := routes.New(cfg, db)

	if err != nil {
		return nil, err
	}

	return &Server{
		cfg,
		r,
	}, nil
}

func (s *Server) Listen() error {
	log.WithField("Addr", s.cfg.ListenAddr).
		Infoln("Server listening...")
	return http.ListenAndServe(s.cfg.ListenAddr, s.r.Router)
}
