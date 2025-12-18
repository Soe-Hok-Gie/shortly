package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"shortly/app"

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
	db, err := app.NewDB(userDB, passDB, hostDB, portDB, nameDB)
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", r))

}
