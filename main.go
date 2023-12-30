package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/narymbaev/techschool/api"
	db "github.com/narymbaev/techschool/db/sqlc"
	"github.com/narymbaev/techschool/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load config:", err )
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
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

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
