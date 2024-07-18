// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BasesColumns holds the columns for the "bases" table.
	BasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// BasesTable holds the schema information for the "bases" table.
	BasesTable = &schema.Table{
		Name:       "bases",
		Columns:    BasesColumns,
		PrimaryKey: []*schema.Column{BasesColumns[0]},
	}
	// BaseInfosColumns holds the columns for the "base_infos" table.
	BaseInfosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "base_id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "logo", Type: field.TypeString},
		{Name: "color", Type: field.TypeString},
	}
	// BaseInfosTable holds the schema information for the "base_infos" table.
	BaseInfosTable = &schema.Table{
		Name:       "base_infos",
		Columns:    BaseInfosColumns,
		PrimaryKey: []*schema.Column{BaseInfosColumns[0]},
	}
	// BaseMembersColumns holds the columns for the "base_members" table.
	BaseMembersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "base_id", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeString},
	}
	// BaseMembersTable holds the schema information for the "base_members" table.
	BaseMembersTable = &schema.Table{
		Name:       "base_members",
		Columns:    BaseMembersColumns,
		PrimaryKey: []*schema.Column{BaseMembersColumns[0]},
	}
	// BaseNavStatesColumns holds the columns for the "base_nav_states" table.
	BaseNavStatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "base_id", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeString},
		{Name: "index", Type: field.TypeInt},
	}
	// BaseNavStatesTable holds the schema information for the "base_nav_states" table.
	BaseNavStatesTable = &schema.Table{
		Name:       "base_nav_states",
		Columns:    BaseNavStatesColumns,
		PrimaryKey: []*schema.Column{BaseNavStatesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BasesTable,
		BaseInfosTable,
		BaseMembersTable,
		BaseNavStatesTable,
	}
)

func init() {
}
