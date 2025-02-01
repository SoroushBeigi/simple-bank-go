package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"testing"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := pgx.Connect(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}
	testQueries = New(conn)

	m.Run()
}
