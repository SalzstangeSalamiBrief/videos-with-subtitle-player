package main

import (
	"backend/internal/config/api"
	"backend/internal/config/oicd"
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
	initializedApiConfiguration, configurationError := api.NewApiConfiguration()
	if configurationError != nil {
		log.Fatal(configurationError)
	}

	initializedKeycloakConfiguration, configurationError := oicd.NewKeycloakConfiguration()
	if configurationError != nil {
		log.Fatal(configurationError)
	}

	// TODO
	log.Printf("Initialized Keycloak Configuration: %+v", initializedKeycloakConfiguration)

	initializedFileTreeManager := fileTreeManager.NewFileTreeManager(initializedApiConfiguration.GetRootPath()).InitializeTree()

	log.Default().Printf("Start server on '%v'", initializedApiConfiguration.GetServerAddress())

	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, os.Interrupt, syscall.SIGTERM)

	corsMiddleware := middlewares.NewCorsMiddleWare().AddConfiguration(middlewares.CorsMiddleWareConfiguration{AllowedCors: initializedApiConfiguration.GetCors()}).Build()
	requestLoggerMiddleware := middlewares.NewRequestLogger().Build()

	handleDiscreteFileRoute := routes.NewGetDiscreteFileByIdRoute(handlers.DiscreteFileByIdHandlerConfig{
		RootPath:        initializedApiConfiguration.GetRootPath(),
		FileTreeManager: initializedFileTreeManager,
	})
	handleContinousFileRoute := routes.CreateGetContinuousFileRoute(handlers.ContinuousFileByIdHandlerConfiguration{RootPath: initializedApiConfiguration.GetRootPath(), FileTreeManager: initializedFileTreeManager})
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
	http.ListenAndServe(initializedApiConfiguration.GetServerAddress(), mux)
}
