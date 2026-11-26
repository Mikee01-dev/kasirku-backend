package main

import (
	"context"
	"log"

	"github.com/Mikee01-dev/kasirku-backend/app/repository"
	"github.com/Mikee01-dev/kasirku-backend/app/service"
	"github.com/Mikee01-dev/kasirku-backend/config"
	"github.com/Mikee01-dev/kasirku-backend/database"
	"github.com/Mikee01-dev/kasirku-backend/route"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	ctx := context.Background()

	dsn := "postgres://" + config.GetEnv("DB_USER", "postgres") + ":" + config.GetEnv("DB_PASSWORD", "") + "@" + config.GetEnv("DB_HOST", "localhost") + ":" + config.GetEnv("DB_PORT", "5432") + "/" + config.GetEnv("DB_NAME", "kasirku_db") + "?sslmode=" + config.GetEnv("DB_SSLMODE", "disable")

	dbPool, err := database.NewPostgres(ctx, dsn)
	if err != nil {
		log.Fatal("Gagal koneksi ke DB: ", err)
	}
	defer dbPool.Close()

	productRepo := repository.NewProductRepository(dbPool)
	productService := service.NewProductService(productRepo)
	productRoute := route.NewProductRoute(productService)

	app := fiber.New()
	
	api := app.Group("/api/v1")
	productRoute.RegisterRoutes(api)

	port := config.GetEnv("PORT", "3000")
	log.Printf("Server berjalan di http://localhost:%s", port)
	log.Fatal(app.Listen(":" + port))
}