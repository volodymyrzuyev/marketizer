// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: queries.sql

package custSql

import (
	"context"
	"database/sql"
)

const addAsset = `-- name: AddAsset :exec
INSERT INTO
    items (ASSET_ID, MARKET_HASH_NAME, PRICE, APPID, TIME, IMAGE)
VALUES
    (?1, ?2, ?3, ?4, ?5, ?6)
`

type AddAssetParams struct {
	AssetID        int64
	MarketHashName string
	Price          int64
	Appid          int64
	Time           int64
	Image          string
}

func (q *Queries) AddAsset(ctx context.Context, arg AddAssetParams) error {
	_, err := q.db.ExecContext(ctx, addAsset,
		arg.AssetID,
		arg.MarketHashName,
		arg.Price,
		arg.Appid,
		arg.Time,
		arg.Image,
	)
	return err
}

const addToFollows = `-- name: AddToFollows :exec
INSERT INTO
    follows (EMAIL, MARKET_HASH_NAME)
VALUES
    (?1, ?2)
`

type AddToFollowsParams struct {
	Email          string
	MarketHashName string
}

func (q *Queries) AddToFollows(ctx context.Context, arg AddToFollowsParams) error {
	_, err := q.db.ExecContext(ctx, addToFollows, arg.Email, arg.MarketHashName)
	return err
}

const add_User = `-- name: Add_User :exec
INSERT INTO 
    users (NAME, EMAIL, PASSWORD)
VALUES
    (?1, ?2, ?3)
`

type Add_UserParams struct {
	Name     sql.NullString
	Email    string
	Password sql.NullString
}

func (q *Queries) Add_User(ctx context.Context, arg Add_UserParams) error {
	_, err := q.db.ExecContext(ctx, add_User, arg.Name, arg.Email, arg.Password)
	return err
}

const create_table1 = `-- name: Create_table1 :exec
CREATE TABLE IF NOT EXISTS users(
    NAME TEXT,
    EMAIL TEXT PRIMARY KEY,
    PASSWORD TEXT
)
`

func (q *Queries) Create_table1(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, create_table1)
	return err
}

const create_table2 = `-- name: Create_table2 :exec
CREATE TABLE IF NOT EXISTS follows(
    ID INTEGER PRIMARY KEY,
    EMAIL INTEGER,
    MARKET_HASH_NAME TEXT
)
`

func (q *Queries) Create_table2(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, create_table2)
	return err
}

const create_table3 = `-- name: Create_table3 :exec
CREATE TABLE IF NOT EXISTS items(
    ASSET_ID INTEGER PRIMARY KEY,
    MARKET_HASH_NAME TEXT,
    PRICE INTEGER,
    APPID INTEGER,
    TIME INTEGER,
    IMAGE TEXT,
    NOTIFIED BOOLEAN
)
`

func (q *Queries) Create_table3(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, create_table3)
	return err
}

const create_table4 = `-- name: Create_table4 :exec
CREATE TABLE IF NOT EXISTS notifications(
    ID INTEGER PRIMARY KEY,
    ASSET_ID INTEGER,
    EMAIL TEXT
)
`

func (q *Queries) Create_table4(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, create_table4)
	return err
}

const getFollowItemsNameASC = `-- name: GetFollowItemsNameASC :many
SELECT items.asset_id, items.market_hash_name, items.price, items.appid, items.time, items.image
FROM items, follows
WHERE items.MARKET_HASH_NAME = follows.MARKET_HASH_NAME 
    AND follows.EMAIL = ?1 
    AND items.MARKET_HASH_NAME LIKE ?2
GROUP BY items.MARKET_HASH_NAME
ORDER BY items.MARKET_HASH_NAME ASC
`

type GetFollowItemsNameASCParams struct {
	Email          string
	MarketHashName string
}

func (q *Queries) GetFollowItemsNameASC(ctx context.Context, arg GetFollowItemsNameASCParams) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, getFollowItemsNameASC, arg.Email, arg.MarketHashName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.AssetID,
			&i.MarketHashName,
			&i.Price,
			&i.Appid,
			&i.Time,
			&i.Image,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFollowItemsNameDSC = `-- name: GetFollowItemsNameDSC :many
SELECT items.asset_id, items.market_hash_name, items.price, items.appid, items.time, items.image
FROM items, follows
WHERE items.MARKET_HASH_NAME = follows.MARKET_HASH_NAME 
    AND follows.EMAIL = ?1 
    AND items.MARKET_HASH_NAME LIKE ?2
GROUP BY items.MARKET_HASH_NAME
ORDER BY items.MARKET_HASH_NAME DESC
`

type GetFollowItemsNameDSCParams struct {
	Email          string
	MarketHashName string
}

