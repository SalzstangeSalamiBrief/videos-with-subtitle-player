package main

import (
	"backend/router"
	directorytree "backend/services/directoryTree"
	usecases "backend/useCases"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

const ADDR = "localhost:3000"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// TODO GRACEFULLY HANDLE ERRORS/SHUTDOWN AND START
	directorytree.InitializeFileTree()
	addRoutesToApp()
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/api/", router.HandleRouting)
	http.ListenAndServe(ADDR, nil)
}

func addRoutesToApp() {
	router.Routes.AddRoute(usecases.GetAudioFileUseCaseRoute)
	router.Routes.AddRoute(usecases.GetFileTreeUseCaseRoute)
}
