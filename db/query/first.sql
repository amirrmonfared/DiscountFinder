-- name: CreateFirstProduct :one
INSERT INTO first (
  brand,
  link,
  price
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetFirstProduct :one
SELECT * FROM first
WHERE id = $1 LIMIT 1;

-- name: GetFirstProductForUpdate :one
SELECT * FROM first
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: UpdateFirstProduct :one
UPDATE first
SET price = $2
WHERE id = $1
RETURNING *;

-- name: ListFirstProduct :many
SELECT * FROM first
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: AddfirstProductPrice :one
UPDATE first
SET price = price + sqlc.arg(price)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteFirstProduct :exec
DELETE FROM first
WHERE id = $1;
