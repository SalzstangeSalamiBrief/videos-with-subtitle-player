package main

import (
	"backend/internal/config"
	"backend/internal/router"
	"backend/internal/routes"
	"backend/pkg/api/handlers"
	"backend/pkg/api/middlewares"
	"backend/pkg/services/fileTreeManager"
	"backend/pkg/services/imageConverter/webp"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	initializedConfiguration := config.InitializeConfiguration()
	err := webp.ExecuteWebpConversion(webp.ExecuteWebpConversionConfiguration{RootPath: initializedConfiguration.RootPath})
	if err != nil {
		log.Fatal(err)
	}

	//initializedImageHandler := imageHandlerSources.NewMagickImageHandler(imageConverter.LowQualityFileSuffix)
	initializedFileTreeManager := fileTreeManager.NewFileTreeManager(initializedConfiguration.RootPath).InitializeTree()

	log.Default().Printf("Start server on '%v'", initializedConfiguration.ServerAddress)

	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
		<-exit
		log.Default().Printf("Shutting down server at %v\n", initializedConfiguration.ServerAddress)
		os.Exit(0)
	}()

	corsMiddleware := middlewares.NewCorsMiddleWare().AddConfiguration(middlewares.CorsMiddleWareConfiguration{AllowedCors: initializedConfiguration.AllowedCors}).Build()
	requestLoggerMiddleware := middlewares.NewRequestLogger().Build()

	handleDiscreteFileRoute := routes.NewGetDiscreteFileByIdRoute(handlers.DiscreteFileByIdHandlerConfig{
		RootPath:        initializedConfiguration.RootPath,
		FileTreeManager: initializedFileTreeManager,
	})
	handleContinousFileRoute := routes.CreateGetContinuousFileRoute(handlers.ContinuousFileByIdHandlerConfiguration{RootPath: initializedConfiguration.RootPath, FileTreeManager: initializedFileTreeManager})
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
	http.ListenAndServe(initializedConfiguration.ServerAddress, mux)
}
