-- name: CreatetProduct :one
INSERT INTO first (
  brand,
  link,
  price
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetProduct :one
SELECT * FROM first
WHERE id = $1 LIMIT 1;

-- name: GetProductForUpdate :one
SELECT * FROM first
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: UpdateProduct :one
UPDATE first
SET price = $2
WHERE id = $1
RETURNING *;

-- name: ListProduct :many
SELECT * FROM first
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteFirstProduct :exec
DELETE FROM first
WHERE id = $1;

-- name: GetLengthOfFirst :one
SELECT count(id) 
FROM first;
