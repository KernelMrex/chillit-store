package placesstore

import (
	"chillit-store/internal/app/models"
	"chillit-store/internal/app/places"
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestServer_GetRandomPlaceByCityName_Success(t *testing.T) {
	// Creating mock datastore
	mockDBConn, mockAPI, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%v' happend when opening a stub database connection", err)
	}
	defer mockDBConn.Close()
	mockDatastore := models.NewMockDatastore(mockDBConn)

	mockAPI.ExpectQuery("SELECT").
		WithArgs("Йошкар-Ола").
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "address", "description"}).
				AddRow(1, "Title of cafe", "Mock address", "Description of a big amount of words"),
		)

	// Creating mock server
	server := newServer(mockDatastore)
	resp, err := server.GetRandomPlaceByCityName(context.Background(), &places.GetRandomPlaceByCityNameRequest{
		CityName: "Йошкар-Ола",
	})
	if err != nil {
		t.Fatalf("an error '%v' was not expected", err)
	}

	// Asserting returned values
	assert.Equal(t, uint64(1), resp.GetPlace().GetId())
	assert.Equal(t, "Title of cafe", resp.GetPlace().GetTitle())
	assert.Equal(t, "Mock address", resp.GetPlace().GetAddress())
	assert.Equal(t, "Description of a big amount of words", resp.GetPlace().GetDescription())

	// Checking expectations
	if err := mockAPI.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: '%v'", err)
	}
}

func TestServer_GetRandomPlaceByCityName_EmptyResponse(t *testing.T) {
	// Creating mock datastore
	mockDBConn, mockAPI, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%v' happend when opening a stub database connection", err)
	}
	defer mockDBConn.Close()
	mockDatastore := models.NewMockDatastore(mockDBConn)

	mockAPI.ExpectQuery("SELECT").
		WithArgs("Йошкар-Ола").
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "address", "description"}),
		)

	// Creating mock server
	server := newServer(mockDatastore)
	if _, err := server.GetRandomPlaceByCityName(context.Background(), &places.GetRandomPlaceByCityNameRequest{
		CityName: "Йошкар-Ола",
	}); err == nil {
		t.Fatalf("an error 'empty result' was expected")
	}

	// Checking expectations
	if err := mockAPI.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: '%v'", err)
	}
}

func TestServer_GetRandomPlaceByCityName_Timeout(t *testing.T) {
	// Creating mock datastore
	mockDBConn, mockAPI, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%v' happend when opening a stub database connection", err)
	}
	defer mockDBConn.Close()
	mockDatastore := models.NewMockDatastore(mockDBConn)

	mockAPI.ExpectQuery("SELECT").
		WithArgs("Йошкар-Ола").WillDelayFor(time.Second * 2).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "address", "description"}),
		)

	// Creating mock server
	server := newServer(mockDatastore)
	if _, err := server.GetRandomPlaceByCityName(context.Background(), &places.GetRandomPlaceByCityNameRequest{
		CityName: "Йошкар-Ола",
	}); err == nil {
		t.Fatalf("an error 'timeout' was expected")
	}

	// Checking expectations
	if err := mockAPI.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: '%v'", err)
	}
}

func TestServer_GetCities_Success(t *testing.T) {
	// Creating mock datastore
	mockDBConn, mockAPI, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%v' happend when opening a stub database connection", err)
	}
	defer mockDBConn.Close()
	mockDatastore := models.NewMockDatastore(mockDBConn)

	mockAPI.ExpectQuery("SELECT").
		WithArgs(3, 0).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title"}).
				AddRow(1, "Йошкар-Ола").
				AddRow(2, "Казань").
				AddRow(3, "Нижний Новгород"),
		)

	// Creating mock server
	server := newServer(mockDatastore)
	resp, err := server.GetCities(context.Background(), &places.GetCitiesRequest{
		Amount: 3,
		Offset: 0,
	})
	if err != nil {
		t.Fatalf("an error '%v' was not expected", err)
	}

	// Asserting returned values
	assert.Equal(t, 3, len(resp.Cities), "should return 3 cities")

	assert.Equal(t, uint64(1), resp.GetCities()[0].GetId())
	assert.Equal(t, "Йошкар-Ола", resp.GetCities()[0].GetTitle())

	assert.Equal(t, uint64(2), resp.GetCities()[1].GetId())
	assert.Equal(t, "Казань", resp.GetCities()[1].GetTitle())

	assert.Equal(t, uint64(3), resp.GetCities()[2].GetId())
	assert.Equal(t, "Нижний Новгород", resp.GetCities()[2].GetTitle())

	// Checking expectations
	if err := mockAPI.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: '%v'", err)
	}
}

func TestServer_GetCities_EmptyResponse(t *testing.T) {
	// Creating mock datastore
	mockDBConn, mockAPI, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%v' happend when opening a stub database connection", err)
	}
	defer mockDBConn.Close()
	mockDatastore := models.NewMockDatastore(mockDBConn)

	mockAPI.ExpectQuery("SELECT").
		WithArgs(3, 0).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title"}))

	// Creating mock server
	server := newServer(mockDatastore)
	if _, err := server.GetCities(context.Background(), &places.GetCitiesRequest{
		Amount: 3,
		Offset: 0,
	}); err != nil {
		t.Fatalf("an error '%v' was not expected", err)
	}

	// Checking expectations
	if err := mockAPI.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: '%v'", err)
	}
}

func TestServer_GetCities_Timeout(t *testing.T) {
	// Creating mock datastore
	mockDBConn, mockAPI, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%v' happend when opening a stub database connection", err)
	}
	defer mockDBConn.Close()
	mockDatastore := models.NewMockDatastore(mockDBConn)

	mockAPI.ExpectQuery("SELECT").
		WithArgs(3, 0).
		WillDelayFor(time.Second * 2).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title"}).
				AddRow(1, "Йошкар-Ола"),
		)

	// Creating mock server
	server := newServer(mockDatastore)
	if _, err := server.GetCities(context.Background(), &places.GetCitiesRequest{
		Amount: 3,
		Offset: 0,
	}); err == nil {
		t.Fatalf("an error 'timeout' was expected, but not given")
	}

	// Checking expectations
	if err := mockAPI.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: '%v'", err)
	}
}
