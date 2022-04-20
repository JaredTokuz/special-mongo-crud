package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jaredtokuz/mongo-crud/api/presenter"
	"github.com/jaredtokuz/mongo-crud/pkg/crud"
)

func FindOne(service crud.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		collection := c.Params("collection")
		id := c.Params("id")
		result, err := service.FindOne(collection, id) 
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ErrorResponse(err)) 
		}
		return c.JSON(presenter.SuccessResponse(result)) 
	}
}

func FindMany(service crud.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		collection := c.Params("collection")
		id := c.Params("id")
		result, err := service.Find(collection, id) 
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ErrorResponse(err)) 
		}
		return c.JSON(presenter.SuccessResponse(result)) 
	}
}

func Expire(service crud.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		collection := c.Params("collection")
		field := c.Params("field")
		days := c.Params("days")
		result, err := service.ExpireDocs(collection, field, days) 
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ErrorResponse(err)) 
		}
		return c.JSON(presenter.SuccessResponse(result)) 
	}
}