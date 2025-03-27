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


-- name: Create_table1 :exec
CREATE TABLE IF NOT EXISTS users(
    NAME TEXT,
    EMAIL TEXT PRIMARY KEY,
    PASSWORD TEXT
);
-- name: Create_table2 :exec
CREATE TABLE IF NOT EXISTS follows(
    ID NUMBER PRIMARY KEY,
    EMAIL NUMBER,
    MARKET_HASH_NAME TEXT
);
-- name: Create_table3 :exec
CREATE TABLE IF NOT EXISTS items(
    ASSET_ID NUMBER PRIMARY KEY,
    MARKET_HASH_NAME TEXT,
    PRICE NUMBER,
    TIME NUMBER
);


