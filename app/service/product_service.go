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

func (s *productService) validate(p *model.Product) error {
	if p.SKU == "" || p.Name == "" || p.Price < 0 || p.Stock < 0 {
		return model.ErrInvalidProductInput
	}
	return nil
}