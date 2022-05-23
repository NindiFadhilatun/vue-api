package main

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"vue-api/internal/data"

	"github.com/DATA-DOG/go-sqlmock"
)

var testApp application
var mockedDB sqlmock.Sqlmock

func TestMain(m *testing.M) {
	testDB, myMock, _ := sqlmock.New()
	mockedDB = myMock

	defer func(testDB *sql.DB) {
		err := testDB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(testDB)

	testApp = application{
		config:      config{},
		infoLog:     log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog:    log.New(os.Stdout, "Error\t", log.Ldate|log.Ltime),
		models:      data.New(testDB),
		environment: "development",
	}

	os.Exit(m.Run())
}
