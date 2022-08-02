package models

import "time"

type Schematic struct {
	ID             int64
	Title          string    //Title of the Page for the submission
	Creator        string    //Submitter's Username
	Description    string    //Authors description of what the schematic is for
	Schematic      string    //base64 of the schematic
	TimeUploaded   time.Time //The date / time that this was submitted
	LastUpdated    time.Time //If this schematic has been updated this will capture the day it was
	SchematicImage string    //idk what this is yet
	Likes          int       //Number of times other user's have liked this schematic
	Downloads      int       //Number of times this schematic has been downloaded
	Category       string    //Idk what this is going to look like just yet
}
