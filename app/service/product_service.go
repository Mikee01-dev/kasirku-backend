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

func (s *productService) GetAll(ctx context.Context, params model.ListParams) ([]model.Product, error) {
	if params.Limit <= 0 || params.Limit > 100 {
		params.Limit = 20
	}
	return s.repo.FindAll(ctx, params)
}

func (s *productService) GetByID(ctx context.Context, id int64) (*model.Product, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *productService) Create(ctx context.Context, p *model.Product) error {
	if err := s.validate(p); err != nil {
		return err
	}
	return s.repo.Create(ctx, p)
}

func (s *productService) Update(ctx context.Context, p *model.Product) error {
	if err := s.validate(p); err != nil {
		return err
	}
	return s.repo.Update(ctx, p)
}

func (s *productService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}