package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jaredtokuz/mongo-crud/api/routes"
	"github.com/jaredtokuz/mongo-crud/pkg/crud"
	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load()
  if err != nil {
    log.Println("Error loading .env file")
  }
	db, cancel, err := databaseConnection()
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	fmt.Println("Database connection success!")
	fmt.Println(db)
	
	crudRepo := crud.NewRepo(db)
	crudService := crud.NewService(crudRepo)

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the clean-architecture mongo book shop!"))
	})
	find := app.Group("/find")
	routes.FindRouter(find, crudService)
	del := app.Group("/delete")
	routes.DeleteRouter(del, crudService)
	ins := app.Group("/insert")
	routes.InsertRouter(ins, crudService)
	updt := app.Group("/update")
	routes.UpdateRouter(updt, crudService)

	defer cancel()
	log.Fatal(app.Listen(":8080"))
}

func databaseConnection() (*mongo.Database, context.CancelFunc, error) {
	mongoURI := os.Getenv("MONGO_URI")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		mongoURI).SetServerSelectionTimeout(5*time.
		Second))
	if err != nil {
		cancel()
		return nil, nil, err
	}
	db := client.Database(os.Getenv("DB_NAME"))
	return db, cancel, nil
}