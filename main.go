package main

import (
	"context"
	"github.com/SoroushBeigi/simple-bank-go/api"
	db "github.com/SoroushBeigi/simple-bank-go/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

const (
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
