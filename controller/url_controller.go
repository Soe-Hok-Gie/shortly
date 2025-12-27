package controller

import "net/http"

type UrlController interface {
	Save(writer http.ResponseWriter, request *http.Request)
	RedirectAndIncrement(writer http.ResponseWriter, request *http.Request)
	GetTopVisited(writer http.ResponseWriter, request *http.Request)
}
