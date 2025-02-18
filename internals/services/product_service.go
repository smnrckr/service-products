package services

import "service-products/internals/models"

type ProductRepository interface {
	GetProducts() ([]models.Product, error)
	GetProductById(productId int) (models.Product, error)
}

type ProductService struct {
	productRepository ProductRepository
}

func NewProductService(productRepository ProductRepository) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (s *ProductService) GetProducts() ([]models.Product, error) {
	return s.productRepository.GetProducts()
}

func (s *ProductService) GetProductById(productId int) (models.Product, error) {
	return s.productRepository.GetProductById(productId)
}
