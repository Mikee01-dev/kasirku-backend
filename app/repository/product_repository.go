package repository

import (
	"context"
	"errors"

	"github.com/Mikee01-dev/kasirku-backend/app/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type productRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) model.ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) FindAll(ctx context.Context, params model.ListParams) ([]model.Product, error) {
	const query = `
		SELECT id, sku, name, category, price, stock, is_active, created_at, updated_at
		FROM products
		WHERE id > $1 AND ($2 = '' OR category = $2)
		ORDER BY id
		LIMIT $3`

	rows, err := r.db.Query(ctx, query, params.AfterID, params.Category, params.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]model.Product, 0)
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(
			&product.ID,
			&product.SKU,
			&product.Name,
			&product.Category,
			&product.Price,
			&product.Stock,
			&product.IsActive,
			&product.CreatedAt,
			&product.UpdatedAt,
		); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepository) FindByID(ctx context.Context, id int64) (*model.Product, error) {
	const query = `
		SELECT id, sku, name, category, price, stock, is_active, created_at, updated_at
		FROM products
		WHERE id = $1`

	var product model.Product
	err := r.db.QueryRow(ctx, query, id).Scan(
		&product.ID,
		&product.SKU,
		&product.Name,
		&product.Category,
		&product.Price,
		&product.Stock,
		&product.IsActive,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, model.ErrProductNotFound
		}
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) Create(ctx context.Context, product *model.Product) error {
	const query = `
		INSERT INTO products (sku, name, category, price, stock, is_active)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at`

	err := r.db.QueryRow(
		ctx, query,
		product.SKU, product.Name, product.Category, product.Price, product.Stock, product.IsActive,
	).Scan(&product.ID, &product.CreatedAt, &product.UpdatedAt)
	return mapProductRepositoryError(err)
}

func (r *productRepository) Update(ctx context.Context, product *model.Product) error {
	const query = `
		UPDATE products
		SET sku = $1, name = $2, category = $3, price = $4, stock = $5,
			is_active = $6, updated_at = NOW()
		WHERE id = $7`

	tag, err := r.db.Exec(
		ctx, query,
		product.SKU, product.Name, product.Category, product.Price, product.Stock, product.IsActive, product.ID,
	)
	if err != nil {
		return mapProductRepositoryError(err)
	}
	if tag.RowsAffected() == 0 {
		return model.ErrProductNotFound
	}

	return nil
}

func (r *productRepository) Delete(ctx context.Context, id int64) error {
	tag, err := r.db.Exec(ctx, `DELETE FROM products WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return model.ErrProductNotFound
	}

	return nil
}

func mapProductRepositoryError(err error) error {
	if err == nil {
		return nil
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" && pgErr.ConstraintName == "products_sku_key" {
		return model.ErrDuplicateSKU
	}
	return err
}
