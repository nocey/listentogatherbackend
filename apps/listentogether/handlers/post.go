package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/listentogether/database"
	"github.com/listentogether/database/models"
)

type Post struct{}

func (p *Post) Create(c *fiber.Ctx) error {
	type NewPost struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		UserID  uint   `json:"user_id"`
	}

	var newPost NewPost
	var post models.Posts

	if err := c.BodyParser(&newPost); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}
	if newPost.Title == "" || newPost.Content == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Title, content, and author ID are required",
		})
	}

	post.Title = newPost.Title
	post.Content = newPost.Content
	user := c.Locals("user").(*models.Users)
	post.UserID = user.ID

	if err := database.DBConn.Create(&post).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create post",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(post)
}

func (p *Post) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Posts
	if err := database.DBConn.First(&post, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Post not found",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(post)
}

func (p *Post) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Posts
	if err := database.DBConn.First(&post, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Post not found",
			"error":   err.Error(),
		})
	}

	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	if err := database.DBConn.Save(&post).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update post",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(post)
}
