package ischema

// Code generated by xo. DO NOT EDIT.

// ColumnPrivilege represents a row from 'information_schema.column_privileges'.
type ColumnPrivilege struct {
	Grantor       SQLIdentifier `json:"grantor"`        // grantor
	Grantee       SQLIdentifier `json:"grantee"`        // grantee
	TableCatalog  SQLIdentifier `json:"table_catalog"`  // table_catalog
	TableSchema   SQLIdentifier `json:"table_schema"`   // table_schema
	TableName     SQLIdentifier `json:"table_name"`     // table_name
	ColumnName    SQLIdentifier `json:"column_name"`    // column_name
	PrivilegeType CharacterData `json:"privilege_type"` // privilege_type
	IsGrantable   YesOrNo       `json:"is_grantable"`   // is_grantable
}
