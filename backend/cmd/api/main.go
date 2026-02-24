package main

import (
	"backend/internal/configuration"
	"backend/internal/database"
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
	apiConfiguration, _ := loadConfigurations()

	initializedFileTreeManager := fileTreeManager.NewFileTreeManager(apiConfiguration.GetRootPath()).InitializeTree()
	//dbConnection, createDbError := createDatabases(dbConfiguration)
	//if createDbError != nil {
	//	if dbConnection != nil {
	//		dbConnection.Close()
	//	}
	//
	//	log.Fatal(createDbError)
	//}
	log.Default().Printf("Start server on '%v'", apiConfiguration.GetServerAddress())

	// TODO CLEANUP OF DB
	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, os.Interrupt, syscall.SIGTERM)

	routerBuilder := router.NewRouterBuilder()

	createdRoutes := createRoutes(apiConfiguration, initializedFileTreeManager)
	for _, route := range createdRoutes {
		routerBuilder.RegisterRoute(route)
	}

	createdMiddlewares := createMiddlewares(apiConfiguration)
	for _, middleware := range createdMiddlewares {
		routerBuilder.RegisterMiddleware(middleware)
	}

	r := routerBuilder.Build()

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./public")))

	mux.Handle("/api/", r)
	http.ListenAndServe(apiConfiguration.GetServerAddress(), mux)
}

func loadConfigurations() (*configuration.ApiConfiguration, *configuration.DbConfiguration) {
	apiConfiguration, apiConfigurationError := configuration.NewApiConfiguration()
	if apiConfigurationError != nil {
		log.Fatal(apiConfigurationError)
	}

	dbConfiguration, dbConfigurationError := configuration.NewDbConfiguration()
	if dbConfigurationError != nil {
		log.Fatal(dbConfigurationError)
	}

	return apiConfiguration, dbConfiguration
}

func createDatabases(dbConfiguration *configuration.DbConfiguration) (*database.FileTreeDatabase, error) {
	fileTreeDb := database.NewFileTreeDatabase().AddConfiguration(dbConfiguration)
	_, migrateDbError := fileTreeDb.MigrateDatabase()
	if migrateDbError != nil {
		log.Fatal(migrateDbError)
	}

	return fileTreeDb, nil
}

func createRoutes(apiConfiguration *configuration.ApiConfiguration, ftm *fileTreeManager.FileTreeManager) []router.Route {
	handleDiscreteFileRoute := routes.NewGetDiscreteFileByIdRoute(handlers.DiscreteFileByIdHandlerConfig{
		RootPath:        apiConfiguration.GetRootPath(),
		FileTreeManager: ftm,
	})
	handleContinousFileRoute := routes.CreateGetContinuousFileRoute(handlers.ContinuousFileByIdHandlerConfiguration{RootPath: apiConfiguration.GetRootPath(), FileTreeManager: ftm})
	getFileTreeRoute := routes.NewGetFileTreeRoute(handlers.FileTreeHandlerConfiguration{FileTreeManager: ftm})

	return []router.Route{
		handleDiscreteFileRoute,
		handleContinousFileRoute,
		getFileTreeRoute,
	}
}

func createMiddlewares(apiConfiguration *configuration.ApiConfiguration) []router.Middleware {
	corsMiddleware := middlewares.NewCorsMiddleWare().AddConfiguration(middlewares.CorsMiddleWareConfiguration{AllowedCors: apiConfiguration.GetCors()}).Build()
	requestLoggerMiddleware := middlewares.NewRequestLogger().Build()

	return []router.Middleware{
		corsMiddleware,
		requestLoggerMiddleware,
	}
}
