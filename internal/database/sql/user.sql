-- name: GetUser :one
SELECT user_id, user_name, user_name, create_at FROM "User"
WHERE user_id = $1 LIMIT 1;

-- name: GetListUsers :many
SELECT * FROM "User"
ORDER BY user_name;

-- name: CreateUser :one
INSERT INTO "User"
(user_name, user_email, first_name, last_name, create_at, active)
VALUES($1, $2, $3, $4, $5, $6) 
RETURNING *;

-- name: UpdateUser :one
UPDATE "User"
SET user_name=$2, first_name=$3, last_name=$4, active=$5
WHERE user_id=$1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM "User"
WHERE user_id = $1;