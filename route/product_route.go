package route

import (
	"strconv"

	"github.com/Mikee01-dev/kasirku-backend/app/model"
	"github.com/Mikee01-dev/kasirku-backend/helper"
	"github.com/gofiber/fiber/v2"
)

type ProductRoute struct {
	service model.ProductService
}

func NewProductRoute(s model.ProductService) *ProductRoute {
	return &ProductRoute{service: s}
}

func (r *ProductRoute) RegisterRoutes(router fiber.Router) {
	products := router.Group("/products")
	products.Get("/", r.getAll)
	products.Get("/:id", r.getByID)
	products.Post("/", r.create)
	products.Put("/:id", r.update)
	products.Delete("/:id", r.delete)
}

func (r *ProductRoute) getAll(c *fiber.Ctx) error {
	products, err := r.service.GetAll(c.Context(), model.ListParams{Limit: 20})
	if err != nil {
		return helper.Fail(c, 500, "Terjadi kesalahan server")
	}
	return helper.Success(c, 200, "Berhasil ambil produk", products)
}

func (r *ProductRoute) getByID(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	product, err := r.service.GetByID(c.Context(), id)
	if err != nil {
		return helper.Fail(c, 404, "Produk tidak ditemukan")
	}
	return helper.Success(c, 200, "Berhasil ambil produk", product)
}

func (r *ProductRoute) create(c *fiber.Ctx) error {
	var p model.Product
	if err := c.BodyParser(&p); err != nil {
		return helper.Fail(c, 400, "Format request salah")
	}
	if err := r.service.Create(c.Context(), &p); err != nil {
		return helper.Fail(c, 400, err.Error())
	}
	return helper.Success(c, 201, "Produk berhasil dibuat", p)
}

func (r *ProductRoute) update(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	var p model.Product
	if err := c.BodyParser(&p); err != nil {
		return helper.Fail(c, 400, "Format request salah")
	}
	p.ID = id
	if err := r.service.Update(c.Context(), &p); err != nil {
		return helper.Fail(c, 400, err.Error())
	}
	return helper.Success(c, 200, "Produk berhasil diupdate", p)
}

func (r *ProductRoute) delete(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	if err := r.service.Delete(c.Context(), id); err != nil {
		return helper.Fail(c, 404, "Produk tidak ditemukan")
	}
	return helper.Success(c, 200, "Produk berhasil dihapus", nil)
}