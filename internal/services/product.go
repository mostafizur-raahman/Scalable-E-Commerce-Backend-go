package services

import "ecommerce/internal/domain/product"

type ProductService struct {
	reader  product.Reader
	getter  product.Getter
	creator product.Creator
	updater product.Updater
	deleter product.Deleter
}

func NewProductService(r product.Reader, g product.Getter, c product.Creator, u product.Updater, d product.Deleter) *ProductService {
	return &ProductService{reader: r, getter: g, creator: c, updater: u, deleter: d}
}

func (s *ProductService) GetAllProducts() ([]product.Product, error)    { return s.reader.GetAll() }
func (s *ProductService) GetProduct(id int64) (*product.Product, error) { return s.getter.GetByID(id) }
func (s *ProductService) CreateProduct(p *product.Product) error        { return s.creator.Create(p) }
func (s *ProductService) UpdateProduct(p *product.Product) error        { return s.updater.Update(p) }
func (s *ProductService) DeleteProduct(id int64) error                  { return s.deleter.Delete(id) }
