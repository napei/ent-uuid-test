// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// TestModelsColumns holds the columns for the "test_models" table.
	TestModelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "test", Type: field.TypeString},
	}
	// TestModelsTable holds the schema information for the "test_models" table.
	TestModelsTable = &schema.Table{
		Name:        "test_models",
		Columns:     TestModelsColumns,
		PrimaryKey:  []*schema.Column{TestModelsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TestModelsTable,
	}
)

func init() {
}
