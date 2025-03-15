package dto

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CreateUserDTO struct {
	Username    string `json:"username" validate:"required,min=2,max=100"`
	Age         int    `json:"age" validate:"required,min=0,max=100"`
	Rating      int    `json:"rating" validate:"required,min=0,max=100"`
	Description string `json:"description" validate:"required,min=2,max=100"`
	Interests   string `json:"interests" validate:"required,min=2,max=100"`
	Type        string `json:"type" validate:"required,min=2,max=100"`

	ImageURL *string `json:"image_url,omitempty" validate:"omitempty,url"`
}

func (u *CreateUserDTO) Validate(c *fiber.Ctx, v *validator.Validate) error {
	if err := c.BodyParser(u); err != nil {
		return err
	}
	if err := v.Struct(u); err != nil {
		return err
	}
	return nil
}

type UserResponseDTO struct {
	Username    string `json:"username"`
	Age         int    `json:"age"`
	Rating      int    `json:"rating"`
	Description string `json:"description"`
	Interests   string `json:"interests"`
	Type        string `json:"type"`

	ImageURL *string `json:"image_url,omitempty"`
	Tags     *string `json:"tags,omitempty"`
}

type UsersQueryDTO struct {
	Tags *[]*string `query:"tag,omitempty"`
	Type *string    `query:"type,omitempty"`
}

type UserUpdateDTO struct {
	Username string `json:"username" validate:"required,min=2,max=100"`
	Tags     string `json:"tags" validate:"required,min=4"`
}

func (u *UserUpdateDTO) Validate(c *fiber.Ctx, v *validator.Validate) error {
	if err := c.BodyParser(u); err != nil {
		return err
	}
	if err := v.Struct(u); err != nil {
		return err
	}
	if strings.Join(strings.Split(u.Tags, ", "), ", ") != u.Tags {
		return errors.New("tags must be a comma-separated list")
	}
	return nil
}
