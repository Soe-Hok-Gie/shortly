package service

import (
	"context"
	"errors"
	"fmt"
	"shortly/model/domain"
	"shortly/model/dto"
	"shortly/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type userServiceImp struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userServiceImp{UserRepository: userRepository}
}

func (service *userServiceImp) Register(ctx context.Context, input dto.CreateUserInput) (dto.UserResponse, error) {

	if input.Username == "" || input.Password == "" {
		return dto.UserResponse{}, errors.New("Username & Password not requaired")
	}
	//hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	input.Password = string(hashedPassword)

	if err != nil {
		return dto.UserResponse{}, fmt.Errorf("failed to hash password: %w", err)
	}

	user := domain.User{
		Username:  string(input.Username),
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
	}
	// Simpan ke repository
	user, err = service.UserRepository.Register(ctx, user)
	if err != nil {
		return dto.UserResponse{}, fmt.Errorf("failed save user: %w", err)
	}

	userResponse := dto.UserResponse{
		Username:   user.Username,
		Created_At: user.CreatedAt,
	}
	return userResponse, nil

}

// error global yang dipakai setiap kali login gagal (username tidak ditemukan atau password salah)
var ErrInvalidCredential = errors.New("invalid credential")

func (service *userServiceImp) Login(ctx context.Context, input dto.CreateUserInput) (dto.UserResponse, error) {

	var userResponse dto.UserResponse

	user, err := service.UserRepository.Login(ctx, input.Username)
	if err != nil {
		return userResponse, ErrInvalidCredential
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		fmt.Println("Password mismatch!")
		return userResponse, ErrInvalidCredential
	}

	userResponse = dto.UserResponse{
		Username: user.Username,
	}
	return userResponse, nil
}
