package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"shortly/model/dto"
	"shortly/service"

	"github.com/gorilla/mux"
)

type urlControllerImp struct {
	urlService service.UrlService
}

func NewUrlController(urlService service.UrlService) UrlController {
	return &urlControllerImp{urlService: urlService}

}

func (controller *urlControllerImp) Save(writer http.ResponseWriter, request *http.Request) {

	// ctx := request.Context()

	// var req dto.CreateURLRequest
	// if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
	// 	http.Error(writer, "invalid request body", http.StatusBadRequest)
	// 	return
	// }

	// if req.LongURL == "" {
	// 	http.Error(writer, "long_url is required", http.StatusBadRequest)
	// 	return
	// }

	// // panggil service
	// url, err := controller.urlService.Save(ctx, req.LongURL)
	// if err != nil {
	// 	http.Error(writer, "failed to save url", http.StatusInternalServerError)
	// 	return
	// }
	// // mapping domain → DTO
	// resp := dto.CreateURLResponse{
	// 	Code:     url.Code,
	// 	ShortURL: "https://sho.rt/" + url.Code,
	// }
	// // return JSON
	// writer.Header().Set("Content-Type", "application/json")
	// writer.WriteHeader(http.StatusCreated)
	// json.NewEncoder(writer).Encode(resp)

	ctx := request.Context()

	var req dto.CreateURLRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(dto.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   "invalid request",
		})
		return
	}

	url, err := controller.urlService.Save(ctx, req.LongURL)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(dto.Response{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   nil,
		})
		return
	}

	response := dto.CreateURLResponse{
		Code:     url.Code,
		ShortURL: "http://localhost:8080/" + url.Code,
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(dto.Response{
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

// func (controller *urlControllerImp) Redirect(writer http.ResponseWriter, request *http.Request) {
// 	ctx := request.Context()
// 	code := mux.Vars(request)["code"]

// 	url, err := controller.urlService.Redirect(ctx, code)
// 	if err != nil {
// 		http.NotFound(writer, request)
// 		return
// 	}
// 	http.Redirect(writer, request, url.LongURL, http.StatusFound)

// }

func (controller *urlControllerImp) RedirectAndIncrement(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	code := mux.Vars(request)["code"]
	log.Println("Route hit, code =", mux.Vars(request)["code"])

	url, err := controller.urlService.RedirectAndIncrement(ctx, code)
	if err != nil {
		log.Println("Redirect error:", err)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(map[string]string{"error": "URL not found"})
		return
	}

	// http.Redirect(writer, request, url.LongURL, http.StatusFound)

	// log.Println("Redirecting to:", url.LongURL, "Hits now:", url.HitCount)
	// Kembalikan JSON
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(url)
}
