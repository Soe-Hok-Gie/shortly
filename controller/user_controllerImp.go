package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"shortly/middleware"
	"shortly/model/dto"
	"shortly/service"
	"strconv"
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
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(dto.Response{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   userResponse,
	})
}

// profile
func (controller *userControllerImp) Profile(writer http.ResponseWriter, request *http.Request) {
	Id := request.Context().Value(middleware.UserIdKey).(int64)
	writer.Write([]byte("Hello user " + strconv.FormatInt(Id, 10)))
}
