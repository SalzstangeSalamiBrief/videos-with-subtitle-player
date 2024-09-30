package main

import (
	"backend/lib"
	"backend/middleware"
	"backend/router"
	usecases "backend/useCases"
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

	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
		<-exit
		fmt.Printf("Shutting down server at %v\n", ADDR)
		os.Exit(0)
	}()

	lib.InitializeFileTree()
	addRoutesToApp()

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./public")))

	routerWithMiddlewares := addRouterWithMiddlewares()
	mux.Handle("/api/", routerWithMiddlewares)
	http.ListenAndServe(ADDR, mux)
}

func addRoutesToApp() {
	router.Routes.AddRoute(usecases.GetContinuousFileUseCaseRoute)
	router.Routes.AddRoute(usecases.GetDiscreteFileUseCaseRoute)
	router.Routes.AddRoute(usecases.GetFileTreeUseCaseRoute)
}

func addRouterWithMiddlewares() http.Handler {
	routerHandler := http.HandlerFunc(router.Router)
	corsHandler := middleware.CorsHandler(routerHandler)
	requestLoggerMiddleware := middleware.RequestLoggerHandler(corsHandler)
	return requestLoggerMiddleware
}
