package list_users

import (
	"github.com/gofiber/fiber/v2"
	"goapi/src/mymodule/domain/entity/user"
	"gorm.io/gorm"
)

type ListUsers struct {
	DB *gorm.DB `di.inject:""`
}

func NewListUsers(db *gorm.DB) *ListUsers {
	return &ListUsers{DB: db}
}

func (context *ListUsers) Handle(c *fiber.Ctx) error {
	var users []user.User
	if err := context.DB.Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch users",
		})
	}
	return c.JSON(users)
}
