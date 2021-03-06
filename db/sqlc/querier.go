// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
)

type Querier interface {
	CreateFirstProduct(ctx context.Context, arg CreateFirstProductParams) (First, error)
	CreateOnSale(ctx context.Context, arg CreateOnSaleParams) (OnSale, error)
	DeleteFirstProduct(ctx context.Context, id int64) error
	DeleteOnSale(ctx context.Context, id int64) error
	GetFirstProduct(ctx context.Context, id int64) (First, error)
	GetFirstProductForUpdate(ctx context.Context, id int64) (First, error)
	GetLengthOfFirst(ctx context.Context) (int64, error)
	GetLengthOnSale(ctx context.Context) (int64, error)
	GetOnSale(ctx context.Context, id int64) (OnSale, error)
	GetOnSaleForUpdate(ctx context.Context, id int64) (OnSale, error)
	ListFirstProduct(ctx context.Context, arg ListFirstProductParams) ([]First, error)
	ListOnSale(ctx context.Context, arg ListOnSaleParams) ([]OnSale, error)
	UpdateFirstProduct(ctx context.Context, arg UpdateFirstProductParams) (First, error)
	UpdateOnSale(ctx context.Context, arg UpdateOnSaleParams) (OnSale, error)
}

var _ Querier = (*Queries)(nil)
