package main

import (
	"context"
	"log"

	"github.com/deepaksing/Travegram/server"
	"github.com/deepaksing/Travegram/store"
	"github.com/deepaksing/Travegram/store/db/postgres"
	_ "github.com/lib/pq"
)

func main() {

	run()
}

func run() {
	ctx := context.Background()

	// 1. database connection (postgres)
	dbConn, err := postgres.NewDB()
	if err != nil {
		log.Fatal(err)
		return
	}
	dbConn.Migrate(ctx)

	store := store.NewStore(dbConn)

	// //2. CRUD API's
	server := server.NewServer(store)
	server.StartServer()
}
