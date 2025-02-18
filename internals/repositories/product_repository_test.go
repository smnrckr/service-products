package repositories_test

import (
	"service-products/internals/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFavoritesListsRepository(t *testing.T) {

	userRepo := repositories.NewProductsRepository(testDb)

	t.Run("Get All Products", func(t *testing.T) {
		products, err := userRepo.GetProducts()
		assert.NoError(t, err)
		assert.NotEmpty(t, products)

	})

	t.Run("Get Product By Id", func(t *testing.T) {
		product, err := userRepo.GetProductById(1)
		assert.NoError(t, err)
		assert.Equal(t, 1, product.Id)

	})
}
