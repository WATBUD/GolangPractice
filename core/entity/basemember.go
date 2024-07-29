// Code generated by ent, DO NOT EDIT.

package entity

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"mai.today/core/entity/basemember"
)

// BaseMember is the model entity for the BaseMember schema.
type BaseMember struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty" bson:"_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// BaseID holds the value of the "base_id" field.
	BaseID string `json:"base_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID       string `json:"user_id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BaseMember) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case basemember.FieldID, basemember.FieldBaseID, basemember.FieldUserID:
			values[i] = new(sql.NullString)
		case basemember.FieldCreatedAt, basemember.FieldDeletedAt, basemember.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BaseMember fields.
func (bm *BaseMember) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case basemember.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				bm.ID = value.String
			}
		case basemember.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				bm.CreatedAt = value.Time
			}
		case basemember.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				bm.DeletedAt = new(time.Time)
				*bm.DeletedAt = value.Time
			}
		case basemember.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				bm.UpdatedAt = value.Time
			}
		case basemember.FieldBaseID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field base_id", values[i])
			} else if value.Valid {
				bm.BaseID = value.String
			}
		case basemember.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				bm.UserID = value.String
			}
		default:
			bm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BaseMember.
// This includes values selected through modifiers, order, etc.
func (bm *BaseMember) Value(name string) (ent.Value, error) {
	return bm.selectValues.Get(name)
}

// Update returns a builder for updating this BaseMember.
// Note that you need to call BaseMember.Unwrap() before calling this method if this BaseMember
// was returned from a transaction, and the transaction was committed or rolled back.
func (bm *BaseMember) Update() *BaseMemberUpdateOne {
	return NewBaseMemberClient(bm.config).UpdateOne(bm)
}

// Unwrap unwraps the BaseMember entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bm *BaseMember) Unwrap() *BaseMember {
	_tx, ok := bm.config.driver.(*txDriver)
	if !ok {
		panic("entity: BaseMember is not a transactional entity")
	}
	bm.config.driver = _tx.drv
	return bm
}

// String implements the fmt.Stringer.
func (bm *BaseMember) String() string {
	var builder strings.Builder
	builder.WriteString("BaseMember(")
	builder.WriteString(fmt.Sprintf("id=%v, ", bm.ID))
	builder.WriteString("created_at=")
	builder.WriteString(bm.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := bm.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(bm.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("base_id=")
	builder.WriteString(bm.BaseID)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(bm.UserID)
	builder.WriteByte(')')
	return builder.String()
}

// BaseMembers is a parsable slice of BaseMember.
type BaseMembers []*BaseMember