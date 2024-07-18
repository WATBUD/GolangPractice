// Code generated by ent, DO NOT EDIT.

package basenavstate

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"mai.today/core/entity/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldEQ(FieldCreatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldEQ(FieldDeletedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldEQ(FieldUpdatedAt, v))
}

// BaseID applies equality check predicate on the "base_id" field. It's identical to BaseIDEQ.
func BaseID(v string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldEQ(FieldBaseID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldEQ(FieldUserID, v))
}

// Index applies equality check predicate on the "index" field. It's identical to IndexEQ.
func Index(v int) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldEQ(FieldIndex, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldLTE(FieldCreatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldLTE(FieldDeletedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldLTE(FieldUpdatedAt, v))
}

// BaseIDEQ applies the EQ predicate on the "base_id" field.
func BaseIDEQ(v string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldEQ(FieldBaseID, v))
}

// BaseIDNEQ applies the NEQ predicate on the "base_id" field.
func BaseIDNEQ(v string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldNEQ(FieldBaseID, v))
}

// BaseIDIn applies the In predicate on the "base_id" field.
func BaseIDIn(vs ...string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldIn(FieldBaseID, vs...))
}

// BaseIDNotIn applies the NotIn predicate on the "base_id" field.
func BaseIDNotIn(vs ...string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldNotIn(FieldBaseID, vs...))
}

// BaseIDGT applies the GT predicate on the "base_id" field.
func BaseIDGT(v string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldGT(FieldBaseID, v))
}

// BaseIDGTE applies the GTE predicate on the "base_id" field.
func BaseIDGTE(v string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldGTE(FieldBaseID, v))
}

// BaseIDLT applies the LT predicate on the "base_id" field.
func BaseIDLT(v string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldLT(FieldBaseID, v))
}

// BaseIDLTE applies the LTE predicate on the "base_id" field.
func BaseIDLTE(v string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldLTE(FieldBaseID, v))
}

// BaseIDContains applies the Contains predicate on the "base_id" field.
func BaseIDContains(v string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldContains(FieldBaseID, v))
}

// BaseIDHasPrefix applies the HasPrefix predicate on the "base_id" field.
func BaseIDHasPrefix(v string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldHasPrefix(FieldBaseID, v))
}

// BaseIDHasSuffix applies the HasSuffix predicate on the "base_id" field.
func BaseIDHasSuffix(v string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldHasSuffix(FieldBaseID, v))
}

// BaseIDEqualFold applies the EqualFold predicate on the "base_id" field.
func BaseIDEqualFold(v string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldEqualFold(FieldBaseID, v))
}

// BaseIDContainsFold applies the ContainsFold predicate on the "base_id" field.
func BaseIDContainsFold(v string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldContainsFold(FieldBaseID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldContainsFold(FieldUserID, v))
}

// IndexEQ applies the EQ predicate on the "index" field.
func IndexEQ(v int) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldEQ(FieldIndex, v))
}

// IndexNEQ applies the NEQ predicate on the "index" field.
func IndexNEQ(v int) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldNEQ(FieldIndex, v))
}

// IndexIn applies the In predicate on the "index" field.
func IndexIn(vs ...int) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldIn(FieldIndex, vs...))
}

// IndexNotIn applies the NotIn predicate on the "index" field.
func IndexNotIn(vs ...int) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldNotIn(FieldIndex, vs...))
}

// IndexGT applies the GT predicate on the "index" field.
func IndexGT(v int) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldGT(FieldIndex, v))
}

// IndexGTE applies the GTE predicate on the "index" field.
func IndexGTE(v int) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldGTE(FieldIndex, v))
}

// IndexLT applies the LT predicate on the "index" field.
func IndexLT(v int) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldLT(FieldIndex, v))
}

// IndexLTE applies the LTE predicate on the "index" field.
func IndexLTE(v int) predicate.BaseNavState {
	return predicate.BaseNavState(sql.FieldLTE(FieldIndex, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BaseNavState) predicate.BaseNavState {
	return predicate.BaseNavState(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BaseNavState) predicate.BaseNavState {
	return predicate.BaseNavState(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.BaseNavState) predicate.BaseNavState {
	return predicate.BaseNavState(sql.NotPredicates(p))
}
