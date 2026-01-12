package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"shortly/app"
	"shortly/controller"
	"shortly/middleware"
	"shortly/repository"
	"shortly/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("shortly")
	//setENV
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	userDB := os.Getenv("DB_USER")
	passDB := os.Getenv("DB_PASS")
	hostDB := os.Getenv("DB_HOST")
	portDB := os.Getenv("DB_PORT")
	nameDB := os.Getenv("DB_NAME")

	fmt.Println("dsn:", userDB, passDB, hostDB, portDB, nameDB)

	//setDB
	db := app.NewDB(userDB, passDB, hostDB, portDB, nameDB)
	//url
	urlRepository := repository.NewUrlRepository(db)
	urlService := service.NewUrlService(urlRepository)
	urlController := controller.NewUrlController(urlService)

	//user
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	r := mux.NewRouter()
	r.HandleFunc("/url", urlController.Save).Methods("POST")
	// r.HandleFunc("/topvisited", urlController.GetTopVisited).Methods("GET")//sebelum ada middleware
	rateLimitMiddleware := middleware.NewRateLimitMiddleware()
	// r.HandleFunc("/{code}", urlController.RedirectAndIncrement).Methods("GET")//sebelum ada middleware
	r.Handle("/code/{code}", rateLimitMiddleware.WithRateLimit()(http.HandlerFunc(urlController.RedirectAndIncrement)))
	r.Handle("/topvisited", rateLimitMiddleware.WithRateLimit()(http.HandlerFunc(urlController.GetTopVisited)))
	r.HandleFunc("/user/register", userController.Save).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))

}
