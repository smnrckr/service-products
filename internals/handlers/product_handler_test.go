package handlers_test

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"service-products/internals/handlers"
	"service-products/internals/models"
	"service-products/internals/repositories"
	"service-products/internals/services"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestFavoritesListsRepository(t *testing.T) {

	userRepo := repositories.NewProductsRepository(testDb)
	userService := services.NewProductService(userRepo)
	handler := handlers.NewProductHandler(userService)

	app := fiber.New()

	handler.SetRoutes(app)

	t.Run("Get All Products", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/products", nil)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)

		jsonDataFromHttp, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		products := []models.Product{}
		err = json.Unmarshal(jsonDataFromHttp, &products)
		assert.NoError(t, err)
		assert.NotEmpty(t, products)

	})

	t.Run("Get Product By Id", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/products/1", nil)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)

		jsonDataFromHttp, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		product := models.Product{}
		err = json.Unmarshal(jsonDataFromHttp, &product)
		assert.NoError(t, err)

		assert.NotEmpty(t, "A001", product.ProductCode)

	})
}
