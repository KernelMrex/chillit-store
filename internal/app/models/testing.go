package models

import "database/sql"

// NewMockDatastore provides interface for testing
func NewMockDatastore(mockDBConn *sql.DB) Datastore {
	return &MysqlDB{mockDBConn}
}
