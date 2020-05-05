package models

import "context"

// Datastore provides interface to communicate with db
type Datastore interface {
	GetRandomItemByCityName(ctx context.Context, cityName string) (*Place, error)
}
