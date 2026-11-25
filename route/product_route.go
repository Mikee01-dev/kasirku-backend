package route

import "github.com/gofiber/fiber/v2"

func ProductRoutes(app *fiber.App) {
	product := app.Group("/products")

	product.Get("/", getProducts)
	product.Get("/:id", getProductByID)
	product.Post("/", createProduct)
	product.Put("/:id", updateProduct)
	product.Delete("/:id", deleteProduct)
}

func getProducts(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "get all products",
	})
}

func getProductByID(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "get product by id",
		"id":      c.Params("id"),
	})
}

func createProduct(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "create product",
	})
}

func updateProduct(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "update product",
		"id":      c.Params("id"),
	})
}

func deleteProduct(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "delete product",
		"id":      c.Params("id"),
	})
}
