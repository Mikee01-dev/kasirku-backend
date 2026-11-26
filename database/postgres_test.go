package database

import (
	"context"
	"testing"
)

func TestNewPostgresRequiresDSN(t *testing.T) {
	pool, err := NewPostgres(context.Background(), " \t ")
	if err == nil {
		t.Fatal("NewPostgres() error = nil, want missing DSN error")
	}
	if pool != nil {
		t.Fatal("NewPostgres() pool != nil for missing DSN")
	}
}
