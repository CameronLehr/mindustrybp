package routes

import (
	"database/sql"
	"html/template"
	"math"
	"mindustrybp/config"
	"mindustrybp/services"
	dbs "mindustrybp/services/db"
	"mindustrybp/services/s2i"
	"net/http"
	"strconv"

	"github.com/gernest/hot"
	"github.com/gorilla/mux"
)

type Routes struct {
	cfg *config.Config
	db  *sql.DB

	Router *mux.Router

	templates *hot.Template

	services.ServiceGroup
}

func New(cfg *config.Config, db *sql.DB) (*Routes, error) {
	templateMap := template.FuncMap{
		"ToHumanReadable": ToHumanReadable,
	}

	config := &hot.Config{
		Watch:          true,
		BaseName:       "templates",
		Dir:            "templates",
		FilesExtension: []string{".html"},
		Funcs:          templateMap,
	}

	templates, err := hot.New(config)

	if err != nil {
		return nil, err
	}

	r := &Routes{
		cfg:       cfg,
		db:        db,
		Router:    mux.NewRouter(),
		templates: templates,
	}

	sg, err := DefaultGroup(db)

	if err != nil {
		return nil, err
	}

	r.ServiceGroup = sg

	// VIEWS
	r.Router.HandleFunc("/", r.View).Methods(http.MethodGet)
	r.Router.HandleFunc("/create", r.Create).Methods(http.MethodGet)
	r.Router.HandleFunc("/{id}", r.SingleSchematic).Methods(http.MethodGet)
	r.Router.HandleFunc("/edit/{id}", r.Create).Methods(http.MethodGet)

	// API
	r.Router.HandleFunc("/schematics", r.CreateSchematic).Methods(http.MethodPost)
	r.Router.HandleFunc("/schematics/{id}", r.EditSchematic).Methods(http.MethodPatch)

	// STATIC
	r.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	return r, nil
}

func ToHumanReadable(num int) string {
	hundred := 3
	thousand := 6
	million := 9
	billion := 12
	magnitude := math.Floor(math.Log10(float64(num)))
	result := ""

	switch {
	case int(magnitude) < hundred:
		result = strconv.Itoa(num)
	case int(magnitude) < thousand:
		result = strconv.Itoa(num/1000) + "k"
	case int(magnitude) < million:
		result = strconv.Itoa(num/1000000) + "m"
	case int(magnitude) < billion:
		result = strconv.Itoa(num/1000000000) + "b"
	default:
		result = "alot"
	}
	return result
}

func DefaultGroup(db *sql.DB) (services.ServiceGroup, error) {

	sg := services.ServiceGroup{}

	dbg, err := dbs.New(db)

	if err != nil {
		return services.ServiceGroup{}, err
	}

	sg.DB = dbg
	sg.S2I = s2i.New()

	return sg, nil
}
