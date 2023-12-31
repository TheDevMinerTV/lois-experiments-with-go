// Code generated by ent, DO NOT EDIT.

package todo

import (
	"test/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Todo {
	return predicate.Todo(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Todo {
	return predicate.Todo(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Todo {
	return predicate.Todo(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Todo {
	return predicate.Todo(sql.FieldLTE(FieldID, id))
}

// Item applies equality check predicate on the "item" field. It's identical to ItemEQ.
func Item(v string) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldItem, v))
}

// Done applies equality check predicate on the "done" field. It's identical to DoneEQ.
func Done(v bool) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldDone, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldUpdatedAt, v))
}

// ItemEQ applies the EQ predicate on the "item" field.
func ItemEQ(v string) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldItem, v))
}

// ItemNEQ applies the NEQ predicate on the "item" field.
func ItemNEQ(v string) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldItem, v))
}

// ItemIn applies the In predicate on the "item" field.
func ItemIn(vs ...string) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldItem, vs...))
}

// ItemNotIn applies the NotIn predicate on the "item" field.
func ItemNotIn(vs ...string) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldItem, vs...))
}

// ItemGT applies the GT predicate on the "item" field.
func ItemGT(v string) predicate.Todo {
	return predicate.Todo(sql.FieldGT(FieldItem, v))
}

// ItemGTE applies the GTE predicate on the "item" field.
func ItemGTE(v string) predicate.Todo {
	return predicate.Todo(sql.FieldGTE(FieldItem, v))
}

// ItemLT applies the LT predicate on the "item" field.
func ItemLT(v string) predicate.Todo {
	return predicate.Todo(sql.FieldLT(FieldItem, v))
}

// ItemLTE applies the LTE predicate on the "item" field.
func ItemLTE(v string) predicate.Todo {
	return predicate.Todo(sql.FieldLTE(FieldItem, v))
}

// ItemContains applies the Contains predicate on the "item" field.
func ItemContains(v string) predicate.Todo {
	return predicate.Todo(sql.FieldContains(FieldItem, v))
}

// ItemHasPrefix applies the HasPrefix predicate on the "item" field.
func ItemHasPrefix(v string) predicate.Todo {
	return predicate.Todo(sql.FieldHasPrefix(FieldItem, v))
}

// ItemHasSuffix applies the HasSuffix predicate on the "item" field.
func ItemHasSuffix(v string) predicate.Todo {
	return predicate.Todo(sql.FieldHasSuffix(FieldItem, v))
}

// ItemEqualFold applies the EqualFold predicate on the "item" field.
func ItemEqualFold(v string) predicate.Todo {
	return predicate.Todo(sql.FieldEqualFold(FieldItem, v))
}

// ItemContainsFold applies the ContainsFold predicate on the "item" field.
func ItemContainsFold(v string) predicate.Todo {
	return predicate.Todo(sql.FieldContainsFold(FieldItem, v))
}

// DoneEQ applies the EQ predicate on the "done" field.
func DoneEQ(v bool) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldDone, v))
}

// DoneNEQ applies the NEQ predicate on the "done" field.
func DoneNEQ(v bool) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldDone, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Todo) predicate.Todo {
	return predicate.Todo(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Todo) predicate.Todo {
	return predicate.Todo(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Todo) predicate.Todo {
	return predicate.Todo(sql.NotPredicates(p))
}
