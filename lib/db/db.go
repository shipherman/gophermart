package db

import (
	"context"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/shipherman/gophermart/ent"
)

var entClient *ent.Client

func NewClient(connString string) *ent.Client {
	//Open a connection to the database
	entClient, err := ent.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	fmt.Println("Connected to database successfully")

	// defer EntClient.Close()
	// AutoMigration with ENT
	if err := entClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		return nil
	}
	return entClient
}

func GetClient() *ent.Client {
	return entClient
}

func SetClient(newClient *ent.Client) {
	entClient = newClient
}
