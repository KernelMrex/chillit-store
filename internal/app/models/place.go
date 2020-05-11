package models

import (
	"context"
	"fmt"
)

// Place represents place table row in database
type Place struct {
	ID          uint64
	Title       string
	Address     string
	Description string
}

const (
	sqlQueryForGetRandomPlaceByCityName = "SELECT p.`id`, p.`title`, p.`address`, p.`description` FROM `place` as p, `city` as c WHERE c.`id` = p.`city_id` AND c.`title` = ? ORDER BY RAND() LIMIT 1"
	sqlQueryForSavePlace                = "INSERT INTO place (`title`, `address`, `description`, `city_id`) SELECT '?', '?', '?', `id` FROM `city` AS c WHERE `title`='?'"
)

// GetRandomPlaceByCityName returns random place from datastore
func (db *MysqlDB) GetRandomPlaceByCityName(ctx context.Context, cityName string) (*Place, error) {
	place := &Place{}
	if err := db.QueryRowContext(ctx, sqlQueryForGetRandomPlaceByCityName, cityName).Scan(
		&place.ID,
		&place.Title,
		&place.Address,
		&place.Description,
	); err != nil {
		return place, fmt.Errorf("could not get random place by city name: %v", err)
	}
	return place, nil
}

// SavePlace stores place info in database return last inserted id
func (db *MysqlDB) SavePlace(ctx context.Context, place *Place, cityName string) (uint64, error) {
	res, err := db.ExecContext(ctx, sqlQueryForSavePlace,
		place.Title,
		place.Address,
		place.Description,
		cityName,
	)
	if err != nil {
		return 0, fmt.Errorf("could not save place, error: %v", err)
	}

	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("could not get inserted id, error: %v", err)
	}

	return uint64(lastInsertedID), nil
}
