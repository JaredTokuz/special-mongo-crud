package routes

import (
	"github.com/jaredtokuz/mongo-crud/api/handlers"
	"github.com/jaredtokuz/mongo-crud/pkg/find"

	"github.com/gofiber/fiber/v2"
)

// BookRouter is the Router for GoFiber App
func FindRouter(app fiber.Router, service find.Service) {
	app.Get("/many", handlers.FindMany(service))
	app.Get("/one", handlers.FindOne(service))
}

func DeleteRouter(app fiber.Router, service delete.Service) {
	app.Delete("/expire", handlers.Expire(service)) // 
	app.Delete("/many", handlers.DeleteMany(service))
	app.Delete("/one", handlers.DeleteOne(service))
}

func UpdateRouter(app fiber.Router, service update.Service) {
	app.Put("/many", handlers.UpdateMany(service))
	app.Put("/one", handlers.UpdateOne(service))
}

func InsertRouter(app fiber.Router, service insert.Service) {
	app.Post("/many", handlers.InsertMany(service))
	app.Post("/one", handlers.InsertOne(service))
}
