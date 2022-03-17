package handler

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/wonpanu/personal-blog/service/pkg/entity"
	"github.com/wonpanu/personal-blog/service/pkg/usecase"
)

type BlogHandler struct {
	uc usecase.IBlogUsecase
}

func (r BlogHandler) Create(c *fiber.Ctx) error {
	var body entity.Blog
	err := c.BodyParser(&body)
	if err != nil {
		log.Println(err, body)
		return c.Status(fiber.ErrBadRequest.Code).SendString("Invalid Payload!")
	}
	reponse, err := r.uc.Create(body)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.ErrBadRequest.Code).SendString("Fail to create a blog.")
	}
	return c.Status(fiber.StatusOK).JSON(reponse)
}

func (r BlogHandler) GetAll(c *fiber.Ctx) error {
	reponse, err := r.uc.GetAll()
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Fail to get all blogs.")
	}
	return c.Status(fiber.StatusOK).JSON(reponse)
}

func (r BlogHandler) GetByBlogID(c *fiber.Ctx) error {
	ID := c.Params("id")
	response, err := r.uc.GetByBlogID(ID)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Fail to get a blog by id.")
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (r BlogHandler) UpdateByBlogID(c *fiber.Ctx) error {
	ID := c.Params("id")
	var body entity.Blog
	err := c.BodyParser(&body)
	if err != nil {
		log.Println(err, body)
		return c.Status(fiber.ErrBadRequest.Code).SendString("Invalid Payload!")
	}
	response, err := r.uc.UpdateByBlogID(ID, body)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.ErrBadRequest.Code).SendString("Fail to update a blog by id.")
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (r BlogHandler) DeleteByBlogID(c *fiber.Ctx) error {
	ID := c.Params("id")
	err := r.uc.DeleteByBlogID(ID)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.ErrBadRequest.Code).SendString("Fail to delete a blog by id.")
	}
	return c.Status(fiber.StatusOK).SendString("Delete blog id %d success.", ID)
}

func NewBlogHandler(blogUsecase usecase.IBlogUsecase) BlogHandler {
	return BlogHandler{
		uc: blogUsecase,
	}
}
