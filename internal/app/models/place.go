package models

// Place provides model for places table in database
type Place struct {
	ID          uint64
	Title       string
	Address     string
	Description string
}
