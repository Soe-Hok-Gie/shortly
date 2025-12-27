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

//pakek redirect
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
		json.NewEncoder(writer).Encode(dto.Response{
			Code:   http.StatusNotFound,
			Status: "status NotFound",
			Data:   "url salah",
		})
		return
	}
	// http.Redirect(writer, request, url.LongURL, http.StatusFound)
	// log.Println("Redirecting to:", url.LongURL, "Hits now:", url.HitCount)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(url)
}

func (controller *urlControllerImp) GetTopVisited(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	result, err := controller.urlService.GetTopVisited(ctx)
	if err != nil {
		log.Println("Error fetching top visited URLs:", err)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		// json.NewEncoder(writer).Encode(map[string]string{"error": "Failed to fetch top visited URLs"})
		// return
		json.NewEncoder(writer).Encode(dto.Response{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   nil,
		})
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(result)
}
