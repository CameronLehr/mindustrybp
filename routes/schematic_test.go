package routes

import (
	"errors"
	"mindustrybp/models"
	"mindustrybp/services"
	"mindustrybp/services/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestRoutes_CreateSchematic(t *testing.T) {
	tests := []struct {
		ServiceGroup   services.ServiceGroup
		name           string
		expectedStatus int
		formValues     map[string][]string
	}{
		{
			name:           "success",
			expectedStatus: http.StatusFound,
			formValues: map[string][]string{
				"title":       {"test"},
				"creator":     {"test"},
				"description": {"test"},
				"schematic":   {"test"},
				"category":    {"test"},
			},
			ServiceGroup: services.ServiceGroup{
				DB: &mock.DB{
					InsertSchematicHook: func(schematic models.Schematic) (models.Schematic, error) {
						return schematic, nil
					},
				},
				S2I: &mock.S2I{
					GenerateImageHook: func(schematic string) (string, error) {
						return "", nil
					},
				},
			},
		},
		{
			name:           "handle invalid form",
			expectedStatus: http.StatusBadRequest,
			formValues:     nil,
		},
		{
			name:           "handle generateImage error",
			expectedStatus: http.StatusBadRequest,
			formValues: map[string][]string{
				"title":       {"test"},
				"creator":     {"test"},
				"description": {"test"},
				"schematic":   {"test"},
				"category":    {"test"},
			},
			ServiceGroup: services.ServiceGroup{
				S2I: &mock.S2I{
					GenerateImageHook: func(schematic string) (string, error) {
						return "", errors.New("Test")
					},
				},
			},
		},
		{
			name:           "handle insertSchematic error",
			expectedStatus: http.StatusInternalServerError,
			formValues: map[string][]string{
				"title":       {"test"},
				"creator":     {"test"},
				"description": {"test"},
				"schematic":   {"test"},
				"category":    {"test"},
			},
			ServiceGroup: services.ServiceGroup{
				DB: &mock.DB{
					InsertSchematicHook: func(schematic models.Schematic) (models.Schematic, error) {
						return schematic, errors.New("Test")
					},
				},
				S2I: &mock.S2I{
					GenerateImageHook: func(schematic string) (string, error) {
						return "", nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Routes{
				ServiceGroup: tt.ServiceGroup,
			}
			req, _ := http.NewRequest("POST", "/schematics", nil)
			req.PostForm = tt.formValues
			w := httptest.NewRecorder()
			r.CreateSchematic(w, req)
			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}

func TestRoutes_EditSchematic(t *testing.T) {
	tests := []struct {
		ServiceGroup   services.ServiceGroup
		name           string
		expectedStatus int
		formValues     map[string][]string
		vars           map[string]string
	}{
		{
			name:           "success",
			expectedStatus: http.StatusFound,
			formValues: map[string][]string{
				"title":       {"test"},
				"creator":     {"test"},
				"description": {"test"},
				"schematic":   {"test"},
				"category":    {"test"},
			},
			vars: map[string]string{
				"id": "1",
			},
			ServiceGroup: services.ServiceGroup{
				DB: &mock.DB{
					UpdateSchematicHook: func(schematic models.Schematic) (models.Schematic, error) {
						return schematic, nil
					},
				},
				S2I: &mock.S2I{
					GenerateImageHook: func(schematic string) (string, error) {
						return "", nil
					},
				},
			},
		},
		{
			name:           "handle invalid form",
			expectedStatus: http.StatusBadRequest,
			formValues:     nil,
		},
		{
			name:           "handle invalid vars",
			expectedStatus: http.StatusBadRequest,
			formValues: map[string][]string{
				"title":       {"test"},
				"creator":     {"test"},
				"description": {"test"},
				"schematic":   {"test"},
				"category":    {"test"},
			},
			vars: nil,
		},
		{
			name:           "handle generateImage error",
			expectedStatus: http.StatusBadRequest,
			formValues: map[string][]string{
				"title":       {"test"},
				"creator":     {"test"},
				"description": {"test"},
				"schematic":   {"test"},
				"category":    {"test"},
			},
			vars: map[string]string{
				"id": "1",
			},
			ServiceGroup: services.ServiceGroup{
				S2I: &mock.S2I{
					GenerateImageHook: func(schematic string) (string, error) {
						return "", errors.New("Test")
					},
				},
			},
		},
		{
			name:           "handle insertSchematic error",
			expectedStatus: http.StatusInternalServerError,
			formValues: map[string][]string{
				"title":       {"test"},
				"creator":     {"test"},
				"description": {"test"},
				"schematic":   {"test"},
				"category":    {"test"},
			},
			vars: map[string]string{
				"id": "1",
			},
			ServiceGroup: services.ServiceGroup{
				DB: &mock.DB{
					UpdateSchematicHook: func(schematic models.Schematic) (models.Schematic, error) {
						return schematic, errors.New("Test")
					},
				},
				S2I: &mock.S2I{
					GenerateImageHook: func(schematic string) (string, error) {
						return "", nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Routes{
				ServiceGroup: tt.ServiceGroup,
			}
			req, _ := http.NewRequest("PATCH", "/schematics", nil)
			req.PostForm = tt.formValues
			w := httptest.NewRecorder()
			req = mux.SetURLVars(req, tt.vars)
			r.EditSchematic(w, req)
			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}
