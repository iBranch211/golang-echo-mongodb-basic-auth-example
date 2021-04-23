package security

import (
	"github.com/labstack/echo/v4"
	"golang-echo-mongodb-basic-auth-example/repository"
	"golang-echo-mongodb-basic-auth-example/util"
)

type AuthValidator struct {
	userRepository repository.UserRepository
}

func NewAuthValidator(userRepository repository.UserRepository) *AuthValidator {
	return &AuthValidator{userRepository: userRepository}
}

func (authValidator *AuthValidator) ValidateCredentials(username, password string, c echo.Context) (bool, error) {
	user, err := authValidator.userRepository.FindByEmail(username)
	if err != nil || util.VerifyPassword(user.Password, password) != nil {
		return false, nil
	}
	return true, nil
}
