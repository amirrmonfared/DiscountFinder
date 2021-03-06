package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
	CreateProduct(ctx context.Context, arg CreateFirstProductParams) (CreateProductResult, error)
	LengthOfFirst(ctx context.Context) (int64, error)
	LengthOfOnSale(ctx context.Context) (int64, error)
}

// SQLStore provides all functions to excute db queries
type SQLStore struct {
	db *sql.DB
	*Queries
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type CreateProductResult struct {
	First First `json:"first"`
}

func (store *SQLStore) CreateProduct(ctx context.Context, arg CreateFirstProductParams) (CreateProductResult, error) {
	var result CreateProductResult

	err := store.execTx(ctx, func(q *Queries) error {

		var err error

		result.First, err = q.CreateFirstProduct(ctx, CreateFirstProductParams{
			Brand: arg.Brand,
			Link:  arg.Link,
			Price: arg.Price,
		})
		if err != nil {
			return err
		}
		return err
	})

	return result, err
}

func (store *SQLStore) LengthOfFirst(ctx context.Context) (int64, error) {
	var result int64

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result, err = q.GetLengthOfFirst(ctx)
		if err != nil {
			return err
		}
		return err
	})
	fmt.Println(result)

	return result, err
}

func (store *SQLStore) LengthOfOnSale(ctx context.Context) (int64, error) {
	var result int64

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result, err = q.GetLengthOnSale(ctx)
		if err != nil {
			return err
		}
		return err
	})
	fmt.Println(result)

	return result, err
}
