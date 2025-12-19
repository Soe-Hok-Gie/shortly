package controller

import (
	"net/http"
	"shortly/service"
)

type urlControllerImp struct {
	urlService service.UrlService
}

func NewUrlController(urlService service.UrlService) UrlController {
	return &urlControllerImp{urlService: urlService}

}

func (controller *urlControllerImp) Save(writer http.ResponseWriter, request *http.Request) {

}
