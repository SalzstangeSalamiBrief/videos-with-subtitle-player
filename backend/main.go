package main

import (
	"log"
	"net/http"
	"videos-with-subtitle-player/router"
	usecases "videos-with-subtitle-player/useCases"

	"github.com/joho/godotenv"
)

const ADDR = "localhost:3000"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	addRoutesToApp()
	http.HandleFunc("/", router.HandleRouting)
	http.ListenAndServe(ADDR, nil)
}

func addRoutesToApp() {
	router.Routes.AddRoute(usecases.GetFileTreeUseCaseRoute)
}
