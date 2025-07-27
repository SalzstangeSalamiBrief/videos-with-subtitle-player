package main

import (
	"backend/internal/config"
	"backend/internal/router"
	"backend/internal/routes"
	"backend/pkg/api/middlewares"
	"backend/pkg/services/ImageQualityReducer"
	"backend/pkg/services/fileTreeManager"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	appConfiguration := config.InitializeConfiguration()
	ImageQualityReducer.InitializeMagickArgs()
	log.Default().Printf("Start server on '%v'", config.AppConfiguration.ServerAddress)

	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
		<-exit
		log.Default().Printf("Shutting down server at %v\n", config.AppConfiguration.ServerAddress)
		os.Exit(0)
	}()

	fileTreeManager.InitializeFileTree()

	corsMiddleware := middlewares.NewCorsMiddleWare().AddConfiguration(middlewares.CorsMiddleWareConfiguration{AllowedCors: appConfiguration.AllowedCors}).Build()
	requestLoggerMiddleware := middlewares.NewRequestLogger().Build()

	r := router.
		NewRouterBuilder().
		RegisterRoute(routes.GetContinuousFileRoute).
		RegisterRoute(routes.GetDiscreteFileUseCaseRoute).
		RegisterRoute(routes.GetFileTreeRoute).
		RegisterMiddleware(requestLoggerMiddleware).
		RegisterMiddleware(corsMiddleware).
		Build()

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./public")))

	mux.Handle("/api/", r)
	http.ListenAndServe(config.AppConfiguration.ServerAddress, mux)
}
