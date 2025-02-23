package services

import (
	"encoding/json"
	"fmt"
	"service-products/internals/models"
	"time"
)

type ProductRepository interface {
	GetProducts() ([]models.Product, error)
	GetProductById(productId int) (models.Product, error)
}

type RedisDB interface {
	Get(key string) (string, error)
	Set(key string, value interface{}, expiration time.Duration) error
}

type ProductService struct {
	productRepository ProductRepository
	redisDB           RedisDB
}

func NewProductService(productRepository ProductRepository, redisdb RedisDB) *ProductService {
	return &ProductService{
		productRepository: productRepository,
		redisDB:           redisdb,
	}
}

func (s *ProductService) GetProducts() ([]models.Product, error) {
	products, err := s.getProductsFromCache()
	if err == nil {
		return products, nil
	}
	fmt.Println("Cacheden veri gelmedi")

	products, err = s.productRepository.GetProducts()
	fmt.Println("Products dbden getirildi")
	if err != nil {
		return nil, err
	}

	err = s.saveProductsToCache(products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *ProductService) GetProductById(productId int) (models.Product, error) {
	redisResult, err := s.redisDB.Get("products")
	var products []models.Product
	if err == nil && redisResult != "" {
		if err := json.Unmarshal([]byte(redisResult), &products); err == nil {
			for _, product := range products {
				if product.Id == productId {
					return product, nil
				}
			}
		}
	}

	product, err := s.productRepository.GetProductById(productId)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (s *ProductService) getProductsFromCache() ([]models.Product, error) {
	redisResult, err := s.redisDB.Get("products")
	if err != nil || redisResult == "" {
		return nil, err
	}

	products := []models.Product{}
	if err := json.Unmarshal([]byte(redisResult), &products); err != nil {
		return nil, err
	}

	return products, nil
}

func (s *ProductService) saveProductsToCache(products []models.Product) error {
	data, err := json.Marshal(products)
	if err != nil {
		return err
	}

	return s.redisDB.Set("products", data, time.Minute*5)
}
