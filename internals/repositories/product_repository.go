package repositories

import (
	"service-products/internals/models"

	"gorm.io/gorm"
)

type PostgreDB interface {
	GetConnection() *gorm.DB
	Close()
}

type ProductRepository struct {
	storage PostgreDB
}

func NewProductsRepository(s PostgreDB) *ProductRepository {
	return &ProductRepository{
		storage: s,
	}
}

func (r *ProductRepository) GetProducts() ([]models.Product, error) {
	products := []models.Product{}
	err := r.storage.GetConnection().Find(&products).Error
	if err != nil {
		return []models.Product{}, err
	}
	return products, nil
}

func (r *ProductRepository) GetProductById(productId int) (models.Product, error) {
	product := models.Product{}
	err := r.storage.GetConnection().Where("id=? ", productId).Find(&product).Error
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}
