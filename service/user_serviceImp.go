package service

import (
	"context"
	"errors"
	"fmt"
	"shortly/model/domain"
	"shortly/model/dto"
	"shortly/repository"

	"golang.org/x/crypto/bcrypt"
)

type userServiceImp struct {
	UserRepository repository.UserRepository
}

func (service *userServiceImp) Save(ctx context.Context, input dto.CreateUserInput) (domain.User, error) {

	if input.Username == "" || input.Password == "" {
		return domain.User{}, errors.New("Username & Password not requaired")
	}
	//hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to hash password: %w", err)
	}
}
