package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/repentance/rest_vid_api/handlers"
)

func Routes(app *fiber.App) {
	// method GET
	app.Get("/", handlers.ReadData)

}
