// Code generated by sqlc. DO NOT EDIT.

package gen

import (
	"database/sql"
)

type Author struct {
	ID   int32
	Name string
	Bio  sql.NullString
}