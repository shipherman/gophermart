// Code generated by ent, DO NOT EDIT.

package withdrawals

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/shipherman/gophermart/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldLTE(FieldID, id))
}

// Order applies equality check predicate on the "order" field. It's identical to OrderEQ.
func Order(v string) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldEQ(FieldOrder, v))
}

// Sum applies equality check predicate on the "sum" field. It's identical to SumEQ.
func Sum(v int) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldEQ(FieldSum, v))
}

// Timestamp applies equality check predicate on the "timestamp" field. It's identical to TimestampEQ.
func Timestamp(v time.Time) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldEQ(FieldTimestamp, v))
}

// OrderEQ applies the EQ predicate on the "order" field.
func OrderEQ(v string) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldEQ(FieldOrder, v))
}

// OrderNEQ applies the NEQ predicate on the "order" field.
func OrderNEQ(v string) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldNEQ(FieldOrder, v))
}

// OrderIn applies the In predicate on the "order" field.
func OrderIn(vs ...string) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldIn(FieldOrder, vs...))
}

// OrderNotIn applies the NotIn predicate on the "order" field.
func OrderNotIn(vs ...string) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldNotIn(FieldOrder, vs...))
}

// OrderGT applies the GT predicate on the "order" field.
func OrderGT(v string) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldGT(FieldOrder, v))
}

// OrderGTE applies the GTE predicate on the "order" field.
func OrderGTE(v string) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldGTE(FieldOrder, v))
}

// OrderLT applies the LT predicate on the "order" field.
func OrderLT(v string) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldLT(FieldOrder, v))
}

// OrderLTE applies the LTE predicate on the "order" field.
func OrderLTE(v string) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldLTE(FieldOrder, v))
}

// OrderContains applies the Contains predicate on the "order" field.
func OrderContains(v string) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldContains(FieldOrder, v))
}

// OrderHasPrefix applies the HasPrefix predicate on the "order" field.
func OrderHasPrefix(v string) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldHasPrefix(FieldOrder, v))
}

// OrderHasSuffix applies the HasSuffix predicate on the "order" field.
func OrderHasSuffix(v string) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldHasSuffix(FieldOrder, v))
}

// OrderEqualFold applies the EqualFold predicate on the "order" field.
func OrderEqualFold(v string) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldEqualFold(FieldOrder, v))
}

// OrderContainsFold applies the ContainsFold predicate on the "order" field.
func OrderContainsFold(v string) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldContainsFold(FieldOrder, v))
}

// SumEQ applies the EQ predicate on the "sum" field.
func SumEQ(v int) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldEQ(FieldSum, v))
}

// SumNEQ applies the NEQ predicate on the "sum" field.
func SumNEQ(v int) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldNEQ(FieldSum, v))
}

// SumIn applies the In predicate on the "sum" field.
func SumIn(vs ...int) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldIn(FieldSum, vs...))
}

// SumNotIn applies the NotIn predicate on the "sum" field.
func SumNotIn(vs ...int) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldNotIn(FieldSum, vs...))
}

// SumGT applies the GT predicate on the "sum" field.
func SumGT(v int) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldGT(FieldSum, v))
}

// SumGTE applies the GTE predicate on the "sum" field.
func SumGTE(v int) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldGTE(FieldSum, v))
}

// SumLT applies the LT predicate on the "sum" field.
func SumLT(v int) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldLT(FieldSum, v))
}

// SumLTE applies the LTE predicate on the "sum" field.
func SumLTE(v int) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldLTE(FieldSum, v))
}

// TimestampEQ applies the EQ predicate on the "timestamp" field.
func TimestampEQ(v time.Time) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldEQ(FieldTimestamp, v))
}

// TimestampNEQ applies the NEQ predicate on the "timestamp" field.
func TimestampNEQ(v time.Time) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldNEQ(FieldTimestamp, v))
}

// TimestampIn applies the In predicate on the "timestamp" field.
func TimestampIn(vs ...time.Time) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldIn(FieldTimestamp, vs...))
}

// TimestampNotIn applies the NotIn predicate on the "timestamp" field.
func TimestampNotIn(vs ...time.Time) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldNotIn(FieldTimestamp, vs...))
}

// TimestampGT applies the GT predicate on the "timestamp" field.
func TimestampGT(v time.Time) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldGT(FieldTimestamp, v))
}

// TimestampGTE applies the GTE predicate on the "timestamp" field.
func TimestampGTE(v time.Time) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldGTE(FieldTimestamp, v))
}

// TimestampLT applies the LT predicate on the "timestamp" field.
func TimestampLT(v time.Time) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldLT(FieldTimestamp, v))
}

// TimestampLTE applies the LTE predicate on the "timestamp" field.
func TimestampLTE(v time.Time) predicate.Withdrawals {
	return predicate.Withdrawals(sql.FieldLTE(FieldTimestamp, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Withdrawals {
	return predicate.Withdrawals(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Withdrawals {
	return predicate.Withdrawals(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Withdrawals) predicate.Withdrawals {
	return predicate.Withdrawals(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Withdrawals) predicate.Withdrawals {
	return predicate.Withdrawals(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Withdrawals) predicate.Withdrawals {
	return predicate.Withdrawals(func(s *sql.Selector) {
		p(s.Not())
	})
}
