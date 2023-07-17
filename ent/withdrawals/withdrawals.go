// Code generated by ent, DO NOT EDIT.

package withdrawals

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the withdrawals type in the database.
	Label = "withdrawals"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOrder holds the string denoting the order field in the database.
	FieldOrder = "order"
	// FieldSum holds the string denoting the sum field in the database.
	FieldSum = "sum"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the withdrawals in the database.
	Table = "withdrawals"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "withdrawals"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_withdrawals"
)

// Columns holds all SQL columns for withdrawals fields.
var Columns = []string{
	FieldID,
	FieldOrder,
	FieldSum,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "withdrawals"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_withdrawals",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Withdrawals queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByOrder orders the results by the order field.
func ByOrder(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrder, opts...).ToFunc()
}

// BySum orders the results by the sum field.
func BySum(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSum, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
