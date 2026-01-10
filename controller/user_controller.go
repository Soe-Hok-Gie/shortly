package controller

import "net/http"

type userController interface {
	Save(writer http.ResponseWriter, request *http.Request)
}
