package db

import (
	"database/sql"
	"log"
	"mindustrybp/models"
	"time"
)

type DB struct {
	db *sql.DB
}

func New(db *sql.DB) (*DB, error) {
	ndb := &DB{
		db,
	}

	if err := ndb.createSchematicTable(); err != nil {
		return nil, err
	}

	return ndb, nil
}

func (db *DB) InsertSchematic(schematic models.Schematic) (models.Schematic, error) {
	insertSchematicSQL := `INSERT INTO schematics (title,creator,description,schematic,timeUploaded,lastUpdated,schematicImage,likes,downloads,category)
	VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`
	_, err := db.db.Exec(insertSchematicSQL, schematic.Title, schematic.Creator, schematic.Description, schematic.Schematic, schematic.TimeUploaded, schematic.LastUpdated, schematic.SchematicImage, schematic.Likes, schematic.Downloads, schematic.Category)
	// TODO: Assign returned ID before returning
	return schematic, err
}

func (db *DB) GetSchematics() ([]models.Schematic, error) {
	getSchematicsSQL := `SELECT * FROM schematics;`
	rows, err := db.db.Query(getSchematicsSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	schematics := []models.Schematic{}
	for rows.Next() {
		tmp := models.Schematic{}
		timeUpload := ""
		lastUpdated := ""
		err = rows.Scan(&tmp.ID, &tmp.Title, &tmp.Creator, &tmp.Description, &tmp.Schematic, &timeUpload, &lastUpdated, &tmp.SchematicImage, &tmp.Likes, &tmp.Downloads, &tmp.Category)
		if err != nil {
			return nil, err
		}
		tmp.TimeUploaded, _ = time.Parse("2006-01-02T15:04:05-0700", timeUpload)
		tmp.LastUpdated, _ = time.Parse("2006-01-02T15:04:05-0700", lastUpdated)
		schematics = append(schematics, tmp)
	}
	return schematics, nil
}

func (db *DB) GetSchematic(id int) (models.Schematic, error) {
	getSchematicsSQL := `SELECT * FROM schematics WHERE id = $1`
	rows, err := db.db.Query(getSchematicsSQL, id)
	if err != nil {
		return models.Schematic{}, err
	}
	defer rows.Close()
	log.Println("Rows:", rows)
	schematic := models.Schematic{}
	log.Println("models.Schematic:", schematic)
	return schematic, nil
}

func (db *DB) UpdateSchematic(schematic models.Schematic) (models.Schematic, error) {
	stmt, err := db.db.Prepare("UPDATE schematics SET title=?, creator=?, description=?, schematic=?, schematicImage=?, category=? WHERE id=?")
	if err != nil {
		return models.Schematic{}, err
	}
	_, err = stmt.Exec(schematic.Title, schematic.Creator, schematic.Description, schematic.Schematic, schematic.SchematicImage, schematic.Category, schematic.ID)
	return schematic, err
}

func (db *DB) DeleteSchematic(id int) error {
	stmt, err := db.db.Prepare("DELETE FROM schematics WHERE id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	return err
}

func (db *DB) SearchSchematic(search string) ([]models.Schematic, error) {
	getSchematicsSQL, err := db.db.Prepare(`SELECT * FROM schematics WHERE title LIKE '%?%';`)
	if err != nil {
		return nil, err
	}
	rows, err := getSchematicsSQL.Query(search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	schematics := []models.Schematic{}
	for rows.Next() {
		tmp := models.Schematic{}
		timeUpload := ""
		lastUpdated := ""
		err = rows.Scan(&tmp.ID, &tmp.Title, &tmp.Creator, &tmp.Description, &tmp.Schematic, &timeUpload, &lastUpdated, &tmp.SchematicImage, &tmp.Likes, &tmp.Downloads, &tmp.Category)
		if err != nil {
			return nil, err
		}
		tmp.TimeUploaded, _ = time.Parse("2006-01-02T15:04:05-0700", timeUpload)
		tmp.LastUpdated, _ = time.Parse("2006-01-02T15:04:05-0700", lastUpdated)
		schematics = append(schematics, tmp)
	}
	return schematics, nil
}

//Create Table (Only use for initial build of DB)
func (db *DB) createSchematicTable() error {
	schematicTable := `CREATE TABLE IF NOT EXISTS schematics (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"title" TEXT NOT NULL, 
		"creator" TEXT NOT NULL,
		"description" TEXT NOT NULL,
		"schematic" TEXT NOT NULL,
		"timeUploaded" TEXT NOT NULL,
		"lastUpdated" TEXT NOT NULL,
		"schematicImage" TEXT NOT NULL,
		"likes" integer NOT NULL,
		"downloads" integer NOT NULL,
		"category" integer NOT NULL
		);`
	_, err := db.db.Exec(schematicTable)
	return err
}
