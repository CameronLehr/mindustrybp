package routes

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Render a template given a Page
func (r *Routes) renderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	err := r.templates.Execute(w, tmpl, p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (r *Routes) View(w http.ResponseWriter, req *http.Request) {
	schems, err := r.GetSchematics()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	r.renderTemplate(w, "view.html", schems)
}

func (r *Routes) Create(w http.ResponseWriter, req *http.Request) {
	r.renderTemplate(w, "create.html", nil)
}

func (r *Routes) EditTemplate(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	idInt, err := strconv.Atoi(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	schematic, err := r.GetSchematic(idInt)

	// TODO: The schematic may just not exist, or this could be a real db failure
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	r.renderTemplate(w, "edit.html", schematic)
}

func (r *Routes) SingleSchematic(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	idInt, err := strconv.Atoi(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	schematic, err := r.GetSchematic(idInt)

	// TODO: The schematic may just not exist, or this could be a real db failure
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	r.renderTemplate(w, "schematic.html", schematic)
}
