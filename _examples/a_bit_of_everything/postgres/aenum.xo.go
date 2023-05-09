package postgres

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql/driver"
	"fmt"
)

// AEnum is the 'a_enum' enum type from schema 'public'.
type AEnum uint16

// AEnum values.
const (
	// AEnumOne is the 'ONE' a_enum.
	AEnumOne AEnum = 1
	// AEnumTwo is the 'TWO' a_enum.
	AEnumTwo AEnum = 2
)

// String satisfies the [fmt.Stringer] interface.
func (ae AEnum) String() string {
	switch ae {
	case AEnumOne:
		return "ONE"
	case AEnumTwo:
		return "TWO"
	}
	return fmt.Sprintf("AEnum(%d)", ae)
}

// MarshalText marshals [AEnum] into text.
func (ae AEnum) MarshalText() ([]byte, error) {
	return []byte(ae.String()), nil
}

// UnmarshalText unmarshals [AEnum] from text.
func (ae *AEnum) UnmarshalText(buf []byte) error {
	switch str := string(buf); str {
	case "ONE":
		*ae = AEnumOne
	case "TWO":
		*ae = AEnumTwo
	default:
		return ErrInvalidAEnum(str)
	}
	return nil
}

// Value satisfies the [driver.Valuer] interface.
func (ae AEnum) Value() (driver.Value, error) {
	return ae.String(), nil
}

// Scan satisfies the [sql.Scanner] interface.
func (ae *AEnum) Scan(v interface{}) error {
	switch x := v.(type) {
	case []byte:
		return ae.UnmarshalText(x)
	case string:
		return ae.UnmarshalText([]byte(x))
	}
	return ErrInvalidAEnum(fmt.Sprintf("%T", v))
}

// NullAEnum represents a null 'a_enum' enum for schema 'public'.
type NullAEnum struct {
	AEnum AEnum
	// Valid is true if [AEnum] is not null.
	Valid bool
}

// Value satisfies the [driver.Valuer] interface.
func (nae NullAEnum) Value() (driver.Value, error) {
	if !nae.Valid {
		return nil, nil
	}
	return nae.AEnum.Value()
}

// Scan satisfies the [sql.Scanner] interface.
func (nae *NullAEnum) Scan(v interface{}) error {
	if v == nil {
		nae.AEnum, nae.Valid = 0, false
		return nil
	}
	err := nae.AEnum.Scan(v)
	nae.Valid = err == nil
	return err
}

// ErrInvalidAEnum is the invalid [AEnum] error.
type ErrInvalidAEnum string

// Error satisfies the error interface.
func (err ErrInvalidAEnum) Error() string {
	return fmt.Sprintf("invalid AEnum(%s)", string(err))
}
