package models

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql" // imported for mysql side effects
)

// MysqlDB provides MySQL implementation for datastore
type MysqlDB struct {
	*sql.DB
}

// NewMysqlDB build MysqlDB struct and initialize it
func NewMysqlDB(url string) (*MysqlDB, error) {
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, errors.New("[ NewMysqlDB ] could not connect to MySQL server error: " + err.Error())
	}
	if err := db.Ping(); err != nil {
		return nil, errors.New("[ NewMysqlDB ] could not connect to MySQL server error: " + err.Error())
	}
	return &MysqlDB{db}, nil
}
