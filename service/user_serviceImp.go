package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
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

var (
	// error global yang dipakai setiap kali register gagal (username & password are required)
	ErrInvalidInput = errors.New("username & password are required")

	// error global yang dipakai setiap kali register gagal (username duplikat)
	ErrUsernameExists = errors.New("username already exists")

	// error global yang dipakai setiap kali login gagal (username tidak ditemukan atau password salah)
	ErrInvalidCredential = errors.New("invalid credential")

	ErrInternal = errors.New("internal server error")
)

func (service *userServiceImp) Register(ctx context.Context, input dto.CreateUserInput) (dto.UserResponse, error) {

	if input.Username == "" || input.Password == "" {
		return dto.UserResponse{}, ErrInvalidInput
	}
	//hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		return dto.UserResponse{}, fmt.Errorf("failed to hash password: %w", err)
	}

	user := domain.User{
		Username:  input.Username,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
	}
	// Simpan ke repository
	user, err = service.UserRepository.Register(ctx, user)
	if err != nil {
		if repository.IsDuplicateKeyError(err) {
			return dto.UserResponse{}, ErrUsernameExists
		}
	}

	userResponse := dto.UserResponse{
		Username:   user.Username,
		Created_At: user.CreatedAt,
	}
	return userResponse, nil

}

func (service *userServiceImp) Login(ctx context.Context, input dto.CreateUserInput) (dto.UserResponse, error) {

	var userResponse dto.UserResponse

	user, err := service.UserRepository.Login(ctx, input.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("user not found: username=%s, dbErr=%v", input.Username, err)
			return dto.UserResponse{}, ErrInvalidCredential
		}
		return dto.UserResponse{}, ErrInternal

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
