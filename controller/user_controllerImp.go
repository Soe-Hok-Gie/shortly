package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"shortly/middleware"
	"shortly/model/dto"
	"shortly/service"
	"shortly/utils"
)

type userControllerImp struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) userController {
	return &userControllerImp{userService: userService}
}

func (controller *userControllerImp) Register(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	var req dto.CreateUserInput
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(dto.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   "Input salah",
		})
		return
	}

	result, err := controller.userService.Register(ctx, req)
	if err != nil {
		status := http.StatusInternalServerError
		msg := "Internal server error"

		// Mapping error ke pesan & status code
		switch {
		case errors.Is(err, service.ErrInvalidInput):
			status = http.StatusBadRequest
			msg = "Username & password are required"
		case errors.Is(err, service.ErrUsernameExists):
			status = http.StatusConflict
			msg = "username already exists"
		}

		writer.WriteHeader(status)
		json.NewEncoder(writer).Encode(dto.Response{
			Code:   status,
			Status: http.StatusText(status),
			Data:   msg,
		})
		return
	}

	// Success response
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated) // 201
	json.NewEncoder(writer).Encode(dto.Response{
		Code:   http.StatusCreated,
		Status: "Create",
		Data:   result,
	})
}

// login
func (controller *userControllerImp) Login(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	var input dto.CreateUserInput

	if err := json.NewDecoder(request.Body).Decode(&input); err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(dto.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   "Input salah",
		})
		return
	}

	userResponse, err := controller.userService.Login(ctx, input)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		//errorcredensial
		if err == service.ErrInvalidCredential {
			writer.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(writer).Encode(dto.Response{
				Code:   http.StatusUnauthorized,
				Status: "unauthorize",
				Data:   "Username or password failed",
			})
			return
		}

		//internal server error
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(dto.Response{
			Code:   http.StatusInternalServerError,
			Status: "internal server error",
			Data:   "server error",
		})
		fmt.Println("pap ma", err)
		return
	}

	//cookie
	http.SetCookie(writer, &http.Cookie{
		//tanpa refresh
		// Name:     "token",
		// Value:    userResponse.AccessToken,
		Name:     "refresh_token",
		Value:    userResponse.RefreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		MaxAge:   900,
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(dto.Response{
		Code:   http.StatusOK,
		Status: "Success",
		//tanpa refresh token
		// Data:   userResponse,
		Data: userResponse.AccessToken,
	})
}

func (controller *userControllerImp) Refresh(writer http.ResponseWriter, request *http.Request) {

	cookie, err := request.Cookie("refresh_token")
	if err != nil {
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userResponse, err := controller.userService.Refresh(request.Context(), cookie.Value)
	fmt.Println("refresh token received:", userResponse.RefreshToken)

	if err != nil {
		fmt.Println("refresh error:", err)

		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if userResponse.RefreshToken != "" {
		http.SetCookie(writer, &http.Cookie{
			Name:     "refresh_token",
			Value:    userResponse.RefreshToken,
			HttpOnly: true,
			Path:     "/",
			MaxAge:   7 * 24 * 60 * 60,
		})

		json.NewEncoder(writer).Encode(map[string]string{
			"access_token": userResponse.AccessToken,
		})
	}
}

// logout
func (controller *userControllerImp) Logout(writer http.ResponseWriter, request *http.Request) {

	cookie, err := request.Cookie("refresh_token")
	if err != nil {
		log.Println("cookie refresh_token tidak ada")
		return
	}
	refreshToken := cookie.Value
	log.Println("refresh token:", refreshToken)

	//hashed
	hashed := utils.Hashrefresh(refreshToken)
	//delete token di DB
	log.Println("before delete refresh")

	err = controller.userService.DeleteRefresh(request.Context(), hashed)
	log.Println("after delete refresh")
	if err != nil {
		log.Println("delete", err)
	}

	http.SetCookie(writer, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	})
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(dto.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   "Successfully logged out",
	})
}

// profile
func (controller *userControllerImp) Profile(writer http.ResponseWriter, request *http.Request) {
	Id, ok := request.Context().Value(middleware.UserIdKey).(int64)
	if !ok {
		writer.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(writer).Encode(dto.Response{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
			Data:   nil,
		})
		return
	}

	userResponse := map[string]interface{}{
		"user_id": Id,
		"message": "hello",
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(dto.Response{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   userResponse,
	})

}
