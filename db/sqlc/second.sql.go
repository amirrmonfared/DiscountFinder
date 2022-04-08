// Code generated by sqlc. DO NOT EDIT.
// source: second.sql

package db

import (
	"context"
)

const addSecondPrice = `-- name: AddSecondPrice :one
UPDATE second
SET price = price + $1
WHERE id = $2
RETURNING id, brand, link, price, created_at
`

type AddSecondPriceParams struct {
	Price string `json:"price"`
	ID    int64  `json:"id"`
}

func (q *Queries) AddSecondPrice(ctx context.Context, arg AddSecondPriceParams) (Second, error) {
	row := q.db.QueryRowContext(ctx, addSecondPrice, arg.Price, arg.ID)
	var i Second
	err := row.Scan(
		&i.ID,
		&i.Brand,
		&i.Link,
		&i.Price,
		&i.CreatedAt,
	)
	return i, err
}

const createSecond = `-- name: CreateSecond :one
INSERT INTO second (
  brand,
  link,
  price
) VALUES (
  $1, $2, $3
) RETURNING id, brand, link, price, created_at
`

type CreateSecondParams struct {
	Brand string `json:"brand"`
	Link  string `json:"link"`
	Price string `json:"price"`
}

func (q *Queries) CreateSecond(ctx context.Context, arg CreateSecondParams) (Second, error) {
	row := q.db.QueryRowContext(ctx, createSecond, arg.Brand, arg.Link, arg.Price)
	var i Second
	err := row.Scan(
		&i.ID,
		&i.Brand,
		&i.Link,
		&i.Price,
		&i.CreatedAt,
	)
	return i, err
}

const deleteSecond = `-- name: DeleteSecond :exec
DELETE FROM second
WHERE id = $1
`

func (q *Queries) DeleteSecond(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteSecond, id)
	return err
}

const getSecond = `-- name: GetSecond :one
SELECT id, brand, link, price, created_at FROM second
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetSecond(ctx context.Context, id int64) (Second, error) {
	row := q.db.QueryRowContext(ctx, getSecond, id)
	var i Second
	err := row.Scan(
		&i.ID,
		&i.Brand,
		&i.Link,
		&i.Price,
		&i.CreatedAt,
	)
	return i, err
}

const getSecondForUpdate = `-- name: GetSecondForUpdate :one
SELECT id, brand, link, price, created_at FROM second
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetSecondForUpdate(ctx context.Context, id int64) (Second, error) {
	row := q.db.QueryRowContext(ctx, getSecondForUpdate, id)
	var i Second
	err := row.Scan(
		&i.ID,
		&i.Brand,
		&i.Link,
		&i.Price,
		&i.CreatedAt,
	)
	return i, err
}

const listSecond = `-- name: ListSecond :many
SELECT id, brand, link, price, created_at FROM second
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListSecondParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListSecond(ctx context.Context, arg ListSecondParams) ([]Second, error) {
	rows, err := q.db.QueryContext(ctx, listSecond, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Second{}
	for rows.Next() {
		var i Second
		if err := rows.Scan(
			&i.ID,
			&i.Brand,
			&i.Link,
			&i.Price,
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

const updateSecond = `-- name: UpdateSecond :one
UPDATE second
SET price = $2
WHERE id = $1
RETURNING id, brand, link, price, created_at
`

type UpdateSecondParams struct {
	ID    int64  `json:"id"`
	Price string `json:"price"`
}

func (q *Queries) UpdateSecond(ctx context.Context, arg UpdateSecondParams) (Second, error) {
	row := q.db.QueryRowContext(ctx, updateSecond, arg.ID, arg.Price)
	var i Second
	err := row.Scan(
		&i.ID,
		&i.Brand,
		&i.Link,
		&i.Price,
		&i.CreatedAt,
	)
	return i, err
}
