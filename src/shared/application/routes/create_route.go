package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goioc/di"
	"goapi/src/shared/domain/contracts"
)

func CreateRoute(app *fiber.App, method string, path string, action string) {
	switch method {
	case "GET":
		{
			app.Get(path, func(c *fiber.Ctx) error {
				return di.GetInstance(action).(contracts.Action).Handle(c)
			})
		}
	case "POST":
		{
			app.Post(path, func(c *fiber.Ctx) error {
				return di.GetInstance(action).(contracts.Action).Handle(c)
			})
		}
	}
}
