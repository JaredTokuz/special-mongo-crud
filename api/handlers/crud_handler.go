package handlers

import (
	"errors"
	"net/http"

	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jaredtokuz/mongo-crud/api/presenter"
	"github.com/jaredtokuz/mongo-crud/pkg/crud"
)


func FindOne(service crud.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
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
		fieldName := c.Params("fieldName")
		days := c.Params("days")
		daysi, _ := strconv.Atoi(days)
		collectionName := c.Query("collection")
		if collectionName == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ErrorResponse(errors.New("you forgot query string: collection"))) 
		}
		result, err := service.DaysOutExpire(collectionName, fieldName, daysi)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ErrorResponse(err)) 
		}
		return c.JSON(presenter.SuccessResponse(result)) 
	}
}

func Shave(service crud.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		count := c.Params("count")
		counti, _ := strconv.ParseInt(count, 10, 64)
		collectionName := c.Query("collection")
		if collectionName == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ErrorResponse(errors.New("you forgot query string: collection"))) 
		}
		result, err := service.DeleteFirstNRows(collectionName, counti)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ErrorResponse(err)) 
		}
		return c.JSON(presenter.SuccessResponse(result)) 
	}
}

func DeleteOne(service crud.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		result, err := service.DeleteOne(id) 
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ErrorResponse(err)) 
		}
		return c.JSON(presenter.SuccessResponse(result)) 
	}
}

func DeleteMany(service crud.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		result, err := service.Delete(id) 
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ErrorResponse(err)) 
		}
		return c.JSON(presenter.SuccessResponse(result)) 
	}
}

func UpdateOne(service crud.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		result, err := service.UpdateOne(id) 
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ErrorResponse(err)) 
		}
		return c.JSON(presenter.SuccessResponse(result)) 
	}
}

func UpdateMany(service crud.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		result, err := service.Update(id) 
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ErrorResponse(err)) 
		}
		return c.JSON(presenter.SuccessResponse(result)) 
	}
}

func InsertOne(service crud.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		result, err := service.InsertOne(id) 
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ErrorResponse(err)) 
		}
		return c.JSON(presenter.SuccessResponse(result)) 
	}
}

func InsertMany(service crud.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		result, err := service.Insert(id) 
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ErrorResponse(err)) 
		}
		return c.JSON(presenter.SuccessResponse(result)) 
	}
}
