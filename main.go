package main

import (
	"context"
	"github.com/SoroushBeigi/simple-bank-go/api"
	db "github.com/SoroushBeigi/simple-bank-go/db/sqlc"
	"github.com/SoroushBeigi/simple-bank-go/util"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}

	store := db.NewStore(conn)
	server,err := api.NewServer(config,store)
	if err!=nil{
		log.Fatal("Cannot create server: ", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
