package handlers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jaredtokuz/mongo-crud/api/presenter"
	"github.com/jaredtokuz/mongo-crud/pkg/crud"
)


func FindOne(service crud.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ErrorResponse(errors.New("path missing id"))) 
		}
		result, err := service.FindOne(id) 
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ErrorResponse(err)) 
		}
		return c.JSON(presenter.SuccessResponse(result)) 
	}
}

func FindMany(service crud.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ErrorResponse(errors.New("path missing id"))) 
		}
		result, err := service.Find(id) 
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ErrorResponse(err)) 
		}
		return c.JSON(presenter.SuccessResponse(result)) 
	}
}

func Expire(service crud.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

	}
}