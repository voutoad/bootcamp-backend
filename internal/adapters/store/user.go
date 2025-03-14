package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/voutoad/bootcamp-backend/internal/domain/dto"
	"github.com/voutoad/bootcamp-backend/internal/domain/ent"
	"golang.org/x/crypto/bcrypt"
)

type UserStore interface {
	CreateUser(user *dto.CreateUserDTO) (*dto.UserResponseDTO, error)
	GetUser(id uuid.UUID) error
	UpdateUser(id uuid.UUID, user string) error
	DeleteUser(id uuid.UUID) error
}

func NewUserStore(client *ent.Client) UserStore {
	return &userStore{
		client: client,
	}
}

var _ UserStore = (*userStore)(nil)

type userStore struct {
	client *ent.Client
}

func (s *userStore) CreateUser(user *dto.CreateUserDTO) (*dto.UserResponseDTO, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u, err := s.client.User.Create().
		SetUsername(user.Username).
		SetPassword(string(password)).
		SetAge(user.Age).
		SetInterests(user.Interests).
		SetRating(user.Rating).
		SetDescription(user.Description).
		SetNillableImageURL(user.ImageURL).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return &dto.UserResponseDTO{
		Username:    u.Username,
		Age:         u.Age,
		Interests:   u.Interests,
		Rating:      u.Rating,
		Description: u.Description,
		ImageURL:    u.ImageURL,
	}, nil
}

func (s *userStore) GetUser(id uuid.UUID) error {
	return nil
}

func (s *userStore) UpdateUser(id uuid.UUID, user string) error {
	return nil
}

func (s *userStore) DeleteUser(id uuid.UUID) error {
	return nil
}
