package store

import (
	"context"

	"github.com/voutoad/bootcamp-backend/internal/domain/dto"
	"github.com/voutoad/bootcamp-backend/internal/domain/ent"
	"github.com/voutoad/bootcamp-backend/internal/domain/ent/predicate"
	"github.com/voutoad/bootcamp-backend/internal/domain/ent/user"
)

type UserStore interface {
	CreateUser(user *dto.CreateUserDTO) (*dto.UserResponseDTO, error)
	GetUserByUsername(username string) (*dto.UserResponseDTO, error)
	GetUsersWithFilters(query *dto.UsersQueryDTO) ([]*dto.UserResponseDTO, error)
	UpdateUserTags(user *dto.UserUpdateDTO) error
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
	u, err := s.client.User.Create().
		SetUsername(user.Username).
		SetAge(user.Age).
		SetInterests(user.Interests).
		SetRating(user.Rating).
		SetDescription(user.Description).
		SetType(user.Type).
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
		Type:        u.Type,
		ImageURL:    u.ImageURL,
	}, nil
}

func (s *userStore) GetUserByUsername(username string) (*dto.UserResponseDTO, error) {
	u, err := s.client.User.Query().Where(user.Username(username)).First(context.Background())
	if err != nil {
		return nil, err
	}

	return &dto.UserResponseDTO{
		Username:    u.Username,
		Age:         u.Age,
		Interests:   u.Interests,
		Rating:      u.Rating,
		Description: u.Description,
		Type:        u.Type,
		ImageURL:    u.ImageURL,
	}, nil
}

func (s *userStore) GetUsersWithFilters(query *dto.UsersQueryDTO) ([]*dto.UserResponseDTO, error) {
	q := s.client.User.Query()
	if query.Tags != nil {
		q = q.Where(user.Or(createPredictsTags(*query.Tags)...))
	}
	if query.Type != nil {
		q = q.Where(user.Type(*query.Type))
	}

	users, err := q.All(context.Background())
	if err != nil {
		return nil, err
	}

	resp := make([]*dto.UserResponseDTO, len(users))
	for i, u := range users {
		resp[i] = &dto.UserResponseDTO{
			Username:    u.Username,
			Age:         u.Age,
			Interests:   u.Interests,
			Rating:      u.Rating,
			Description: u.Description,
			Type:        u.Type,
			ImageURL:    u.ImageURL,
		}
	}
	return resp, nil
}

func (s *userStore) UpdateUserTags(u *dto.UserUpdateDTO) error {
	n, err := s.client.User.Update().Where(user.Username(u.Username)).SetTags(u.Tags).Save(context.Background())
	if err != nil {
		return err
	}
	if n == 0 {
		return &ent.NotFoundError{}
	}
	return nil
}

func createPredictsTags(tags []*string) []predicate.User {
	valiuesTags := make([]string, len(tags))
	for i, tag := range tags {
		valiuesTags[i] = *tag
	}
	resp := make([]predicate.User, 0)
	resp = append(resp, user.TagsIn(valiuesTags...))
	resp = append(resp, user.TagsIsNil())
	return resp
}
