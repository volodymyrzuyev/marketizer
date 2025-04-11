-- name: Get_All_Users :many
SELECT
    *
FROM
    users;

-- name: Add_User :exec
INSERT INTO 
    users (NAME, EMAIL, PASSWORD)
VALUES
    (?1, ?2, ?3);

-- name: Get_User :one
SELECT *
FROM users
WHERE EMAIL = ?1;

-- name: AddAsset :exec
INSERT INTO
    items (ASSET_ID, MARKET_HASH_NAME, PRICE, APPID, TIME, IMAGE)
VALUES
    (?1, ?2, ?3, ?4, ?5, ?6);

-- name: GetItemsPriceDESC :many
SELECT * 
FROM items 
WHERE MARKET_HASH_NAME LIKE ?1
ORDER BY price DESC;

-- name: GetItemsPriceASC :many
SELECT * 
FROM items 
WHERE MARKET_HASH_NAME LIKE ?1
ORDER BY price ASC;

-- name: GetItemsNameDESC :many
SELECT * 
FROM items 
WHERE MARKET_HASH_NAME LIKE ?1
ORDER BY MARKET_HASH_NAME DESC;

-- name: GetItemsNameASC :many
SELECT * 
FROM items 
WHERE MARKET_HASH_NAME LIKE ?1
ORDER BY MARKET_HASH_NAME ASC;

-- name: GetItemsTimeDESC :many
SELECT * 
FROM items 
WHERE MARKET_HASH_NAME LIKE ?1
ORDER BY time DESC;

-- name: GetItemsTimeASC :many
SELECT * 
FROM items 
WHERE MARKET_HASH_NAME LIKE ?1
ORDER BY time ASC;


-- name: GetItemsToNotify :many
SELECT 
    i.ASSET_ID
FROM items i, follows f
WHERE i.MARKET_HASH_NAME = f.MARKET_HASH_NAME 
    AND f.EMAIL = ?1 
    AND NOT EXISTS (
        SELECT n.ID
        FROM notifications n 
        WHERE n.ASSET_ID = i.ASSET_ID);

-- name: SetItemAsNotified :exec
INSERT INTO 
    notifications (ASSET_ID, EMAIL)
VALUES
    (?1, ?2);

-- name: AddToFollows :exec
INSERT INTO
    follows (EMAIL, MARKET_HASH_NAME)
VALUES
    (?1, ?2);

-- name: GetFollows :many
SELECT MARKET_HASH_NAME
FROM follows
WHERE EMAIL = ?1;

-- name: GetFollowItemsPriceASC :many
SELECT items.*
FROM items, follows
WHERE items.MARKET_HASH_NAME = follows.MARKET_HASH_NAME 
    AND follows.EMAIL = ?1 
    AND items.MARKET_HASH_NAME LIKE ?2
GROUP BY items.MARKET_HASH_NAME
ORDER BY price ASC;

-- name: GetFollowItemsPriceDSC :many
SELECT items.*
FROM items, follows
WHERE items.MARKET_HASH_NAME = follows.MARKET_HASH_NAME 
    AND follows.EMAIL = ?1 
    AND items.MARKET_HASH_NAME LIKE ?2
GROUP BY items.MARKET_HASH_NAME
ORDER BY price DESC;

-- name: GetFollowItemsNameASC :many
SELECT items.*
FROM items, follows
WHERE items.MARKET_HASH_NAME = follows.MARKET_HASH_NAME 
    AND follows.EMAIL = ?1 
    AND items.MARKET_HASH_NAME LIKE ?2
GROUP BY items.MARKET_HASH_NAME
ORDER BY items.MARKET_HASH_NAME ASC;

-- name: GetFollowItemsNameDSC :many
SELECT items.*
FROM items, follows
WHERE items.MARKET_HASH_NAME = follows.MARKET_HASH_NAME 
    AND follows.EMAIL = ?1 
    AND items.MARKET_HASH_NAME LIKE ?2
GROUP BY items.MARKET_HASH_NAME
ORDER BY items.MARKET_HASH_NAME DESC;

-- name: GetFollowItemsTimeDSC :many
SELECT items.*
FROM items, follows
WHERE items.MARKET_HASH_NAME = follows.MARKET_HASH_NAME 
    AND follows.EMAIL = ?1 
    AND items.MARKET_HASH_NAME LIKE ?2
GROUP BY items.MARKET_HASH_NAME
ORDER BY time ASC;

-- name: GetFollowItemsTimeASC :many
SELECT items.*
FROM items, follows
WHERE items.MARKET_HASH_NAME = follows.MARKET_HASH_NAME 
    AND follows.EMAIL = ?1 
    AND items.MARKET_HASH_NAME LIKE ?2
GROUP BY items.MARKET_HASH_NAME
ORDER BY time DESC;

-- name: Create_table1 :exec
CREATE TABLE IF NOT EXISTS users(
    NAME TEXT,
    EMAIL TEXT PRIMARY KEY,
    PASSWORD TEXT
);
-- name: Create_table2 :exec
CREATE TABLE IF NOT EXISTS follows(
    ID INTEGER PRIMARY KEY,
    EMAIL INTEGER,
    MARKET_HASH_NAME TEXT
);
-- name: Create_table3 :exec
CREATE TABLE IF NOT EXISTS items(
    ASSET_ID INTEGER PRIMARY KEY,
    MARKET_HASH_NAME TEXT,
    PRICE INTEGER,
    APPID INTEGER,
    TIME INTEGER,
    IMAGE TEXT,
    NOTIFIED BOOLEAN
);

-- name: Create_table4 :exec
CREATE TABLE IF NOT EXISTS notifications(
    ID INTEGER PRIMARY KEY,
    ASSET_ID INTEGER,
    EMAIL TEXT
);

