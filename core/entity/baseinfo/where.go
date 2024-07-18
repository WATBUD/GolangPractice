// Code generated by ent, DO NOT EDIT.

package baseinfo

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"mai.today/core/entity/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldEQ(FieldCreatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldEQ(FieldDeletedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldEQ(FieldUpdatedAt, v))
}

// BaseID applies equality check predicate on the "base_id" field. It's identical to BaseIDEQ.
func BaseID(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldEQ(FieldBaseID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldEQ(FieldName, v))
}

// Logo applies equality check predicate on the "logo" field. It's identical to LogoEQ.
func Logo(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldEQ(FieldLogo, v))
}

// Color applies equality check predicate on the "color" field. It's identical to ColorEQ.
func Color(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldEQ(FieldColor, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldLTE(FieldCreatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldLTE(FieldDeletedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldLTE(FieldUpdatedAt, v))
}

// BaseIDEQ applies the EQ predicate on the "base_id" field.
func BaseIDEQ(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldEQ(FieldBaseID, v))
}

// BaseIDNEQ applies the NEQ predicate on the "base_id" field.
func BaseIDNEQ(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldNEQ(FieldBaseID, v))
}

// BaseIDIn applies the In predicate on the "base_id" field.
func BaseIDIn(vs ...string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldIn(FieldBaseID, vs...))
}

// BaseIDNotIn applies the NotIn predicate on the "base_id" field.
func BaseIDNotIn(vs ...string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldNotIn(FieldBaseID, vs...))
}

// BaseIDGT applies the GT predicate on the "base_id" field.
func BaseIDGT(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldGT(FieldBaseID, v))
}

// BaseIDGTE applies the GTE predicate on the "base_id" field.
func BaseIDGTE(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldGTE(FieldBaseID, v))
}

// BaseIDLT applies the LT predicate on the "base_id" field.
func BaseIDLT(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldLT(FieldBaseID, v))
}

// BaseIDLTE applies the LTE predicate on the "base_id" field.
func BaseIDLTE(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldLTE(FieldBaseID, v))
}

// BaseIDContains applies the Contains predicate on the "base_id" field.
func BaseIDContains(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldContains(FieldBaseID, v))
}

// BaseIDHasPrefix applies the HasPrefix predicate on the "base_id" field.
func BaseIDHasPrefix(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldHasPrefix(FieldBaseID, v))
}

// BaseIDHasSuffix applies the HasSuffix predicate on the "base_id" field.
func BaseIDHasSuffix(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldHasSuffix(FieldBaseID, v))
}

// BaseIDEqualFold applies the EqualFold predicate on the "base_id" field.
func BaseIDEqualFold(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldEqualFold(FieldBaseID, v))
}

// BaseIDContainsFold applies the ContainsFold predicate on the "base_id" field.
func BaseIDContainsFold(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldContainsFold(FieldBaseID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldContainsFold(FieldName, v))
}

// LogoEQ applies the EQ predicate on the "logo" field.
func LogoEQ(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldEQ(FieldLogo, v))
}

// LogoNEQ applies the NEQ predicate on the "logo" field.
func LogoNEQ(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldNEQ(FieldLogo, v))
}

// LogoIn applies the In predicate on the "logo" field.
func LogoIn(vs ...string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldIn(FieldLogo, vs...))
}

// LogoNotIn applies the NotIn predicate on the "logo" field.
func LogoNotIn(vs ...string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldNotIn(FieldLogo, vs...))
}

// LogoGT applies the GT predicate on the "logo" field.
func LogoGT(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldGT(FieldLogo, v))
}

// LogoGTE applies the GTE predicate on the "logo" field.
func LogoGTE(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldGTE(FieldLogo, v))
}

// LogoLT applies the LT predicate on the "logo" field.
func LogoLT(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldLT(FieldLogo, v))
}

// LogoLTE applies the LTE predicate on the "logo" field.
func LogoLTE(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldLTE(FieldLogo, v))
}

// LogoContains applies the Contains predicate on the "logo" field.
func LogoContains(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldContains(FieldLogo, v))
}

// LogoHasPrefix applies the HasPrefix predicate on the "logo" field.
func LogoHasPrefix(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldHasPrefix(FieldLogo, v))
}

// LogoHasSuffix applies the HasSuffix predicate on the "logo" field.
func LogoHasSuffix(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldHasSuffix(FieldLogo, v))
}

// LogoEqualFold applies the EqualFold predicate on the "logo" field.
func LogoEqualFold(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldEqualFold(FieldLogo, v))
}

// LogoContainsFold applies the ContainsFold predicate on the "logo" field.
func LogoContainsFold(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldContainsFold(FieldLogo, v))
}

// ColorEQ applies the EQ predicate on the "color" field.
func ColorEQ(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldEQ(FieldColor, v))
}

// ColorNEQ applies the NEQ predicate on the "color" field.
func ColorNEQ(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldNEQ(FieldColor, v))
}

// ColorIn applies the In predicate on the "color" field.
func ColorIn(vs ...string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldIn(FieldColor, vs...))
}

// ColorNotIn applies the NotIn predicate on the "color" field.
func ColorNotIn(vs ...string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldNotIn(FieldColor, vs...))
}

// ColorGT applies the GT predicate on the "color" field.
func ColorGT(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldGT(FieldColor, v))
}

// ColorGTE applies the GTE predicate on the "color" field.
func ColorGTE(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldGTE(FieldColor, v))
}

// ColorLT applies the LT predicate on the "color" field.
func ColorLT(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldLT(FieldColor, v))
}

// ColorLTE applies the LTE predicate on the "color" field.
func ColorLTE(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldLTE(FieldColor, v))
}

// ColorContains applies the Contains predicate on the "color" field.
func ColorContains(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldContains(FieldColor, v))
}

// ColorHasPrefix applies the HasPrefix predicate on the "color" field.
func ColorHasPrefix(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldHasPrefix(FieldColor, v))
}

// ColorHasSuffix applies the HasSuffix predicate on the "color" field.
func ColorHasSuffix(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldHasSuffix(FieldColor, v))
}

// ColorEqualFold applies the EqualFold predicate on the "color" field.
func ColorEqualFold(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldEqualFold(FieldColor, v))
}

// ColorContainsFold applies the ContainsFold predicate on the "color" field.
func ColorContainsFold(v string) predicate.BaseInfo {
	return predicate.BaseInfo(sql.FieldContainsFold(FieldColor, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BaseInfo) predicate.BaseInfo {
	return predicate.BaseInfo(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BaseInfo) predicate.BaseInfo {
	return predicate.BaseInfo(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.BaseInfo) predicate.BaseInfo {
	return predicate.BaseInfo(sql.NotPredicates(p))
}
