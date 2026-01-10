package controller

import (
	"net/http"
	"shortly/service"
)

type userControllerImp struct {
	userService service.UserService
}

func (controller *userControllerImp) Save(writer http.ResponseWriter, request *http.Request) {

}
