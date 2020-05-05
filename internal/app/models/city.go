package models

import (
	"context"
	"fmt"
)

// City represents city table row in database
type City struct {
	ID    uint64
	Title string
}

const sqlQueryForGetCities = "SELECT `id`, `title` FROM `city` ORDER BY `title` ASC LIMIT ? OFFSET ?"

// GetCities returns cities sorted by title
func (db *MysqlDB) GetCities(ctx context.Context, limit uint64, offset uint64) ([]*City, error) {
	rows, err := db.QueryContext(ctx, sqlQueryForGetCities, limit, offset)
	if err != nil {
		return []*City{}, fmt.Errorf("could not execute GetCities query error: %v", err)
	}
	defer rows.Close()
	cities := make([]*City, 0)
	for rows.Next() {
		city := &City{}
		if err := rows.Scan(&city.ID, &city.Title); err != nil {
			return []*City{}, fmt.Errorf("could not scan GetCities query result, error: %v", err)
		}
		cities = append(cities, city)
	}
	return cities, nil
}
