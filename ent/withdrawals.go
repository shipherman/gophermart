// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/shipherman/gophermart/ent/user"
	"github.com/shipherman/gophermart/ent/withdrawals"
)

// Withdrawals is the model entity for the Withdrawals schema.
type Withdrawals struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Order holds the value of the "order" field.
	Order string `json:"order,omitempty"`
	// Sum holds the value of the "sum" field.
	Sum float64 `json:"sum,omitempty"`
	// Timestamp holds the value of the "timestamp" field.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WithdrawalsQuery when eager-loading is set.
	Edges            WithdrawalsEdges `json:"edges"`
	user_withdrawals *int
	selectValues     sql.SelectValues
}

// WithdrawalsEdges holds the relations/edges for other nodes in the graph.
type WithdrawalsEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WithdrawalsEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Withdrawals) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case withdrawals.FieldSum:
			values[i] = new(sql.NullFloat64)
		case withdrawals.FieldID:
			values[i] = new(sql.NullInt64)
		case withdrawals.FieldOrder:
			values[i] = new(sql.NullString)
		case withdrawals.FieldTimestamp:
			values[i] = new(sql.NullTime)
		case withdrawals.ForeignKeys[0]: // user_withdrawals
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Withdrawals fields.
func (w *Withdrawals) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case withdrawals.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			w.ID = int(value.Int64)
		case withdrawals.FieldOrder:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field order", values[i])
			} else if value.Valid {
				w.Order = value.String
			}
		case withdrawals.FieldSum:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field sum", values[i])
			} else if value.Valid {
				w.Sum = value.Float64
			}
		case withdrawals.FieldTimestamp:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field timestamp", values[i])
			} else if value.Valid {
				w.Timestamp = value.Time
			}
		case withdrawals.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_withdrawals", value)
			} else if value.Valid {
				w.user_withdrawals = new(int)
				*w.user_withdrawals = int(value.Int64)
			}
		default:
			w.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Withdrawals.
// This includes values selected through modifiers, order, etc.
func (w *Withdrawals) Value(name string) (ent.Value, error) {
	return w.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Withdrawals entity.
func (w *Withdrawals) QueryUser() *UserQuery {
	return NewWithdrawalsClient(w.config).QueryUser(w)
}

// Update returns a builder for updating this Withdrawals.
// Note that you need to call Withdrawals.Unwrap() before calling this method if this Withdrawals
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Withdrawals) Update() *WithdrawalsUpdateOne {
	return NewWithdrawalsClient(w.config).UpdateOne(w)
}

// Unwrap unwraps the Withdrawals entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Withdrawals) Unwrap() *Withdrawals {
	_tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("ent: Withdrawals is not a transactional entity")
	}
	w.config.driver = _tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Withdrawals) String() string {
	var builder strings.Builder
	builder.WriteString("Withdrawals(")
	builder.WriteString(fmt.Sprintf("id=%v, ", w.ID))
	builder.WriteString("order=")
	builder.WriteString(w.Order)
	builder.WriteString(", ")
	builder.WriteString("sum=")
	builder.WriteString(fmt.Sprintf("%v", w.Sum))
	builder.WriteString(", ")
	builder.WriteString("timestamp=")
	builder.WriteString(w.Timestamp.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// WithdrawalsSlice is a parsable slice of Withdrawals.
type WithdrawalsSlice []*Withdrawals
