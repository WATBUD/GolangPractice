// Code generated by ent, DO NOT EDIT.

package baseinfo

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the baseinfo type in the database.
	Label = "base_info"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldBaseID holds the string denoting the base_id field in the database.
	FieldBaseID = "base_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldLogo holds the string denoting the logo field in the database.
	FieldLogo = "logo"
	// FieldColor holds the string denoting the color field in the database.
	FieldColor = "color"
	// Table holds the table name of the baseinfo in the database.
	Table = "base_infos"
)

// Columns holds all SQL columns for baseinfo fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldDeletedAt,
	FieldUpdatedAt,
	FieldBaseID,
	FieldName,
	FieldLogo,
	FieldColor,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the BaseInfo queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByBaseID orders the results by the base_id field.
func ByBaseID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBaseID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByLogo orders the results by the logo field.
func ByLogo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLogo, opts...).ToFunc()
}

// ByColor orders the results by the color field.
func ByColor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldColor, opts...).ToFunc()
}