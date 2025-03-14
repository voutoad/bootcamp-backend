package main

import (
	"github.com/voutoad/bootcamp-backend/internal/application"
	"github.com/voutoad/bootcamp-backend/internal/config"
)

// @title PickMe API
// @version 1.0
// @description Simple backend for PickMe product
// @host localhost:5000
// @BasePath /api/v1
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	app, err := application.NewApplication(cfg)
	if err != nil {
		panic(err)
	}
	if err := app.Start(); err != nil {
		panic(err)
	}
}
