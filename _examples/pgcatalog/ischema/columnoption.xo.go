package ischema

// Code generated by xo. DO NOT EDIT.

import (
	"github.com/xo/xo/_examples/pgcatalog/pgtypes"
)

// ColumnOption represents a row from 'information_schema.column_options'.
type ColumnOption struct {
	TableCatalog pgtypes.SQLIdentifier `json:"table_catalog"` // table_catalog
	TableSchema  pgtypes.SQLIdentifier `json:"table_schema"`  // table_schema
	TableName    pgtypes.SQLIdentifier `json:"table_name"`    // table_name
	ColumnName   pgtypes.SQLIdentifier `json:"column_name"`   // column_name
	OptionName   pgtypes.SQLIdentifier `json:"option_name"`   // option_name
	OptionValue  pgtypes.CharacterData `json:"option_value"`  // option_value

}