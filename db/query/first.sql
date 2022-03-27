-- name: CreateFirst :one
INSERT INTO first (
  brand,
  link,
  price
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetFirst :one
SELECT * FROM first
WHERE id = $1 LIMIT 1;

-- name: GetFirstForUpdate :one
SELECT * FROM first
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: UpdateFirst :one
UPDATE first
SET price = $2
WHERE id = $1
RETURNING *;

-- name: ListFirst :many
SELECT * FROM first
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: AddfirstPrice :one
UPDATE first
SET price = price + sqlc.arg(price)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteFirst :exec
DELETE FROM first
WHERE id = $1;
