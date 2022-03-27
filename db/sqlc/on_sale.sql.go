// Code generated by sqlc. DO NOT EDIT.
// source: on_sale.sql

package db

import (
	"context"
)

const addOnSalePrice = `-- name: AddOnSalePrice :one
UPDATE on_sale
SET price = price + $1
WHERE id = $2
RETURNING id, brand, link, price, saleper, created_at
`

type AddOnSalePriceParams struct {
	Price int64 `json:"price"`
	ID    int64 `json:"id"`
}

func (q *Queries) AddOnSalePrice(ctx context.Context, arg AddOnSalePriceParams) (OnSale, error) {
	row := q.db.QueryRowContext(ctx, addOnSalePrice, arg.Price, arg.ID)
	var i OnSale
	err := row.Scan(
		&i.ID,
		&i.Brand,
		&i.Link,
		&i.Price,
		&i.Saleper,
		&i.CreatedAt,
	)
	return i, err
}

const createOnSale = `-- name: CreateOnSale :one
INSERT INTO on_sale (
  brand,
  link,
  price,
  saleper
) VALUES (
  $1, $2, $3, $4
) RETURNING id, brand, link, price, saleper, created_at
`

type CreateOnSaleParams struct {
	Brand   string `json:"brand"`
	Link    string `json:"link"`
	Price   int64  `json:"price"`
	Saleper int64  `json:"saleper"`
}

func (q *Queries) CreateOnSale(ctx context.Context, arg CreateOnSaleParams) (OnSale, error) {
	row := q.db.QueryRowContext(ctx, createOnSale,
		arg.Brand,
		arg.Link,
		arg.Price,
		arg.Saleper,
	)
	var i OnSale
	err := row.Scan(
		&i.ID,
		&i.Brand,
		&i.Link,
		&i.Price,
		&i.Saleper,
		&i.CreatedAt,
	)
	return i, err
}

const deleteOnSale = `-- name: DeleteOnSale :exec
DELETE FROM on_sale
WHERE id = $1
`

func (q *Queries) DeleteOnSale(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteOnSale, id)
	return err
}

const getOnSale = `-- name: GetOnSale :one
SELECT id, brand, link, price, saleper, created_at FROM on_sale
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetOnSale(ctx context.Context, id int64) (OnSale, error) {
	row := q.db.QueryRowContext(ctx, getOnSale, id)
	var i OnSale
	err := row.Scan(
		&i.ID,
		&i.Brand,
		&i.Link,
		&i.Price,
		&i.Saleper,
		&i.CreatedAt,
	)
	return i, err
}

const getOnSaleForUpdate = `-- name: GetOnSaleForUpdate :one
SELECT id, brand, link, price, saleper, created_at FROM on_sale
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetOnSaleForUpdate(ctx context.Context, id int64) (OnSale, error) {
	row := q.db.QueryRowContext(ctx, getOnSaleForUpdate, id)
	var i OnSale
	err := row.Scan(
		&i.ID,
		&i.Brand,
		&i.Link,
		&i.Price,
		&i.Saleper,
		&i.CreatedAt,
	)
	return i, err
}

const listOnSale = `-- name: ListOnSale :many
SELECT id, brand, link, price, saleper, created_at FROM on_sale
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListOnSaleParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListOnSale(ctx context.Context, arg ListOnSaleParams) ([]OnSale, error) {
	rows, err := q.db.QueryContext(ctx, listOnSale, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []OnSale{}
	for rows.Next() {
		var i OnSale
		if err := rows.Scan(
			&i.ID,
			&i.Brand,
			&i.Link,
			&i.Price,
			&i.Saleper,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOnSale = `-- name: UpdateOnSale :one
UPDATE on_sale
SET price = $2
WHERE id = $1
RETURNING id, brand, link, price, saleper, created_at
`

type UpdateOnSaleParams struct {
	ID    int64 `json:"id"`
	Price int64 `json:"price"`
}

func (q *Queries) UpdateOnSale(ctx context.Context, arg UpdateOnSaleParams) (OnSale, error) {
	row := q.db.QueryRowContext(ctx, updateOnSale, arg.ID, arg.Price)
	var i OnSale
	err := row.Scan(
		&i.ID,
		&i.Brand,
		&i.Link,
		&i.Price,
		&i.Saleper,
		&i.CreatedAt,
	)
	return i, err
}
