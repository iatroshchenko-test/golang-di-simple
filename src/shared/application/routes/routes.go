package routes

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
	"goapi/src/mymodule/application/use_case/list_users"
)

func CreateRoutes(app *fiber.App, container *dig.Container) {
	app.Get("/users", func(c *fiber.Ctx) error {
		err := container.Invoke(func(action *list_users.ListUsers) error {
			return action.Handle(c)
		})
		return err
	})
}