func (q *Queries) GetFollowItemsNameDSC(ctx context.Context, arg GetFollowItemsNameDSCParams) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, getFollowItemsNameDSC, arg.Email, arg.MarketHashName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.AssetID,
			&i.MarketHashName,
			&i.Price,
			&i.Appid,
			&i.Time,
			&i.Image,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFollowItemsPriceASC = `-- name: GetFollowItemsPriceASC :many
SELECT items.asset_id, items.market_hash_name, items.price, items.appid, items.time, items.image
FROM items, follows
WHERE items.MARKET_HASH_NAME = follows.MARKET_HASH_NAME 
    AND follows.EMAIL = ?1 
    AND items.MARKET_HASH_NAME LIKE ?2
GROUP BY items.MARKET_HASH_NAME
ORDER BY price ASC
`

type GetFollowItemsPriceASCParams struct {
	Email          string
	MarketHashName string
}

func (q *Queries) GetFollowItemsPriceASC(ctx context.Context, arg GetFollowItemsPriceASCParams) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, getFollowItemsPriceASC, arg.Email, arg.MarketHashName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.AssetID,
			&i.MarketHashName,
			&i.Price,
			&i.Appid,
			&i.Time,
			&i.Image,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFollowItemsPriceDSC = `-- name: GetFollowItemsPriceDSC :many
SELECT items.asset_id, items.market_hash_name, items.price, items.appid, items.time, items.image
FROM items, follows
WHERE items.MARKET_HASH_NAME = follows.MARKET_HASH_NAME 
    AND follows.EMAIL = ?1 
    AND items.MARKET_HASH_NAME LIKE ?2
GROUP BY items.MARKET_HASH_NAME
ORDER BY price DESC
`

type GetFollowItemsPriceDSCParams struct {
	Email          string
	MarketHashName string
}

func (q *Queries) GetFollowItemsPriceDSC(ctx context.Context, arg GetFollowItemsPriceDSCParams) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, getFollowItemsPriceDSC, arg.Email, arg.MarketHashName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.AssetID,
			&i.MarketHashName,
			&i.Price,
			&i.Appid,
			&i.Time,
			&i.Image,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFollowItemsTimeASC = `-- name: GetFollowItemsTimeASC :many
SELECT items.asset_id, items.market_hash_name, items.price, items.appid, items.time, items.image
FROM items, follows
WHERE items.MARKET_HASH_NAME = follows.MARKET_HASH_NAME 
    AND follows.EMAIL = ?1 
    AND items.MARKET_HASH_NAME LIKE ?2
GROUP BY items.MARKET_HASH_NAME
ORDER BY time DESC
`

type GetFollowItemsTimeASCParams struct {
	Email          string
	MarketHashName string
}

func (q *Queries) GetFollowItemsTimeASC(ctx context.Context, arg GetFollowItemsTimeASCParams) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, getFollowItemsTimeASC, arg.Email, arg.MarketHashName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.AssetID,
			&i.MarketHashName,
			&i.Price,
			&i.Appid,
			&i.Time,
			&i.Image,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFollowItemsTimeDSC = `-- name: GetFollowItemsTimeDSC :many
SELECT items.asset_id, items.market_hash_name, items.price, items.appid, items.time, items.image
FROM items, follows
WHERE items.MARKET_HASH_NAME = follows.MARKET_HASH_NAME 
    AND follows.EMAIL = ?1 
    AND items.MARKET_HASH_NAME LIKE ?2
GROUP BY items.MARKET_HASH_NAME
ORDER BY time ASC
`

type GetFollowItemsTimeDSCParams struct {
	Email          string
	MarketHashName string
}

func (q *Queries) GetFollowItemsTimeDSC(ctx context.Context, arg GetFollowItemsTimeDSCParams) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, getFollowItemsTimeDSC, arg.Email, arg.MarketHashName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.AssetID,
			&i.MarketHashName,
			&i.Price,
			&i.Appid,
			&i.Time,
			&i.Image,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFollows = `-- name: GetFollows :many
SELECT MARKET_HASH_NAME
FROM follows
WHERE EMAIL = ?1
`

