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

	// TODO GRACEFULLY HANDLE ERRORS/SHUTDOWN AND START

	addRoutesToApp()
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/api/", router.HandleRouting)
	http.ListenAndServe(ADDR, nil)
}

func addRoutesToApp() {
	router.Routes.AddRoute(usecases.GetFileTreeUseCaseRoute)
	router.Routes.AddRoute(usecases.GetAudioFileUseCaseRoute)
}
