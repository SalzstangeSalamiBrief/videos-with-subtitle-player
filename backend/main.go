package main

import (
	"backend/router"
	directorytree "backend/services/directoryTree"
	usecases "backend/useCases"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

const ADDR = "localhost:3000"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
		<-exit
		fmt.Printf("Shutting down server at %v\n", ADDR)
		os.Exit(0)
	}()

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
