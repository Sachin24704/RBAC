package main

import (
	"backend/db"
	"backend/server"
	"log"
)

func main() {
	client := db.NewDBClient()
	defer client.Close()
	log.Println("DB setup completed")
	db.SeedDatabase(client)
	server.StartServer(client)
}
