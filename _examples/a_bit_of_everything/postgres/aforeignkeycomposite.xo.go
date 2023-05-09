package postgres

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// AForeignKeyComposite represents a row from 'public.a_foreign_key_composite'.
type AForeignKeyComposite struct {
	AKey1 sql.NullInt64 `json:"a_key1"` // a_key1
	AKey2 sql.NullInt64 `json:"a_key2"` // a_key2
}

// APrimaryComposite returns the APrimaryComposite associated with the [AForeignKeyComposite]'s (AKey1, AKey2).
//
// Generated from foreign key 'a_foreign_key_composite_a_key1_a_key2_fkey'.
func (afkc *AForeignKeyComposite) APrimaryComposite(ctx context.Context, db DB) (*APrimaryComposite, error) {
	return APrimaryCompositeByAKey1AKey2(ctx, db, int(afkc.AKey1.Int64), int(afkc.AKey2.Int64))
}
