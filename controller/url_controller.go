package controller

import "net/http"

type UrlController interface {
	Save(writer http.ResponseWriter, request *http.Request)
}
