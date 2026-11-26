package repository

import (
	"errors"
	"testing"

	"github.com/Mikee01-dev/kasirku-backend/app/model"
	"github.com/jackc/pgx/v5/pgconn"
)

func TestMapProductRepositoryError(t *testing.T) {
	databaseErr := errors.New("database unavailable")
	primaryKeyErr := &pgconn.PgError{Code: "23505", ConstraintName: "products_pkey"}

	tests := []struct {
		name string
		err  error
		want error
	}{
		{name: "nil error", err: nil, want: nil},
		{name: "duplicate SKU", err: &pgconn.PgError{Code: "23505", ConstraintName: "products_sku_key"}, want: model.ErrDuplicateSKU},
		{name: "other unique constraint", err: primaryKeyErr, want: primaryKeyErr},
		{name: "other database error", err: databaseErr, want: databaseErr},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mapProductRepositoryError(tt.err)
			if (got == nil) != (tt.want == nil) {
				t.Fatalf("mapProductRepositoryError() = %v, want %v", got, tt.want)
			}
			if tt.want != nil && !errors.Is(got, tt.want) {
				t.Fatalf("mapProductRepositoryError() = %v, want %v", got, tt.want)
			}
		})
	}
}
