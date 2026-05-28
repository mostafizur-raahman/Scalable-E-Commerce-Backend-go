package database

import (
	"ecommerce/internal/domain"
	"errors"
	"sync"
)

type Repository struct {
	products map[int64]*domain.Product
	mu       sync.RWMutex
	nextId   int64
}

func NewProductRepository() *Repository {
	return &Repository{
		products: make(map[int64]*domain.Product),
		nextId:   1,
	}
}

func (r *Repository) Create(product *domain.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if product.Id == 0 {
		product.Id = r.nextId
		r.nextId++
	} else if _, exists := r.products[product.Id]; exists {
		return errors.New("product with this ID already exists")
	}

	newProduct := *product
	r.products[newProduct.Id] = &newProduct
	return nil
}

func (r *Repository) GetAll() ([]domain.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	products := make([]domain.Product, 0, len(r.products))
	for _, p := range r.products {
		products = append(products, *p)
	}
	return products, nil
}

func (r *Repository) GetByID(id int64) (*domain.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	product, exists := r.products[id]
	if !exists {
		return nil, errors.New("product not found")
	}

	newProduct := *product
	return &newProduct, nil
}
