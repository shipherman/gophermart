// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/shipherman/gophermart/ent/schema"
	"github.com/shipherman/gophermart/ent/user"
	"github.com/shipherman/gophermart/ent/withdrawals"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescLogin is the schema descriptor for login field.
	userDescLogin := userFields[0].Descriptor()
	// user.LoginValidator is a validator for the "login" field. It is called by the builders before save.
	user.LoginValidator = userDescLogin.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[1].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	withdrawalsFields := schema.Withdrawals{}.Fields()
	_ = withdrawalsFields
	// withdrawalsDescOrder is the schema descriptor for order field.
	withdrawalsDescOrder := withdrawalsFields[0].Descriptor()
	// withdrawals.OrderValidator is a validator for the "order" field. It is called by the builders before save.
	withdrawals.OrderValidator = withdrawalsDescOrder.Validators[0].(func(string) error)
	// withdrawalsDescSum is the schema descriptor for sum field.
	withdrawalsDescSum := withdrawalsFields[1].Descriptor()
	// withdrawals.SumValidator is a validator for the "sum" field. It is called by the builders before save.
	withdrawals.SumValidator = withdrawalsDescSum.Validators[0].(func(string) error)
}
