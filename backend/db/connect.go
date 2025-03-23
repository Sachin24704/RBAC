package db

import (
	"backend/ent"
	"context"
	"log"

	_ "github.com/lib/pq"
)

func NewDBClient() *ent.Client {
	connStr := "postgresql://sachin:sachin@localhost:5432/rbac?sslmode=disable"
	client, err := ent.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to open connection to db: %v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal(err)
	}
	log.Print("connected to db")
	return client
}
