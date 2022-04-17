package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
	CreateProduct(ctx context.Context, arg CreateProductParams) (CreateProductResult, error)
	LengthOfFirst(ctx context.Context) (int64, error)
	ReviewProduct(ctx context.Context, arg CreateSecondParams) (CreateSecondProductResult, error)
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

type CreateProductParams struct {
	Brand string `json:"brand"`
    Link  string `json:"link"`
    Price string `json:"price"`
}

type CreateProductResult struct {
	First First `json:"first"`
}

type CreateSecondProductResult struct {
	Second Second `json:"second"`
}

func (store *SQLStore) CreateProduct(ctx context.Context, arg CreateProductParams) (CreateProductResult, error) {
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
	
	fmt.Println("Product saved")

	return result, err
}

func (store *SQLStore) ReviewProduct(ctx context.Context, arg CreateSecondParams) (CreateSecondProductResult, error) {
	var result CreateSecondProductResult

	err := store.execTx(ctx, func(q *Queries) error {

		var err error

		result.Second, err = q.CreateSecond(ctx, CreateSecondParams{
			Brand: arg.Brand,
			Link: arg.Link,
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
		if err != nil{
			return err
		}
		return err
	})
	fmt.Println(result)

	return result, err
}

func (store *SQLStore) LengthOfSecond(ctx context.Context) (int64, error) {
	var result int64

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result, err = q.GetLengthOfSecond(ctx)
		if err != nil{
			return err
		}
		return err
	})
	fmt.Println(result)

	return result, err
}
