package db

import (
	"context"
	"github.com/SoroushBeigi/simple-bank-go/util"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"testing"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	config, configErr := util.LoadConfig("../..")
	if configErr != nil {
		log.Fatal("cannot laod config: ", configErr)
	}
	var err error
	testDB, err = pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}

	testQueries = New(testDB)

	code := m.Run()

	testDB.Close()
	os.Exit(code)
}
