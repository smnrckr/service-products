package handlers

import (
	"service-products/internals/models"

	"github.com/gofiber/fiber/v2"
)

type ProductService interface {
	GetProducts() ([]models.Product, error)
	GetProductById(productId int) (models.Product, error)
}

type ProductHander struct {
	productService ProductService
}

func NewProductHandler(productService ProductService) *ProductHander {
	return &ProductHander{
		productService: productService,
	}
}

func (h *ProductHander) handleGetProducts(c *fiber.Ctx) error {

	products, err := h.productService.GetProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(products)
}

func (h *ProductHander) handleGetProductById(c *fiber.Ctx) error {
	productId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}

	product, err := h.productService.GetProductById(productId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(product)
}

func (h *ProductHander) SetRoutes(app *fiber.App) {

	productsGroup := app.Group("/products")
	productsGroup.Get("/", h.handleGetProducts)
	productsGroup.Get("/:id<int>", h.handleGetProductById)

}
