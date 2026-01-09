package service

import (
	"context"
	"shortly/model/domain"
	"shortly/model/dto"
	"shortly/repository"
)

type userServiceImp struct {
	UserRepository repository.UserRepository
}

func (service *userServiceImp) Save(ctx context.Context, input dto.CreateUserInput) (domain.User, error) {

}
