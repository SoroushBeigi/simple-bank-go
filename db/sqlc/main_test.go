package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"testing"
)

var testQueries *Queries
var testDB *pgxpool.Pool

const (
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}

	testQueries = New(testDB)

	code := m.Run()

	testDB.Close()
	os.Exit(code)
}
