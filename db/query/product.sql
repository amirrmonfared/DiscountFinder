-- name: CreateProduct :one
INSERT INTO products (
  brand,
  link,
  price
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetProduct :one
SELECT * FROM products
WHERE id = $1 LIMIT 1;

-- name: UpdateProduct :one
UPDATE products
SET price = $2
WHERE id = $1
RETURNING *;

-- name: ListProduct :many
SELECT * FROM products
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;

-- name: GetLengthOfProducts :one
SELECT count(id) 
FROM products;
