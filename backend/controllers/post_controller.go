package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbalmahad/post-app.git/backend/models"
	"github.com/iqbalmahad/post-app.git/backend/repositories"
)

type PostController struct {
	repository repositories.PostRepository
}

func NewPostController(repository repositories.PostRepository) *PostController {
	return &PostController{repository}
}

func (controller *PostController) CreatePost(c *fiber.Ctx) error {
	var post models.Post
	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := controller.repository.Create(&post); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(post)
}

func (controller *PostController) GetPosts(c *fiber.Ctx) error {
	posts, err := controller.repository.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(posts)
}

func (controller *PostController) GetPost(c *fiber.Ctx) error {
	id := c.Params("id")
	post, err := controller.repository.FindByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Post not found"})
	}

	return c.JSON(post)
}

func (controller *PostController) UpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Post
	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	post.ID = id
	if err := controller.repository.Update(&post); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(post)
}

func (controller *PostController) DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := controller.repository.Delete(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
