package setup

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/voutoad/bootcamp-backend/internal/adapters/controllers/rest/handler"
	"github.com/voutoad/bootcamp-backend/internal/adapters/controllers/rest/ping"
	"github.com/voutoad/bootcamp-backend/internal/adapters/controllers/rest/user"
	"github.com/voutoad/bootcamp-backend/internal/adapters/store"
)

type Rest interface {
	InitRoutes(app *fiber.App)
}

type rest struct {
	pingHandler handler.Handler
	userHandler handler.Handler
}

func NewRest(userStore store.UserStore) Rest {
	v := validator.New()
	return &rest{
		pingHandler: ping.NewPingHandler(),
		userHandler: user.NewUserHandler(userStore, v),
	}
}

func (r *rest) InitRoutes(app *fiber.App) {
	apiv1 := app.Group("/api/v1")
	r.pingHandler.RegisterRoutes(apiv1)
	r.userHandler.RegisterRoutes(apiv1)
}
