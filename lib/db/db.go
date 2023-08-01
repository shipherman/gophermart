package db

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/shipherman/gophermart/generated/ent"
	"github.com/shipherman/gophermart/lib/models"
)

type DBClient struct {
	Client     *ent.Client
	ConnString string
}

type DBClientInt interface {
	Start() error
	Stop() error
	InsertOrder(models.OrderResponse) error
	UpdateOrder(models.OrderResponse) error
	SelectOrderOwner(string) (string, error)
	SelectOrders(string) ([]models.OrderResponse, error)
	InsertUser(ent.User) error
	SelectUserExistence(string, string) (bool, error)
	SelectUser(string) (*ent.User, error)
	SelectBalance(string) (response models.BalanceResponse, err error)
	UpdateBalance(models.OrderResponse) error
	InsertWithdraw(string, models.WithdrawResponse) error
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
