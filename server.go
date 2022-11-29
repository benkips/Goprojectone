package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/repentance/rest_vid_api/routes"
	"log"
	"os"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	// Define file to logs
	file, err := os.OpenFile("./my_logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	// Set config for logger
	loggerConfig := logger.Config{
		Output: file, // add file to save output
	}

	app.Use(logger.New(loggerConfig))

	// then we must define our routes function
	routes.Routes(app)
	// Listen on PORT 300
	app.Listen(":3000")
}
