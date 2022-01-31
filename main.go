package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func HelloWorld(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "Hello World! Azure Functions!",
		"time":    time.Now(),
	})
}

func main() {
	app := fiber.New()

	api := app.Group("/api")
	api.Get("/hello-world", HelloWorld)

	FunctionPort, exists := os.LookupEnv("FUNCTIONS_HTTPWORKER_PORT")
	if exists {
		fmt.Println("FUNCTIONS_HTTPWORKER_PORT: " + FunctionPort)
	}

	log.Println("Azure Functions Server Running...!")
	if err := app.Listen(":" + FunctionPort); err != nil {
		log.Fatalln("Failed to Fiber Listen Server")
	}

}
