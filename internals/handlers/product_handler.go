package handlers

import (
	"service-products/internals/models"

	"github.com/go-swagno/swagno/components/endpoint"
	"github.com/go-swagno/swagno/components/http/response"
	"github.com/go-swagno/swagno/components/parameter"
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

var UserEndpoints = []*endpoint.EndPoint{
	endpoint.New(
		endpoint.GET,
		"/products",
		endpoint.WithTags("product"),
		endpoint.WithSuccessfulReturns([]response.Response{response.New(
			[]models.Product{}, "OK", "200")}),
		endpoint.WithErrors([]response.Response{response.New(models.ErrorResponse{}, "Bad Request", "400")}),
		endpoint.WithDescription("tüm ürünleri döner"),
	),
	endpoint.New(
		endpoint.GET,
		"/products/{id}",
		endpoint.WithTags("product"),
		endpoint.WithParams(parameter.IntParam("id", parameter.Path, parameter.WithRequired(), parameter.WithDescription("Ürün Id"))),
		endpoint.WithSuccessfulReturns([]response.Response{response.New(
			models.Product{}, "OK", "200")}),
		endpoint.WithErrors([]response.Response{response.New(models.ErrorResponse{}, "Bad Request", "400")}),
		endpoint.WithDescription("id'e göre ürün döner"),
	),
}
