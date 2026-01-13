package service

import (
	"context"
	"errors"
	"fmt"
	"shortly/model/domain"
	"shortly/model/dto"
	"shortly/repository"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type userServiceImp struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userServiceImp{UserRepository: userRepository}
}

func (service *userServiceImp) Save(ctx context.Context, input dto.CreateUserInput) (dto.UserResponse, error) {

	if input.Username == "" || input.Password == "" {
		return dto.UserResponse{}, errors.New("Username & Password not requaired")
	}
	//hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.UserResponse{}, fmt.Errorf("failed to hash password: %w", err)
	}

	user := domain.User{
		Username:  strings.ToLower(input.Username),
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
	}
	// Simpan ke repository
	user, err = service.UserRepository.Save(ctx, user)
	if err != nil {
		return dto.UserResponse{}, fmt.Errorf("gagal menyimpan user: %w", err)
	}

	userResponse := dto.UserResponse{
		Username:   user.Username,
		Created_At: user.CreatedAt,
	}
	return userResponse, nil

}

func (service *userServiceImp) Login(ctx context.Context, input dto.CreateUserInput) (dto.UserResponse, error) {

}
