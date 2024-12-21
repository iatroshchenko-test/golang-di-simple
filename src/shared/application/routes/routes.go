package routes

import (
	"github.com/gofiber/fiber/v2"
)

func CreateRoutes(app *fiber.App) {
	CreateRoute(app, "GET", "/users", "listUsers")
	CreateRoute(app, "GET", "/posts", "listPosts")
}
