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

// -- INSERT INTO products (product_name, product_code, product_price) VALUES
// -- ('Apple iPhone 14', 'A001', '1299.00'),
// -- ('Samsung Galaxy S23', 'S002', '999.00'),
// -- ('Sony WH-1000XM5 Headphones', 'S003', '349.99'),
// -- ('Dell XPS 13', 'D004', '1199.00'),
// -- ('Apple MacBook Pro 16"', 'A005', '2399.00'),
// -- ('Nintendo Switch OLED', 'N006', '349.99'),
// -- ('Sony PlayStation 5', 'S007', '499.00'),
// -- ('Microsoft Xbox Series X', 'M008', '499.00'),
// -- ('Apple iPad Pro 12.9"', 'A009', '1099.00'),
// -- ('Samsung Galaxy Tab S8', 'S010', '799.00'),
// -- ('Bose QuietComfort 45', 'B011', '329.00'),
// -- ('Logitech MX Master 3 Mouse', 'L012', '99.99'),
// -- ('Canon EOS 90D DSLR Camera', 'C013', '1199.00'),
// -- ('Nikon Z9 Mirrorless Camera', 'N014', '5499.00'),
// -- ('GoPro HERO10 Black', 'G015', '499.00'),
// -- ('Fitbit Charge 5', 'F016', '149.95'),
// -- ('Apple Watch Series 8', 'A017', '399.00'),
// -- ('Microsoft Surface Laptop 4', 'M018', '1299.00'),
// -- ('JBL Charge 5 Bluetooth Speaker', 'J019', '179.95'),
// -- ('Samsung 65" QLED TV', 'S020', '1299.00');
