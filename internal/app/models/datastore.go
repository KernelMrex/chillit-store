package models

import "context"

type Datastore interface {
	GetPlacesById(ctx context.Context, offset uint64, limit uint64) ([]*Place, error)
	AddPlace(ctx context.Context, place *Place) (uint64, error)
}
