package validator

import (
	"errors"

	"example.com/mittaus/ddd-example/application"
	"example.com/mittaus/ddd-example/domain"
	"github.com/asaskevich/govalidator"
	// "example.com/mittaus/ddd-example/application"
)

type userValidator struct{}

func New() application.UserValidator {
	return userValidator{}
}

func (userValidator) CheckUser(user domain.User) error {
	if ok := govalidator.IsEmail(user.Email); !ok {
		return errors.New("invalid email")
	}

	return nil
}