func (q *Queries) GetFollows(ctx context.Context, email string) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getFollows, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var market_hash_name string
		if err := rows.Scan(&market_hash_name); err != nil {
			return nil, err
		}
		items = append(items, market_hash_name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getItemsNameASC = `-- name: GetItemsNameASC :many
SELECT asset_id, market_hash_name, price, appid, time, image 
FROM items 
WHERE MARKET_HASH_NAME LIKE ?1
ORDER BY MARKET_HASH_NAME ASC
`

func (q *Queries) GetItemsNameASC(ctx context.Context, marketHashName string) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, getItemsNameASC, marketHashName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.AssetID,
			&i.MarketHashName,
			&i.Price,
			&i.Appid,
			&i.Time,
			&i.Image,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getItemsNameDESC = `-- name: GetItemsNameDESC :many
SELECT asset_id, market_hash_name, price, appid, time, image 
FROM items 
WHERE MARKET_HASH_NAME LIKE ?1
ORDER BY MARKET_HASH_NAME DESC
`

func (q *Queries) GetItemsNameDESC(ctx context.Context, marketHashName string) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, getItemsNameDESC, marketHashName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.AssetID,
			&i.MarketHashName,
			&i.Price,
			&i.Appid,
			&i.Time,
			&i.Image,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getItemsPriceASC = `-- name: GetItemsPriceASC :many
SELECT asset_id, market_hash_name, price, appid, time, image 
FROM items 
WHERE MARKET_HASH_NAME LIKE ?1
ORDER BY price ASC
`

func (q *Queries) GetItemsPriceASC(ctx context.Context, marketHashName string) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, getItemsPriceASC, marketHashName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.AssetID,
			&i.MarketHashName,
			&i.Price,
			&i.Appid,
			&i.Time,
			&i.Image,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getItemsPriceDESC = `-- name: GetItemsPriceDESC :many
SELECT asset_id, market_hash_name, price, appid, time, image 
FROM items 
WHERE MARKET_HASH_NAME LIKE ?1
ORDER BY price DESC
`

func (q *Queries) GetItemsPriceDESC(ctx context.Context, marketHashName string) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, getItemsPriceDESC, marketHashName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.AssetID,
			&i.MarketHashName,
			&i.Price,
			&i.Appid,
			&i.Time,
			&i.Image,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getItemsTimeASC = `-- name: GetItemsTimeASC :many
SELECT asset_id, market_hash_name, price, appid, time, image 
FROM items 
WHERE MARKET_HASH_NAME LIKE ?1
ORDER BY time ASC
`

func (q *Queries) GetItemsTimeASC(ctx context.Context, marketHashName string) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, getItemsTimeASC, marketHashName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.AssetID,
			&i.MarketHashName,
			&i.Price,
			&i.Appid,
			&i.Time,
			&i.Image,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getItemsTimeDESC = `-- name: GetItemsTimeDESC :many
SELECT asset_id, market_hash_name, price, appid, time, image 
FROM items 
WHERE MARKET_HASH_NAME LIKE ?1
ORDER BY time DESC
`

func (q *Queries) GetItemsTimeDESC(ctx context.Context, marketHashName string) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, getItemsTimeDESC, marketHashName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.AssetID,
			&i.MarketHashName,
			&i.Price,
			&i.Appid,
			&i.Time,
			&i.Image,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getItemsToNotify = `-- name: GetItemsToNotify :many
SELECT 
    i.ASSET_ID
FROM items i, follows f
WHERE i.MARKET_HASH_NAME = f.MARKET_HASH_NAME 
    AND f.EMAIL = ?1 
    AND NOT EXISTS (
        SELECT n.ID
        FROM notifications n 
        WHERE n.ASSET_ID = i.ASSET_ID)
`

func (q *Queries) GetItemsToNotify(ctx context.Context, email string) ([]int64, error) {
	rows, err := q.db.QueryContext(ctx, getItemsToNotify, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int64
	for rows.Next() {
		var asset_id int64
		if err := rows.Scan(&asset_id); err != nil {
			return nil, err
		}
		items = append(items, asset_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const get_All_Users = `-- name: Get_All_Users :many
SELECT
    name, email, password
FROM
    users
`

func (q *Queries) Get_All_Users(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, get_All_Users)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(&i.Name, &i.Email, &i.Password); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const get_User = `-- name: Get_User :one
SELECT name, email, password
FROM users
WHERE EMAIL = ?1
`

func (q *Queries) Get_User(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, get_User, email)
	var i User
	err := row.Scan(&i.Name, &i.Email, &i.Password)
	return i, err
}

const removeFromFollows = `-- name: RemoveFromFollows :exec
DELETE FROM follows
WHERE EMAIL = ?1 AND MARKET_HASH_NAME = ?2
`

type RemoveFromFollowsParams struct {
	Email          string
	MarketHashName string
}

func (q *Queries) RemoveFromFollows(ctx context.Context, arg RemoveFromFollowsParams) error {
	_, err := q.db.ExecContext(ctx, removeFromFollows, arg.Email, arg.MarketHashName)
	return err
}

const setItemAsNotified = `-- name: SetItemAsNotified :exec
INSERT INTO 
    notifications (ASSET_ID, EMAIL)
VALUES
    (?1, ?2)
`

type SetItemAsNotifiedParams struct {
	AssetID int64
	Email   string
}

func (q *Queries) SetItemAsNotified(ctx context.Context, arg SetItemAsNotifiedParams) error {
	_, err := q.db.ExecContext(ctx, setItemAsNotified, arg.AssetID, arg.Email)
	return err
}

const unsetItemAsNotified = `-- name: UnsetItemAsNotified :exec
DELETE FROM notifications 
WHERE ASSET_ID = ?1 AND EMAIL = ?2
`

type UnsetItemAsNotifiedParams struct {
	AssetID int64
	Email   string
}

func (q *Queries) UnsetItemAsNotified(ctx context.Context, arg UnsetItemAsNotifiedParams) error {
	_, err := q.db.ExecContext(ctx, unsetItemAsNotified, arg.AssetID, arg.Email)
	return err
}
