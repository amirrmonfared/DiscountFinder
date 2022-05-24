package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
	StoreProduct(ctx context.Context, arg CreateProductParams) (CreateProductResult, error)
	StoreOnSale(ctx context.Context, arg CreateOnSaleParams) (CreateOnSaleResult, error)
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
	Product Product `json:"product"`
}

func (store *SQLStore) StoreProduct(ctx context.Context, arg CreateProductParams) (CreateProductResult, error) {
	var result CreateProductResult

	err := store.execTx(ctx, func(q *Queries) error {

		var err error

		result.Product, err = q.CreateProduct(ctx, CreateProductParams{
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

type CreateOnSaleResult struct {
	OnSale OnSale `json:"pn_sale"`
}

func (store *SQLStore) StoreOnSale(ctx context.Context, arg CreateOnSaleParams) (CreateOnSaleResult, error) {
	var result CreateOnSaleResult

	err := store.execTx(ctx, func(q *Queries) error {

		var err error

		result.OnSale, err = q.CreateOnSale(ctx, CreateOnSaleParams{
			Brand:         arg.Brand,
			Link:          arg.Link,
			Price:         arg.Price,
			PreviousPrice: arg.PreviousPrice,
		})
		if err != nil {
			return err
		}
		return err
	})

	return result, err
}
