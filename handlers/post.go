package handlers

import (
	"github.com/blockchaindev100/Go-Blog-Site/models"
	"github.com/gofiber/fiber/v2"
)

// @Summary Get Posts
// @Schemes http
// @Description Get all the posts in the blog site
// @Tags Posts
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} []models.Post
// @Router /blog [get]
func (h *Handlers) GetPosts(c *fiber.Ctx) error {
	posts, err := h.Repo.GetPosts()
	if err != nil {
		h.Logger.Error(err)
		return fiber.ErrInternalServerError
	}
	for i := 0; i < len(posts); i++ {
		str := posts[i].Post_id.String()
		commands, err := h.Repo.GetCommandsByPostId(str)
		if err != nil {
			h.Logger.Error(err)
		} else {
			posts[i].Commands = commands
		}
		categories := []models.Category{}
		for j := 0; j < len(posts[i].Categories); j++ {
			category, err := h.Repo.GetCategoriesById(posts[i].Categories[j])
			if err != nil {
				h.Logger.Error(err)
			} else {
				categories = append(categories, category)
			}
		}
		posts[i].Category = categories
	}
	return c.JSON(posts)
}

// @Summary Create Posts
// @Schemes http
// @Description Create post in the blog site
// @Tags Posts
// @Accept json
// @Produce json
// @Param user body models.Post true "Blog details"
// @Security ApiKeyAuth
// @Success 200 {object} models.Response
// @Router /blog [post]
func (h *Handlers) CreatePost(c *fiber.Ctx) error {
	var post models.Post
	id := c.Get("user_id")
	if err := c.BodyParser(&post); err != nil {
		h.Logger.Error(err)
		return fiber.ErrBadRequest
	}
	if err := h.Validator.Struct(post); err != nil {
		h.Logger.Error(err)
		return fiber.ErrBadRequest
	}
	if err := h.Repo.CreatePosts(&post, id); err != nil {
		h.Logger.Error(err)
		return fiber.ErrInternalServerError
	}
	return c.JSON(fiber.Map{
		"message": "post created successfully",
	})
}

// @Summary Update Posts
// @Schemes http
// @Description Update the post in the blog site
// @Tags Posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Param user body models.Post true "Blog details"
// @Security ApiKeyAuth
// @Success 200 {object} models.Response
// @Router /blog [put]
func (h *Handlers) UpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Post
	if err := c.BodyParser(&post); err != nil {
		h.Logger.Error(err)
		return fiber.ErrBadRequest
	}
	if err := h.Validator.Struct(post); err != nil {
		h.Logger.Error(err)
		return fiber.ErrBadRequest
	}
	if err := h.Repo.UpdatePost(&post, id); err != nil {
		h.Logger.Error(err)
		return fiber.ErrInternalServerError
	}
	return c.JSON(fiber.Map{
		"message": "updated successful",
	})
}

// @Summary Delete Posts
// @Schemes http
// @Description Delete the post in the blog site
// @Tags Posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Security ApiKeyAuth
// @Success 200 {object} models.Response
// @Router /blog [Delete]
func (h *Handlers) DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.Repo.DeletePost(id); err != nil {
		h.Logger.Error(err)
		return fiber.ErrInternalServerError
	}
	return c.JSON(fiber.Map{
		"message": "post deleted successful",
	})
}
