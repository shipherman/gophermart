// Code generated by ent, DO NOT EDIT.

package order

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/shipherman/gophermart/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldID, id))
}

// Ordernum applies equality check predicate on the "ordernum" field. It's identical to OrdernumEQ.
func Ordernum(v int) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldOrdernum, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldStatus, v))
}

// OrdernumEQ applies the EQ predicate on the "ordernum" field.
func OrdernumEQ(v int) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldOrdernum, v))
}

// OrdernumNEQ applies the NEQ predicate on the "ordernum" field.
func OrdernumNEQ(v int) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldOrdernum, v))
}

// OrdernumIn applies the In predicate on the "ordernum" field.
func OrdernumIn(vs ...int) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldOrdernum, vs...))
}

// OrdernumNotIn applies the NotIn predicate on the "ordernum" field.
func OrdernumNotIn(vs ...int) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldOrdernum, vs...))
}

// OrdernumGT applies the GT predicate on the "ordernum" field.
func OrdernumGT(v int) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldOrdernum, v))
}

// OrdernumGTE applies the GTE predicate on the "ordernum" field.
func OrdernumGTE(v int) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldOrdernum, v))
}

// OrdernumLT applies the LT predicate on the "ordernum" field.
func OrdernumLT(v int) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldOrdernum, v))
}

// OrdernumLTE applies the LTE predicate on the "ordernum" field.
func OrdernumLTE(v int) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldOrdernum, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Order {
	return predicate.Order(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Order {
	return predicate.Order(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Order {
	return predicate.Order(sql.FieldContainsFold(FieldStatus, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Order) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Order) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
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
func Not(p predicate.Order) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		p(s.Not())
	})
}
