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
    IMAGE TEXT NOT NULL
);


