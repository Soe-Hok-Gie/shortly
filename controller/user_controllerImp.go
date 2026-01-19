package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"shortly/model/dto"
	"shortly/service"
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
		status := http.StatusConflict
		if err.Error() == "username already exists" {
			status = http.StatusConflict
		}
		log.Println("username already exists")

		writer.WriteHeader(status)
		json.NewEncoder(writer).Encode(dto.Response{
			Code:   status,
			Status: http.StatusText(status),
			Data:   err.Error(),
		})
		return

	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(dto.Response{
		Code:   http.StatusCreated,
		Status: "Register Sukses",
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
			Data:   err.Error(),
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
