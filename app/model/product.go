package model

import (
	"context"
	"errors"
	"time"
)

type Product struct {
	ID        int64     `json:"id"`
	SKU       string    `json:"sku"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Price     float64   `json:"price"`
	Stock     int       `json:"stock"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ListParams struct {
	Limit    int
	AfterID  int64
	Category string
}

var (
	ErrProductNotFound     = errors.New("product not found")
	ErrInvalidProductInput = errors.New("invalid product input")
	ErrDuplicateSKU        = errors.New("sku already exists")
	ErrInsufficientStock   = errors.New("insufficient stock")
)

type ProductRepository interface {
	FindAll(ctx context.Context, params ListParams) ([]Product, error)
	FindByID(ctx context.Context, id int64) (*Product, error)
	Create(ctx context.Context, product *Product) error
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id int64) error
}

type ProductService interface {
	GetAll(ctx context.Context, params ListParams) ([]Product, error)
	GetByID(ctx context.Context, id int64) (*Product, error)
	Create(ctx context.Context, product *Product) error
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id int64) error
}