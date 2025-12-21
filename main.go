package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"shortly/app"
	"shortly/controller"
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

	urlRepository := repository.NewUrlRepository(db)
	urlService := service.NewUrlService(urlRepository)
	urlController := controller.NewUrlController(urlService)

	r := mux.NewRouter()
	r.HandleFunc("/url", urlController.Save).Methods("Post")
	log.Fatal(http.ListenAndServe(":8080", r))

}
