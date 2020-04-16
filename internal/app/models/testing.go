package models

import (
	"context"
	"errors"
)

type MockDB struct {
	Places []*Place
}

func NewMockDB() *MockDB {
	return &MockDB{
		Places: []*Place{},
	}
}

func (db *MockDB) GetPlacesById(ctx context.Context, offset uint64, limit uint64) ([]*Place, error) {
	if uint64(len(db.Places)) < offset+limit {
		return nil, errors.New("please select correct range for testing")
	}
	return db.Places[offset:limit], nil
}

func (db *MockDB) AddPlace(ctx context.Context, place *Place) (uint64, error) {
	db.Places = append(db.Places, &Place{
		Id:          uint64(len(db.Places)),
		Title:       place.Title,
		Address:     place.Address,
		Description: place.Description,
	})
	return uint64(len(db.Places) - 1), nil
}
