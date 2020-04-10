package models

import (
	"context"
	"errors"
	"fmt"
)

type Place struct {
	Id          uint64
	Title       string
	Address     string
	Description string
}

const placesMaxLimit = 20

func (db *MysqlDB) GetPlacesById(ctx context.Context, offset uint64, limit uint64) ([]Place, error) {
	// Creating query
	if placesMaxLimit <= limit || limit <= 0 {
		return nil, errors.New(fmt.Sprintf("[ GetPlacesById ] bad range; must be [%d..%d]", 1, limit))
	}
	rows, err := db.QueryContext(ctx,
		"SELECT 'id', 'title', 'address', 'description' FROM 'place' LIMIT ? OFFSET ?",
		limit,
		offset,
	)
	if err != nil {
		return nil, errors.New("[ GetPlacesById ] could not execute query")
	}

	// Parsing results
	places := make([]Place, 0)
	for rows.Next() {
		var place Place
		if err := rows.Scan(&place.Id, &place.Title, &place.Address, &place.Description); err != nil {
			return nil, errors.New("[ GetPlacesById ] could not parse query")
		}
		places = append(places, place)
	}

	return places, nil
}
