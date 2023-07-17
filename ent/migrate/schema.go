// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// OrdersColumns holds the columns for the "orders" table.
	OrdersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "ordernum", Type: field.TypeInt, Unique: true},
		{Name: "accural", Type: field.TypeInt},
		{Name: "status", Type: field.TypeString},
		{Name: "timestamp", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp with time zone"}},
		{Name: "user_orders", Type: field.TypeInt, Nullable: true},
	}
	// OrdersTable holds the schema information for the "orders" table.
	OrdersTable = &schema.Table{
		Name:       "orders",
		Columns:    OrdersColumns,
		PrimaryKey: []*schema.Column{OrdersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "orders_users_orders",
				Columns:    []*schema.Column{OrdersColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "login", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "balance", Type: field.TypeInt},
		{Name: "withdraw", Type: field.TypeInt},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// WithdrawalsColumns holds the columns for the "withdrawals" table.
	WithdrawalsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "order", Type: field.TypeString, Unique: true},
		{Name: "sum", Type: field.TypeString},
	}
	// WithdrawalsTable holds the schema information for the "withdrawals" table.
	WithdrawalsTable = &schema.Table{
		Name:       "withdrawals",
		Columns:    WithdrawalsColumns,
		PrimaryKey: []*schema.Column{WithdrawalsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		OrdersTable,
		UsersTable,
		WithdrawalsTable,
	}
)

func init() {
	OrdersTable.ForeignKeys[0].RefTable = UsersTable
}
