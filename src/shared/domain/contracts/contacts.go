package contracts

import "github.com/gofiber/fiber/v2"

type Action interface {
	Handle(c *fiber.Ctx) error
}
