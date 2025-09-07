package main

import (
	"backend/internal/config/api"
	"backend/internal/router"
	"backend/internal/routes"
	"backend/pkg/api/handlers"
	"backend/pkg/api/middlewares"
	"backend/pkg/services/fileTreeManager"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	initializedConfiguration, configurationError := api.NewApiConfiguration()
	if configurationError != nil {
		log.Fatal(configurationError)
	}

	initializedFileTreeManager := fileTreeManager.NewFileTreeManager(initializedConfiguration.GetRootPath()).InitializeTree()

	log.Default().Printf("Start server on '%v'", initializedConfiguration.GetServerAddress())

	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
		<-exit
		log.Default().Printf("Shutting down server at %v\n", initializedConfiguration.GetServerAddress())
		os.Exit(0)
	}()

	corsMiddleware := middlewares.NewCorsMiddleWare().AddConfiguration(middlewares.CorsMiddleWareConfiguration{AllowedCors: initializedConfiguration.GetCors()}).Build()
	requestLoggerMiddleware := middlewares.NewRequestLogger().Build()

	handleDiscreteFileRoute := routes.NewGetDiscreteFileByIdRoute(handlers.DiscreteFileByIdHandlerConfig{
		RootPath:        initializedConfiguration.GetRootPath(),
		FileTreeManager: initializedFileTreeManager,
	})
	handleContinousFileRoute := routes.CreateGetContinuousFileRoute(handlers.ContinuousFileByIdHandlerConfiguration{RootPath: initializedConfiguration.GetRootPath(), FileTreeManager: initializedFileTreeManager})
	getFileTreeRoute := routes.NewGetFileTreeRoute(handlers.FileTreeHandlerConfiguration{FileTreeManager: initializedFileTreeManager})

	r := router.
		NewRouterBuilder().
		RegisterRoute(handleContinousFileRoute).
		RegisterRoute(handleDiscreteFileRoute).
		RegisterRoute(getFileTreeRoute).
		RegisterMiddleware(requestLoggerMiddleware).
		RegisterMiddleware(corsMiddleware).
		Build()

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./public")))

	mux.Handle("/api/", r)
	http.ListenAndServe(initializedConfiguration.GetServerAddress(), mux)
}
