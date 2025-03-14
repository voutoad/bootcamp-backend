package application

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/voutoad/bootcamp-backend/internal/adapters/controllers/rest/setup"
	"github.com/voutoad/bootcamp-backend/internal/adapters/database/postgres"
	"github.com/voutoad/bootcamp-backend/internal/adapters/store"
	"github.com/voutoad/bootcamp-backend/internal/config"
	"github.com/voutoad/bootcamp-backend/internal/domain/ent"

	_ "github.com/voutoad/bootcamp-backend/docs"
)

type Application interface {
	Start() error
}

type application struct {
	app           *fiber.App
	db            *ent.Client
	serverAddress string
}

func NewApplication(cfg *config.Config) (Application, error) {
	db, err := postgres.NewDB(cfg)
	if err != nil {
		return nil, err
	}
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	app.Get("/swagger/*", swagger.HandlerDefault)
	userStore := store.NewUserStore(db)
	rest := setup.NewRest(userStore)
	rest.InitRoutes(app)
	return &application{
		app:           app,
		serverAddress: cfg.Server.Address,
		db:            db,
	}, nil
}

func (a *application) Start() error {
	defer a.db.Close()
	if err := a.db.Schema.Create(context.Background()); err != nil {
		return err
	}
	if err := a.app.Listen(a.serverAddress); err != nil {
		return err
	}
	return nil
}
