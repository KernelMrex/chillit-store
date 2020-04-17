package placesstore

import (
	"chillit-store/internal/app/models"
	"chillit-store/internal/app/places"
	"context"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestServer_AddPlace(t *testing.T) {
	// Creating mock connection to DB
	conn, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer conn.Close()
	datastore := models.NewMockDB(conn)
	mock.ExpectExec("INSERT INTO place").
		WithArgs("title", "address", "description").
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Creating server and executing query
	server := newServer(datastore)
	resp, err := server.AddPlace(context.Background(), &places.AddPlaceRequest{
		Place: &places.Place{
			Title:       "title",
			Address:     "address",
			Description: "description",
		},
	})

	// Checking results of execution
	if err != nil || resp.Id != 1 {
		t.Errorf("error was not expected while inserting place: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestServer_AddPlace_Error(t *testing.T) {
	// Creating mock connection to DB
	conn, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer conn.Close()
	datastore := models.NewMockDB(conn)
	mock.ExpectExec("INSERT INTO place").
		WithArgs("", "", "").
		WillReturnError(fmt.Errorf("insertation in not null sections"))

	// Creating server and executing query
	server := newServer(datastore)
	_, err = server.AddPlace(context.Background(), &places.AddPlaceRequest{
		Place: &places.Place{},
	})

	// Checking results of execution
	if err == nil {
		t.Errorf("error was expected while inserting place: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
