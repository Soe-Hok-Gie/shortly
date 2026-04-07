package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"shortly/model/domain"
	"shortly/model/dto"
	"shortly/repository"
	"shortly/utils"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type userServiceImp struct {
	UserRepository         repository.UserRepository
	refreshTokenRepository repository.RefreshTokenRepository
}

func NewUserService(
	userRepository repository.UserRepository,
	refreshTokenRepository repository.RefreshTokenRepository,
) UserService {
	return &userServiceImp{UserRepository: userRepository,
		refreshTokenRepository: refreshTokenRepository,
	}
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

func (service *userServiceImp) Login(ctx context.Context, input dto.CreateUserInput) (dto.LoginResponse, error) {

	user, err := service.UserRepository.Login(ctx, input.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("username not found!")
			return dto.LoginResponse{}, ErrInvalidCredential
		}
		return dto.LoginResponse{}, ErrInternal

	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		fmt.Println("Password mismatch!")
		return dto.LoginResponse{}, ErrInvalidCredential
	}

	token, err := utils.GenerateToken(user.Id)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	tokenrefresh, err := utils.GenerateRefreshToken(user.Id)

	if err != nil {
		return dto.LoginResponse{}, err
	}

	//hash sebelum save db
	hashed := utils.Hashrefresh(tokenrefresh)
	err = service.refreshTokenRepository.Save(ctx, &domain.RefreshToken{
		UserID:    user.Id,
		TokenHash: hashed,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	})

	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{
		AccessToken:  token,
		TokenType:    "JWT",
		RefreshToken: tokenrefresh,
	}, nil
}

func (service userServiceImp) Refresh(ctx context.Context, refreshtoken string) (dto.LoginResponse, error) {
	hashed := utils.Hashrefresh(refreshtoken)

	stored, err := service.refreshTokenRepository.FindByHash(ctx, hashed)
	fmt.Println("stored", stored)
	if err != nil {
		return dto.LoginResponse{}, errors.New("refresh token invalid")
	}
	//cek expired
	if time.Now().After(stored.ExpiresAt) {
		return dto.LoginResponse{}, errors.New("refresh token by expired")
	}

	//generate accsess token baru
	access, _ := utils.GenerateToken(stored.UserID)

	// refresh token tetap sama → tidak di-rotate
	return dto.LoginResponse{
		AccessToken:  access,
		RefreshToken: refreshtoken,
	}, nil

}

func (service userServiceImp) DeleteRefresh(ctx context.Context, hashed string) error {
	return service.refreshTokenRepository.DeleteByHash(ctx, hashed)

}
