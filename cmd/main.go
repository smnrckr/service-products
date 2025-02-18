package main

import (
	"os"
	"service-products/internals/handlers"
	"service-products/internals/repositories"
	"service-products/internals/services"
	"service-products/pkg/postgresql"
	"service-products/utils"

	"github.com/go-swagno/swagno"
	"github.com/go-swagno/swagno-fiber/swagger"
	"github.com/gofiber/fiber/v2"
)

func init() {
	utils.LoadEnviromentVariables()
}

func main() {
	host := os.Getenv("HOST")
	dbuser := os.Getenv("USER_NAME")
	dbname := os.Getenv("DB_NAME")
	dbpassword := os.Getenv("PASSWORD")
	port := os.Getenv("PORT")

	db := postgresql.NewDB(postgresql.DbConfig{Host: host, Dbuser: dbuser, Dbname: dbname, Dbpassword: dbpassword, Port: port})

	productRepository := repositories.NewProductsRepository(db)
	productService := services.NewProductService(productRepository)
	productHander := handlers.NewProductHandler(productService)

	app := fiber.New()
	productHander.SetRoutes(app)

	sw := swagno.New(swagno.Config{Title: "Service Favorites", Version: "v1.0.0", Host: "localhost:8082"})
	sw.AddEndpoints(handlers.UserEndpoints)
	swagger.SwaggerHandler(app, sw.MustToJson(), swagger.WithPrefix("/swagger"))

	app.Listen(":8082")
}
