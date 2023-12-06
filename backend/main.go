package main

import (
	"fmt"
	"log"
	"os"
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

	rootPath := os.Getenv("ROOT_PATH")
	vaka := usecases.GetFlatFileTree(rootPath)
	fmt.Print(vaka)
	// TODO GRACEFULLY HANDLE ERRORS/SHUTDOWN AND START
	// addRoutesToApp()
	// http.HandleFunc("/", router.HandleRouting)
	// http.ListenAndServe(ADDR, nil)
}

func addRoutesToApp() {
	router.Routes.AddRoute(usecases.GetFileTreeUseCaseRoute)
	router.Routes.AddRoute(usecases.GetAudioFileUseCaseRoute)
}
