// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"time"
)

type First struct {
	ID        int64     `json:"id"`
	Brand     string    `json:"brand"`
	Link      string    `json:"link"`
	Price     string    `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

type OnSale struct {
	ID        int64     `json:"id"`
	Brand     string    `json:"brand"`
	Link      string    `json:"link"`
	Price     string    `json:"price"`
	Saleper   int64     `json:"saleper"`
	CreatedAt time.Time `json:"created_at"`
}

type Second struct {
	ID        int64     `json:"id"`
	Brand     string    `json:"brand"`
	Link      string    `json:"link"`
	Price     string    `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
