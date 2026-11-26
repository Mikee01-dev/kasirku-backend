package service

import (
	"context"
	"github.com/Mikee01-dev/kasirku-backend/app/model"
)

type productService struct {
	repo model.ProductRepository
}

func NewProductService(repo model.ProductRepository) model.ProductService {
	return &productService{repo: repo}
}