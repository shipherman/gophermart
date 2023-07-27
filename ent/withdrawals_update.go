// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shipherman/gophermart/ent/predicate"
	"github.com/shipherman/gophermart/ent/user"
	"github.com/shipherman/gophermart/ent/withdrawals"
)

// WithdrawalsUpdate is the builder for updating Withdrawals entities.
type WithdrawalsUpdate struct {
	config
	hooks    []Hook
	mutation *WithdrawalsMutation
}

// Where appends a list predicates to the WithdrawalsUpdate builder.
func (wu *WithdrawalsUpdate) Where(ps ...predicate.Withdrawals) *WithdrawalsUpdate {
	wu.mutation.Where(ps...)
	return wu
}

// SetOrder sets the "order" field.
func (wu *WithdrawalsUpdate) SetOrder(s string) *WithdrawalsUpdate {
	wu.mutation.SetOrder(s)
	return wu
}

// SetSum sets the "sum" field.
func (wu *WithdrawalsUpdate) SetSum(f float64) *WithdrawalsUpdate {
	wu.mutation.ResetSum()
	wu.mutation.SetSum(f)
	return wu
}

// AddSum adds f to the "sum" field.
func (wu *WithdrawalsUpdate) AddSum(f float64) *WithdrawalsUpdate {
	wu.mutation.AddSum(f)
	return wu
}

// SetTimestamp sets the "timestamp" field.
func (wu *WithdrawalsUpdate) SetTimestamp(t time.Time) *WithdrawalsUpdate {
	wu.mutation.SetTimestamp(t)
	return wu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (wu *WithdrawalsUpdate) SetUserID(id int) *WithdrawalsUpdate {
	wu.mutation.SetUserID(id)
	return wu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (wu *WithdrawalsUpdate) SetNillableUserID(id *int) *WithdrawalsUpdate {
	if id != nil {
		wu = wu.SetUserID(*id)
	}
	return wu
}

// SetUser sets the "user" edge to the User entity.
func (wu *WithdrawalsUpdate) SetUser(u *User) *WithdrawalsUpdate {
	return wu.SetUserID(u.ID)
}

// Mutation returns the WithdrawalsMutation object of the builder.
func (wu *WithdrawalsUpdate) Mutation() *WithdrawalsMutation {
	return wu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (wu *WithdrawalsUpdate) ClearUser() *WithdrawalsUpdate {
	wu.mutation.ClearUser()
	return wu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WithdrawalsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, wu.sqlSave, wu.mutation, wu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WithdrawalsUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WithdrawalsUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WithdrawalsUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wu *WithdrawalsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(withdrawals.Table, withdrawals.Columns, sqlgraph.NewFieldSpec(withdrawals.FieldID, field.TypeInt))
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.Order(); ok {
		_spec.SetField(withdrawals.FieldOrder, field.TypeString, value)
	}
	if value, ok := wu.mutation.Sum(); ok {
		_spec.SetField(withdrawals.FieldSum, field.TypeFloat64, value)
	}
	if value, ok := wu.mutation.AddedSum(); ok {
		_spec.AddField(withdrawals.FieldSum, field.TypeFloat64, value)
	}
	if value, ok := wu.mutation.Timestamp(); ok {
		_spec.SetField(withdrawals.FieldTimestamp, field.TypeTime, value)
	}
	if wu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   withdrawals.UserTable,
			Columns: []string{withdrawals.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   withdrawals.UserTable,
			Columns: []string{withdrawals.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{withdrawals.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wu.mutation.done = true
	return n, nil
}

// WithdrawalsUpdateOne is the builder for updating a single Withdrawals entity.
type WithdrawalsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WithdrawalsMutation
}

// SetOrder sets the "order" field.
func (wuo *WithdrawalsUpdateOne) SetOrder(s string) *WithdrawalsUpdateOne {
	wuo.mutation.SetOrder(s)
	return wuo
}

// SetSum sets the "sum" field.
func (wuo *WithdrawalsUpdateOne) SetSum(f float64) *WithdrawalsUpdateOne {
	wuo.mutation.ResetSum()
	wuo.mutation.SetSum(f)
	return wuo
}

// AddSum adds f to the "sum" field.
func (wuo *WithdrawalsUpdateOne) AddSum(f float64) *WithdrawalsUpdateOne {
	wuo.mutation.AddSum(f)
	return wuo
}

// SetTimestamp sets the "timestamp" field.
func (wuo *WithdrawalsUpdateOne) SetTimestamp(t time.Time) *WithdrawalsUpdateOne {
	wuo.mutation.SetTimestamp(t)
	return wuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (wuo *WithdrawalsUpdateOne) SetUserID(id int) *WithdrawalsUpdateOne {
	wuo.mutation.SetUserID(id)
	return wuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (wuo *WithdrawalsUpdateOne) SetNillableUserID(id *int) *WithdrawalsUpdateOne {
	if id != nil {
		wuo = wuo.SetUserID(*id)
	}
	return wuo
}

// SetUser sets the "user" edge to the User entity.
func (wuo *WithdrawalsUpdateOne) SetUser(u *User) *WithdrawalsUpdateOne {
	return wuo.SetUserID(u.ID)
}

// Mutation returns the WithdrawalsMutation object of the builder.
func (wuo *WithdrawalsUpdateOne) Mutation() *WithdrawalsMutation {
	return wuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (wuo *WithdrawalsUpdateOne) ClearUser() *WithdrawalsUpdateOne {
	wuo.mutation.ClearUser()
	return wuo
}

// Where appends a list predicates to the WithdrawalsUpdate builder.
func (wuo *WithdrawalsUpdateOne) Where(ps ...predicate.Withdrawals) *WithdrawalsUpdateOne {
	wuo.mutation.Where(ps...)
	return wuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WithdrawalsUpdateOne) Select(field string, fields ...string) *WithdrawalsUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Withdrawals entity.
func (wuo *WithdrawalsUpdateOne) Save(ctx context.Context) (*Withdrawals, error) {
	return withHooks(ctx, wuo.sqlSave, wuo.mutation, wuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WithdrawalsUpdateOne) SaveX(ctx context.Context) *Withdrawals {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WithdrawalsUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WithdrawalsUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wuo *WithdrawalsUpdateOne) sqlSave(ctx context.Context) (_node *Withdrawals, err error) {
	_spec := sqlgraph.NewUpdateSpec(withdrawals.Table, withdrawals.Columns, sqlgraph.NewFieldSpec(withdrawals.FieldID, field.TypeInt))
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Withdrawals.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, withdrawals.FieldID)
		for _, f := range fields {
			if !withdrawals.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != withdrawals.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.Order(); ok {
		_spec.SetField(withdrawals.FieldOrder, field.TypeString, value)
	}
	if value, ok := wuo.mutation.Sum(); ok {
		_spec.SetField(withdrawals.FieldSum, field.TypeFloat64, value)
	}
	if value, ok := wuo.mutation.AddedSum(); ok {
		_spec.AddField(withdrawals.FieldSum, field.TypeFloat64, value)
	}
	if value, ok := wuo.mutation.Timestamp(); ok {
		_spec.SetField(withdrawals.FieldTimestamp, field.TypeTime, value)
	}
	if wuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   withdrawals.UserTable,
			Columns: []string{withdrawals.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   withdrawals.UserTable,
			Columns: []string{withdrawals.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Withdrawals{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{withdrawals.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wuo.mutation.done = true
	return _node, nil
}
