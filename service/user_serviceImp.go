package service

import (
	"context"
	"errors"
	"shortly/model/domain"
	"shortly/model/dto"
	"shortly/repository"
)

type userServiceImp struct {
	UserRepository repository.UserRepository
}

func (service *userServiceImp) Save(ctx context.Context, input dto.CreateUserInput) (domain.User, error) {

	if input.Username == "" || input.Password == "" {
		return domain.User{}, errors.New("Username & Password not requaired")
	}

}
