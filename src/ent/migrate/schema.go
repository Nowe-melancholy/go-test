// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// MiddlewaresColumns holds the columns for the "middlewares" table.
	MiddlewaresColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "l_id", Type: field.TypeString, SchemaType: map[string]string{"postgres": "int"}},
		{Name: "d_id", Type: field.TypeString, SchemaType: map[string]string{"postgres": "int"}},
		{Name: "sys_id", Type: field.TypeString, SchemaType: map[string]string{"postgres": "int"}},
	}
	// MiddlewaresTable holds the schema information for the "middlewares" table.
	MiddlewaresTable = &schema.Table{
		Name:       "middlewares",
		Columns:    MiddlewaresColumns,
		PrimaryKey: []*schema.Column{MiddlewaresColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		MiddlewaresTable,
	}
)

func init() {
}