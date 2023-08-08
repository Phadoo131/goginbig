package db

import (
	"testing"
	"log"
	"os"
	"database/sql"

	_"github.com/lib/pq"
)
var testQuerie *Queries
var testDB *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgresql://bigsimpleapi:phadoo131@localhost:5444/BookStore?sslmode=disable"
)

func TestMain(m *testing.M){
	var err error

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to the DB: ", err)
	}

	testQuerie = New(testDB)

	os.Exit(m.Run())
}