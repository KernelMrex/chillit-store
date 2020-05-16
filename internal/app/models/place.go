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
	sqlQueryForGetPlacesByCityID        = "SELECT `id`, `title`, `address`, `description` FROM `place` WHERE `city_id`=? LIMIT ? OFFSET ?"
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

// GetPlacesByCityID ...
func (db *MysqlDB) GetPlacesByCityID(ctx context.Context, cityID uint64, amount uint64, offset uint64) ([]*Place, error) {
	resp, err := db.QueryContext(ctx, sqlQueryForGetPlacesByCityID, cityID, amount, offset)
	if err != nil {
		return nil, fmt.Errorf("could not get places by city id '%d': %v", cityID, err)
	}
	defer resp.Close()

	places := make([]*Place, 0)
	for resp.Next() {
		place := &Place{}
		if err := resp.Scan(&place.ID, &place.Title, &place.Address, &place.Description); err != nil {
			return nil, fmt.Errorf("could not scan places: %v", err)
		}
		places = append(places, place)
	}

	return places, nil
}
