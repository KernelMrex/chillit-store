package placesstore

import (
	"chillit-store/internal/app/models"
	"chillit-store/internal/app/places"
	"context"
	"fmt"
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

func TestServer_SavePlace_Success(t *testing.T) {
	// Creating mock datastore
	mockDBConn, mockAPI, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%v' happend when opening a stub database connection", err)
	}
	defer mockDBConn.Close()
	mockDatastore := models.NewMockDatastore(mockDBConn)

	mockAPI.ExpectExec("INSERT").
		WithArgs("Камелот", "Ул. Пушкина 6", "Описание описание описание", "Казань").
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Creating mock server
	server := newServer(mockDatastore)
	resp, err := server.AddPlace(context.Background(), &places.AddPlaceRequest{
		CityName: "Казань",
		Place: &places.Place{
			Title:       "Камелот",
			Address:     "Ул. Пушкина 6",
			Description: "Описание описание описание",
		},
	})
	if err != nil {
		t.Fatalf("an error '%v' was not expected", err)
	}

	assert.Equal(t, uint64(1), resp.GetId(), "last inserted id must be 1")

	// Checking expectations
	if err := mockAPI.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: '%v'", err)
	}
}

func TestServer_SavePlace_Duplicate(t *testing.T) {
	// Creating mock datastore
	mockDBConn, mockAPI, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%v' happend when opening a stub database connection", err)
	}
	defer mockDBConn.Close()
	mockDatastore := models.NewMockDatastore(mockDBConn)

	mockAPI.ExpectExec("INSERT").
		WithArgs("Камелот", "Ул. Пушкина 6", "Описание описание описание", "Казань").
		WillReturnError(fmt.Errorf("duplicate mock error"))

	// Creating mock server
	server := newServer(mockDatastore)
	if _, err := server.AddPlace(context.Background(), &places.AddPlaceRequest{
		CityName: "Казань",
		Place: &places.Place{
			Title:       "Камелот",
			Address:     "Ул. Пушкина 6",
			Description: "Описание описание описание",
		},
	}); err == nil {
		t.Fatalf("an error 'duplicate mock error' was expected, but nil given")
	}

	// Checking expectations
	if err := mockAPI.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: '%v'", err)
	}
}

func TestServer_SavePlace_Timeout(t *testing.T) {
	// Creating mock datastore
	mockDBConn, mockAPI, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%v' happend when opening a stub database connection", err)
	}
	defer mockDBConn.Close()
	mockDatastore := models.NewMockDatastore(mockDBConn)

	mockAPI.ExpectExec("INSERT").
		WithArgs("Камелот", "Ул. Пушкина 6", "Описание описание описание", "Казань").
		WillDelayFor(time.Second * 5)

	// Creating mock server
	server := newServer(mockDatastore)
	if _, err := server.AddPlace(context.Background(), &places.AddPlaceRequest{
		CityName: "Казань",
		Place: &places.Place{
			Title:       "Камелот",
			Address:     "Ул. Пушкина 6",
			Description: "Описание описание описание",
		},
	}); err == nil {
		t.Fatalf("an error 'timeout exceeded' was expected, but nil given")
	}

	// Checking expectations
	if err := mockAPI.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: '%v'", err)
	}
}

func TestServer_GetPlacesByCityID_Success(t *testing.T) {
	// Creating mock datastore
	mockDBConn, mockAPI, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%v' happend when opening a stub database connection", err)
	}
	defer mockDBConn.Close()
	mockDatastore := models.NewMockDatastore(mockDBConn)

	mockAPI.ExpectQuery("SELECT").
		WithArgs(2, 5, 0).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "address", "description"}).
				AddRow(1, "Title1", "Address1", "Description1").
				AddRow(2, "Title2", "Address2", "Description2").
				AddRow(3, "Title3", "Address3", "Description3"),
		)

	// Creating mock server
	server := newServer(mockDatastore)
	resp, err := server.GetPlacesByCityID(context.Background(), &places.GetPlacesByCityIDRequest{
		CityID: 2,
		Amount: 5,
		Offset: 0,
	})
	if err != nil {
		t.Fatalf("error was not expected: %v", err)
	}

	// Checking results of request
	assert.Equal(t, 3, len(resp.GetPlaces()), "amount of places should be 3")

	assert.Equal(t, uint64(1), resp.GetPlaces()[0].GetId())
	assert.Equal(t, "Title1", resp.GetPlaces()[0].GetTitle())
	assert.Equal(t, "Address1", resp.GetPlaces()[0].GetAddress())
	assert.Equal(t, "Description1", resp.GetPlaces()[0].GetDescription())

	assert.Equal(t, uint64(2), resp.GetPlaces()[1].GetId())
	assert.Equal(t, "Title2", resp.GetPlaces()[1].GetTitle())
	assert.Equal(t, "Address2", resp.GetPlaces()[1].GetAddress())
	assert.Equal(t, "Description2", resp.GetPlaces()[1].GetDescription())

	assert.Equal(t, uint64(3), resp.GetPlaces()[2].GetId())
	assert.Equal(t, "Title3", resp.GetPlaces()[2].GetTitle())
	assert.Equal(t, "Address3", resp.GetPlaces()[2].GetAddress())
	assert.Equal(t, "Description3", resp.GetPlaces()[2].GetDescription())

	// Checking expectations
	if err := mockAPI.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: '%v'", err)
	}
}

func TestServer_GetPlacesByCityID_Timeout(t *testing.T) {
	// Creating mock datastore
	mockDBConn, mockAPI, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%v' happend when opening a stub database connection", err)
	}
	defer mockDBConn.Close()
	mockDatastore := models.NewMockDatastore(mockDBConn)

	mockAPI.ExpectQuery("SELECT").
		WithArgs(2, 5, 0).
		WillDelayFor(time.Second * 5)

	// Creating mock server
	server := newServer(mockDatastore)
	if _, err := server.GetPlacesByCityID(context.Background(), &places.GetPlacesByCityIDRequest{
		CityID: 2,
		Amount: 5,
		Offset: 0,
	}); err == nil {
		t.Fatalf("error 'timeout' was expected, but nil given")
	}

	// Checking expectations
	if err := mockAPI.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: '%v'", err)
	}
}
