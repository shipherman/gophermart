package db

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/shipherman/gophermart/ent"
)

type DBClient struct {
	Client     *ent.Client
	ConnString string
}

// Create Client instance
func NewClient(connString string) *DBClient {
	return &DBClient{ConnString: connString}
}

// Connect to DB
func (dbc *DBClient) Start() error {
	var err error
	dbc.Client, err = ent.Open("postgres", dbc.ConnString)
	if err != nil {
		return err
	}

	fmt.Println("Connected to database successfully")

	// AutoMigration with ENT
	if err := dbc.Client.Schema.Create(context.Background()); err != nil {
		return fmt.Errorf("failed creating schema resources: %v", err)
	}
	return nil
}

// Close connection to DB
func (dbc *DBClient) Stop() error {
	return dbc.Client.Close()
}
