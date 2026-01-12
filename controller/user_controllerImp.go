package controller

import (
	"encoding/json"
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

func (controller *userControllerImp) Save(writer http.ResponseWriter, request *http.Request) {
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

	result, err := controller.userService.Save(ctx, req)
	if err != nil {
		status := http.StatusBadRequest
		if err.Error() == "username already exists" {
			status = http.StatusConflict
		}

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
