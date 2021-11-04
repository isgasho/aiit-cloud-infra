package main

import (
	"database/sql"
	"log"

	"github.com/mi-bear/infra-control/api"
	"github.com/mi-bear/infra-control/config"
	"github.com/mi-bear/infra-control/infrastructure/db"
	"github.com/mi-bear/infra-control/infrastructure/server"
)

func main() {
	conn, err := db.NewDB()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	defer func(conn *sql.DB) {
		if err := conn.Close(); err != nil {
			log.Fatalf("Failed to close the database connection: %v", err)
		}
	}(conn)

	srv := server.NewServer(
		api.BuildRouter(conn),
	)

	log.Printf("Serving on localhost:%v\n", config.Config.ServerPort)
	log.Fatal(srv.ListenAndServe())
}
