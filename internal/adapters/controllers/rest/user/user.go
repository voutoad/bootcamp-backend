package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/voutoad/bootcamp-backend/internal/adapters/controllers/rest/handler"
	"github.com/voutoad/bootcamp-backend/internal/adapters/store"
	"github.com/voutoad/bootcamp-backend/internal/domain/dto"
	"github.com/voutoad/bootcamp-backend/internal/domain/ent"
)

type UserHandler interface {
	handler.Handler
	createUser(c *fiber.Ctx) error
	getUser(c *fiber.Ctx) error
	getUsers(c *fiber.Ctx) error
}

type userHandler struct {
	userStore store.UserStore
	validator *validator.Validate
}

func NewUserHandler(userStore store.UserStore, validator *validator.Validate) UserHandler {
	return &userHandler{
		userStore: userStore,
		validator: validator,
	}
}

func (h *userHandler) RegisterRoutes(group fiber.Router) {
	users := group.Group("/users")
	users.Post("/", h.createUser)
	users.Get("/", h.getUsers)
	users.Patch("/", h.updateUserTags)
	users.Get("/:username", h.getUser)
}

// createUser godoc
// @Summary      Create user
// @Description  Create a new user with given body
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user   body      dto.CreateUserDTO  true  "User data"
// @Success      201  {object}  dto.UserResponseDTO
// @Failure      400
// @Failure      500
// @Router       /users [post]
func (h *userHandler) createUser(c *fiber.Ctx) error {
	user := &dto.CreateUserDTO{}
	if err := user.Validate(c, h.validator); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	resp, err := h.userStore.CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(resp)
}

// getUserByUsername godoc
// @Summary      Get user by his username
// @Description  Get user with given username or get a 404 error
// @Tags         users
// @Produce      json
// @Param        username   path      string  true  "Username"
// @Success      200  {object}  dto.UserResponseDTO
// @Failure      400
// @Failure      500
// @Router       /users/{username} [get]
func (h *userHandler) getUser(c *fiber.Ctx) error {
	username := c.Params("username")
	user, err := h.userStore.GetUserByUsername(username)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

// getUsers godoc
// @Summary      Get users
// @Description  Get user with optional filters
// @Tags         users
// @Produce      json
// @Param        type   query      string  true  "Type"
// @Param        tag   query      string  true  "Tag"
// @Success      200  {object}  []dto.UserResponseDTO
// @Failure      400
// @Failure      500
// @Router       /users [get]
func (h *userHandler) getUsers(c *fiber.Ctx) error {
	query := &dto.UsersQueryDTO{}
	if err := c.QueryParser(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	users, err := h.userStore.GetUsersWithFilters(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(users)
}

// updateUserTags godoc
// @Summary      Update user tags
// @Description  Updating user tags with given username and tags
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        userUpdate   body      dto.UserUpdateDTO  true  "User Update"
// @Success      204
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /users [patch]
func (h *userHandler) updateUserTags(c *fiber.Ctx) error {
	req := &dto.UserUpdateDTO{}
	if err := req.Validate(c, h.validator); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := h.userStore.UpdateUserTags(req); err != nil {
		if ent.IsNotFound(err) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusNoContent).JSON(nil)
}
