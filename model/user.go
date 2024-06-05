package model

import "github.com/go-playground/validator/v10"

var Validate = validator.New()


type User struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json: "password" validate:"required,min=6,max=128"`

}

type LoginRequest struct{
	User
}
