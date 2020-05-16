package models

import "context"

// Datastore provides interface to communicate with db
type Datastore interface {
	GetRandomPlaceByCityName(ctx context.Context, cityName string) (*Place, error)
	GetCities(ctx context.Context, limit uint64, offset uint64) ([]*City, error)
	SavePlace(ctx context.Context, place *Place, cityName string) (uint64, error)
	GetPlacesByCityID(ctx context.Context, cityID uint64, amount uint64, offset uint64) ([]*Place, error)
}
