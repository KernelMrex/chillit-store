package models

import (
	"context"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestShouldAddOnePlace(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	datastore := &MysqlDB{db}

	mock.ExpectExec("INSERT INTO place").
		WithArgs("test 1", "test 2", "test 3").
		WillReturnResult(sqlmock.NewResult(1, 1))
	if lastInserted, err := datastore.AddPlace(context.Background(), &Place{
		Title:       "test 1",
		Address:     "test 2",
		Description: "test 3",
	}); err != nil || lastInserted != 1 {
		t.Errorf("error was not expected while inserting place: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestBadTitle(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	datastore := &MysqlDB{db}

	mock.ExpectExec("INSERT INTO place").
		WithArgs("", "test 2", "test 3").
		WillReturnError(fmt.Errorf("some error"))
	if lastInserted, err := datastore.AddPlace(context.Background(), &Place{
		Title:       "",
		Address:     "test 2",
		Description: "test 3",
	}); err == nil || lastInserted != 0 {
		t.Errorf("error was expected while inserting place: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: '%s'", err)
	}
}

func TestBadAddress(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	datastore := &MysqlDB{db}

	mock.ExpectExec("INSERT INTO place").
		WithArgs("test 1", "", "test 3").
		WillReturnError(fmt.Errorf("some error"))
	if lastInserted, err := datastore.AddPlace(context.Background(), &Place{
		Title:       "test 1",
		Address:     "",
		Description: "test 3",
	}); err == nil || lastInserted != 0 {
		t.Errorf("error was expected while inserting place: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: '%s'", err)
	}
}

func TestBadDescription(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	datastore := &MysqlDB{db}

	mock.ExpectExec("INSERT INTO place").
		WithArgs("test 1", "test 2", "").
		WillReturnError(fmt.Errorf("some error"))
	if lastInserted, err := datastore.AddPlace(context.Background(), &Place{
		Title:       "test 1",
		Address:     "test 2",
		Description: "",
	}); err == nil || lastInserted != 0 {
		t.Errorf("error was expected while inserting place: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: '%s'", err)
	}
}

func TestShouldReturnTwoPlaces(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	datastore := &MysqlDB{db}

	mock.ExpectQuery("SELECT").
		WithArgs(3, 1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "address", "description"}).
				AddRow(2, "t2", "a2", "d2").
				AddRow(3, "t3", "a3", "d3").
				AddRow(4, "t4", "a4", "d4"),
		)

	const offset uint64 = 1
	rows, err := datastore.GetPlacesById(context.Background(), offset, 3)
	if err != nil {
		t.Errorf("error was not expected while selecting: '%s'", err)
	}

	for i, row := range rows {
		j := uint64(i) + offset + 1
		if row.Id != j ||
			row.Title != fmt.Sprintf("t%d", j) ||
			row.Address != fmt.Sprintf("a%d", j) ||
			row.Description != fmt.Sprintf("d%d", j) {
			t.Errorf("results are not expected %v | %s", *row, fmt.Sprintf("t%d", j))
		}
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
