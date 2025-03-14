package ping

import (
	"github.com/gofiber/fiber/v2"
	"github.com/voutoad/bootcamp-backend/internal/adapters/controllers/rest/handler"
)

type pingHandler struct{}

func NewPingHandler() handler.Handler {
	return &pingHandler{}
}

func (h *pingHandler) RegisterRoutes(group fiber.Router) {
	group.Get("/ping", h.ping)
}

func (h *pingHandler) ping(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(map[string]string{"message": "BOOTCAAAAMP"})
}
