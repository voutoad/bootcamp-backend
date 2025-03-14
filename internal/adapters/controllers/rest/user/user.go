package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/voutoad/bootcamp-backend/internal/adapters/controllers/rest/handler"
	"github.com/voutoad/bootcamp-backend/internal/adapters/store"
	"github.com/voutoad/bootcamp-backend/internal/domain/dto"
)

type UserHandler interface {
	handler.Handler
	createUser(c *fiber.Ctx) error
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
	group.Post("/users", h.createUser)
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
