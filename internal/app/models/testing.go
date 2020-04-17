package models

import "database/sql"

// MockDB structure for testing
type MockDB struct {
	*MysqlDB
}

// NewMockDB provides configured structure for testing
func NewMockDB(conn *sql.DB) *MockDB {
	return &MockDB{
		&MysqlDB{conn},
	}
}
