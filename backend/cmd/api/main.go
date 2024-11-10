package main

import (
	"backend/internal/config"
	"backend/internal/router"
	"backend/internal/routes"
	"backend/pkg/api/middlewares"
	"backend/pkg/services/fileTreeManager"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const ADDR = "localhost:3000"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.InitializeConfiguration()

	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
		<-exit
		fmt.Printf("Shutting down server at %v\n", ADDR)
		os.Exit(0)
	}()

	fileTreeManager.InitializeFileTree()

	r := router.
		NewRouterBuilder().
		RegisterRoute(routes.GetContinuousFileRoute).
		RegisterRoute(routes.GetDiscreteFileUseCaseRoute).
		RegisterRoute(routes.GetFileTreeRoute).
		RegisterMiddleware(middlewares.RequestLoggerMiddleware).
		RegisterMiddleware(middlewares.CorsMiddleware).
		Build()

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./public")))

	mux.Handle("/api/", r)
	http.ListenAndServe(ADDR, mux)
}
