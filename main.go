package main

import (
	"database/sql"
	"log"

	"github.com/linhhuynhcoding/learn-go/api"
	"github.com/linhhuynhcoding/learn-go/db"
	"github.com/linhhuynhcoding/learn-go/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBUrl)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
