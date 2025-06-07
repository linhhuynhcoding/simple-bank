package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/linhhuynhcoding/learn-go/api"
	"github.com/linhhuynhcoding/learn-go/db"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank_go?sslmode=disable"
	serverAddress = "0.0.0.0:8088"
)

func main() {
	fmt.Println("Hello, world!")
	// var x int = "string" // This should trigger an error: cannot use "string" (type string) as int
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
