package users

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// Initialize the validator
func init() {
	validate = validator.New()
}

type CreateUserRequest struct {
	Name string `json:"name" validate:"required,min=1,max=100"`
	Role string `json:"role" validate:"required,oneof=admin user guest"`
}

// ValidateCreateUser validates a user creation request
func ValidateCreateUser(req *CreateUserRequest) error {
	return validate.Struct(req)
}
