package ischema

// Code generated by xo. DO NOT EDIT.

// UserMappingOption represents a row from 'information_schema.user_mapping_options'.
type UserMappingOption struct {
	AuthorizationIdentifier SQLIdentifier `json:"authorization_identifier"` // authorization_identifier
	ForeignServerCatalog    SQLIdentifier `json:"foreign_server_catalog"`   // foreign_server_catalog
	ForeignServerName       SQLIdentifier `json:"foreign_server_name"`      // foreign_server_name
	OptionName              SQLIdentifier `json:"option_name"`              // option_name
	OptionValue             CharacterData `json:"option_value"`             // option_value
}
