package routes

import (
	"github.com/jaredtokuz/mongo-crud/api/handlers"
	"github.com/jaredtokuz/mongo-crud/pkg/crud"

	"github.com/gofiber/fiber/v2"
)

// BookRouter is the Router for GoFiber App
func FindRouter(app fiber.Router, service crud.Service) {
	app.Get("/many/:id", handlers.FindMany(service))
	app.Get("/one/:id", handlers.FindOne(service))
}

func DeleteRouter(app fiber.Router, service crud.Service) {
	app.Delete("/expire/:field/:days", handlers.Expire(service)) 
	app.Delete("/many/:id", handlers.DeleteMany(service))
	app.Delete("/one/:id", handlers.DeleteOne(service))
}

func UpdateRouter(app fiber.Router, service crud.Service) {
	app.Put("/many/rename/:current/:new", handlers.UpdateMany(service))
	app.Put("/many/unset/:field", handlers.UpdateMany(service))
	app.Put("/many/set/:field/:value", handlers.UpdateMany(service))
	app.Put("/one/rename/:current/:new", handlers.UpdateOne(service))
	app.Put("/one/unset/:field", handlers.UpdateOne(service))
	app.Put("/one/set/:field/:value", handlers.UpdateOne(service))
}

// ? some type of insert randomized/planned doc type function
// func InsertRouter(app fiber.Router, service crud.Service) {
// 	app.Post("/many", handlers.InsertMany(service))
// 	app.Post("/one", handlers.InsertOne(service))
// }
