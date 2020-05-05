package models

import "context"

// Datastore provides interface to communicate with db
type Datastore interface {
	GetRandomPlaceByCityName(ctx context.Context, cityName string) (*Place, error)
	GetCities(ctx context.Context, limit uint64, offset uint64) ([]*City, error)
}
