package mock

import "mindustrybp/models"

type DB struct {
	UpdateSchematicHook func (schematic models.Schematic) (models.Schematic, error)
	InsertSchematicHook func (schematic models.Schematic) (models.Schematic, error)
	GetSchematicsHook func () ([]models.Schematic, error)
	GetSchematicHook func (id int) (models.Schematic, error)
}

func (db *DB) UpdateSchematic(schematic models.Schematic) (models.Schematic, error){
	return db.UpdateSchematicHook(schematic)
}
func (db *DB) InsertSchematic(schematic models.Schematic) (models.Schematic, error){
	return db.InsertSchematicHook(schematic)
}
func (db *DB) GetSchematics() ([]models.Schematic, error){
	return db.GetSchematicsHook()
}
func (db *DB) GetSchematic(id int) (models.Schematic, error){
	return db.GetSchematicHook(id)
}

type S2I struct {
	GenerateImageHook func (schematic string) (string, error)
}

func (s2i *S2I) GenerateImage(schematic string) (string, error){
	return s2i.GenerateImageHook(schematic)
}