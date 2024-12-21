package list_posts

import (
	"github.com/gofiber/fiber/v2"
	"goapi/src/mymodule/domain/entity/post"
	"goapi/src/shared/domain/contracts"
	"gorm.io/gorm"
)

type ListPosts struct {
	DB *gorm.DB `di.inject:"db"`
}

func NewListPosts(db *gorm.DB) contracts.Action {
	return &ListPosts{DB: db}
}

func (context *ListPosts) Handle(c *fiber.Ctx) error {
	var posts []post.Post
	if err := context.DB.Find(&posts).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch posts",
		})
	}
	return c.JSON(posts)
}
