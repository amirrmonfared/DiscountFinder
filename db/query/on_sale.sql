-- name: CreateOnSale :one
INSERT INTO on_sale (
  brand,
  link,
  price,
  saleper
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetOnSale :one
SELECT * FROM on_sale
WHERE id = $1 LIMIT 1;

-- name: GetOnSaleForUpdate :one
SELECT * FROM on_sale
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: UpdateOnSale :one
UPDATE on_sale
SET price = $2
WHERE id = $1
RETURNING *;

-- name: ListOnSale :many
SELECT * FROM on_sale
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteOnSale :exec
DELETE FROM on_sale
WHERE id = $1;

-- name: GetLengthOnSale :one
SELECT count(id) 
FROM on_sale;
