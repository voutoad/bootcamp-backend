package postgres

import (
	"fmt"
	"log"

	"github.com/voutoad/bootcamp-backend/internal/config"
	"github.com/voutoad/bootcamp-backend/internal/domain/ent"

	_ "github.com/lib/pq"
)

func NewDB(cfg *config.Config) (*ent.Client, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Name,
		cfg.DB.Password,
	)
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to open psql with error: %v", err)
		return nil, err
	}
	return client, nil
}
