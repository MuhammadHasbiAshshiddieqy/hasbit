package model

import (
	"github.com/google/uuid"
)

// Log - log every external request
type Log struct{
	UserID			uuid.UUID	`json:"user_id"`
	URL				string		`json:"url"`
	Endpoint		string		`json:"endpoint"`
	Error			string		`json:"error"`
}