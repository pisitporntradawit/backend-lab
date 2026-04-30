package products

import (
    "context"
    "fmt"
)

// Interface ให้ test ได้โดยไม่ต้องมี DB จริง
type ProductRepository interface {
    GetProducts(ctx context.Context) ([]ProductsModel, error)
}

type Service struct {
    repo ProductRepository
}

func NewService(repo ProductRepository) *Service {
    return &Service{repo: repo}
}

func (s *Service) GetProducts(ctx context.Context) ([]ProductsModel, error) {
    products, err := s.repo.GetProducts(ctx)
    if err != nil {
        return nil, fmt.Errorf("service.GetProducts: %w", err)
    }
    return products, nil
}

