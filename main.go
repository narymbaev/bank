package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/narymbaev/techschool/api"
	db "github.com/narymbaev/techschool/db/sqlc"
	"log"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:root@localhost:5432/bank?sslmode=disable"
	serverAddress = "localhost:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	err = conn.Ping()
	if err != nil {
		log.Fatal("cannot ping database:", err)
	}

	fmt.Println("Connected to the database successfully")

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
