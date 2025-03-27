// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package custSql

import (
	"database/sql"
)

type Follow struct {
	ID             interface{}
	Email          interface{}
	MarketHashName sql.NullString
}

type Item struct {
	AssetID        interface{}
	MarketHashName sql.NullString
	Price          interface{}
	Time           interface{}
}

type User struct {
	Name     sql.NullString
	Email    string
	Password sql.NullString
}
