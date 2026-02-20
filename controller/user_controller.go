package controller

import "net/http"

type userController interface {
	Register(writer http.ResponseWriter, request *http.Request)
	Login(writer http.ResponseWriter, request *http.Request)
	Logout(writer http.ResponseWriter, request *http.Request)
	Profile(writer http.ResponseWriter, request *http.Request)
}
