package db

import (
	"context"
	"fmt"
	"log"

	"github.com/shipherman/gophermart/ent"

	_ "github.com/lib/pq"
)

func Connect() {}

var EntClient *ent.Client

func init() {
	//Open a connection to the database
	Client, err := ent.Open("postgres", "host=localhost port=5432 dbname=postgres user=postgres password=pass sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database successfully")
	defer Client.Close()
	// AutoMigration with ENT
	if err := Client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	EntClient = Client
}
