-- name: CreateSecond :one
INSERT INTO second (
  brand,
  link,
  price
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetSecond :one
SELECT * FROM second
WHERE id = $1 LIMIT 1;

-- name: GetSecondForUpdate :one
SELECT * FROM second
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: UpdateSecond :one
UPDATE second
SET price = $2
WHERE id = $1
RETURNING *;

-- name: ListSecond :many
SELECT * FROM second
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: AddSecondPrice :one
UPDATE second
SET price = price + sqlc.arg(price)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteSecond :exec
DELETE FROM second
WHERE id = $1;
