package dto

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CreateUserDTO struct {
	Username    string  `json:"username"`
	Password    string  `json:"password"`
	Age         int     `json:"age"`
	Rating      int     `json:"rating"`
	Description string  `json:"description"`
	Interests   string  `json:"interests"`
	ImageURL    *string `json:"image_url,omitempty"`
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
	Username    string  `json:"username"`
	Age         int     `json:"age"`
	Rating      int     `json:"rating"`
	Description string  `json:"description"`
	Interests   string  `json:"interests"`
	ImageURL    *string `json:"image_url,omitempty"`
	Tags        *string `json:"tags,omitempty"`
}
