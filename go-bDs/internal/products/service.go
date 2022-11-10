package products

import (
	"context"

	"github.com/bootcamp-go/go-bDs/internal/domain"
)

type Service interface {
	GetByName(ctx context.Context, name string) (domain.Product, error)
	GetAll(ctx context.Context) ([]domain.Product, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetByName(ctx context.Context, name string) (domain.Product, error) {
	product, err := s.repository.GetByName(ctx, name)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (s *service) GetAll(ctx context.Context) ([]domain.Product, error) {
	products, err := s.repository.GetAll(ctx)
	if err != nil {
		return []domain.Product{}, err
	}

	return products, err
}

func (s *service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}
