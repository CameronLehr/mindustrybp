package routes

import (
	"fmt"
	"mindustrybp/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (r *Routes) CreateSchematic(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	title := req.FormValue("title")
	creator := req.FormValue("creator")
	description := req.FormValue("description")
	schematic := req.FormValue("schematic")
	category := req.FormValue("category")

	test := models.Schematic{
		Title:       title,
		Creator:     creator,
		Description: description,
		Schematic:   schematic,
		Category:    category,
	}
	image, err := r.GenerateImage(test.Schematic)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	test.SchematicImage = image
	new, err := r.InsertSchematic(test)

	// TODO: The schematic may just not exist, or this could be a real db failure
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	http.Redirect(w, req, "/"+strconv.Itoa(int(new.ID)), http.StatusFound)
}

func (r *Routes) EditSchematic(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	params := mux.Vars(req)
	id := params["id"]
	idInt, err := strconv.Atoi(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	title := req.FormValue("title")
	creator := req.FormValue("creator")
	description := req.FormValue("description")
	schematic := req.FormValue("schematic")
	schematicImage := req.FormValue("schematicImage")
	category := req.FormValue("category")
	fmt.Println(title, creator, description, schematic, schematicImage, category)
	test := models.Schematic{
		ID:             int64(idInt),
		Title:          title,
		Creator:        creator,
		Description:    description,
		Schematic:      schematic,
		SchematicImage: schematicImage,
		Category:       category,
	}

	_, err = r.UpdateSchematic(test)

	// TODO: The schematic may just not exist, or this could be a real db failure
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	http.Redirect(w, req, "/", http.StatusFound)
}
