package main

import (
	"fmt"
	"os"
	"service-products/internals/handlers"
	"service-products/internals/repositories"
	"service-products/internals/services"
	"service-products/pkg/postgresql"
	"service-products/pkg/redis"
	"service-products/utils"
	"strconv"

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

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDBstr := os.Getenv("REDIS_DB")
	redisDB, err := strconv.Atoi(redisDBstr)
	if err != nil {
		fmt.Println(err)
	}
	rdb := redis.NewClient(redis.RedisConfig{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: redisPassword,
		Db:       redisDB,
	})

	productRepository := repositories.NewProductsRepository(db)
	productService := services.NewProductService(productRepository, rdb)
	productHander := handlers.NewProductHandler(productService)

	app := fiber.New()
	productHander.SetRoutes(app)

	sw := swagno.New(swagno.Config{Title: "Service Favorites", Version: "v1.0.0", Host: "localhost:8082"})
	sw.AddEndpoints(handlers.UserEndpoints)
	swagger.SwaggerHandler(app, sw.MustToJson(), swagger.WithPrefix("/swagger"))

	app.Listen(":8082")
}
