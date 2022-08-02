package services

import (
	"mindustrybp/models"
)

type ServiceGroup struct {
	DB
	S2I
}

type DB interface {
	UpdateSchematic(schematic models.Schematic) (models.Schematic, error)
	InsertSchematic(schematic models.Schematic) (models.Schematic, error)
	GetSchematics() ([]models.Schematic, error)
	GetSchematic(id int) (models.Schematic, error)
}

type S2I interface {
	GenerateImage(schematic string) (string, error)
}
