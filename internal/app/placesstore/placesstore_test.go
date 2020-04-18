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

func TestServer_GetPlaces(t *testing.T) {
	// Creating mock connection to DB
	conn, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer conn.Close()
	datastore := models.NewMockDB(conn)
	mock.ExpectQuery("SELECT").
		WithArgs(20, 0).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "address", "description"}).
				AddRow(1, "title 1", "address 1", "descr 1").
				AddRow(2, "title 2", "address 2", "descr 2").
				AddRow(3, "title 3", "address 3", "descr 3"),
		)

	// Creating server and executing query
	server := newServer(datastore)
	resp, err := server.GetPlaces(context.Background(), &places.GetPlacesRequest{
		Offset: 0,
		Amount: 20,
	})

	// Check if error
	if err != nil {
		t.Errorf("error was not expected while getting places: %s", err)
	}

	// Checking results of execution
	for i, place := range resp.Places {
		j := i + 1
		if place.Id != uint64(j) ||
			place.Title != fmt.Sprintf("title %d", j) ||
			place.Address != fmt.Sprintf("address %d", j) ||
			place.Description != fmt.Sprintf("descr %d", j) {
			t.Errorf("expected values does not match")
			return
		}
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestServer_GetPlaces_BadRange(t *testing.T) {
	// Creating mock connection to DB
	conn, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer conn.Close()
	datastore := models.NewMockDB(conn)

	// Creating server and executing query
	server := newServer(datastore)
	_, err = server.GetPlaces(context.Background(), &places.GetPlacesRequest{
		Offset: 0,
		Amount: 22,
	})

	// Check if error
	if err == nil {
		t.Errorf("error was expected while getting places: %s", err)
	}
}
