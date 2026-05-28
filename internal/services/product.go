package services

import (
	"ecommerce/internal/domain"
)

type Service struct {
	reader  domain.Reader
	getter  domain.Getter
	creator domain.Creator
}

func NewProductService(reader domain.Reader, getter domain.Getter, creator domain.Creator) *Service {
	return &Service{
		reader:  reader,
		getter:  getter,
		creator: creator,
	}
}

func (s *Service) GetAllProducts() ([]domain.Product, error) {
	return s.reader.GetAll()
}

func (s *Service) GetProduct(id int64) (*domain.Product, error) {
	return s.getter.GetByID(id)
}

func (s *Service) CreateProduct(product *domain.Product) error {
	return s.creator.Create(product)
}
