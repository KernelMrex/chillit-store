package models

type Datastore interface {
	GetPlacesById(offset uint64, limit uint16) ([]Place, error)
}
