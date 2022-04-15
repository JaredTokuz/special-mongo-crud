package routes

import (
	"github.com/jaredtokuz/mongo-crud/api/handlers"
	"github.com/jaredtokuz/mongo-crud/pkg/find"

	"github.com/gofiber/fiber/v2"
)

// BookRouter is the Router for GoFiber App
func BookRouter(app fiber.Router, service find.Service) {
	app.Get("/find", handlers.GetBooks(service))
	app.Post("/find", handlers.AddBook(service))
	app.Put("/find", handlers.UpdateBook(service))
	app.Delete("/find", handlers.RemoveBook(service))
}