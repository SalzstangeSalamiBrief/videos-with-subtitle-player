package routes

import (
	"backend/internal/router"
	"backend/pkg/api/handlers"
	"net/http"
)

const getContinuousFileByIdPath = `\/file\/continuous\/([0-9A-Fa-f]{8}[-]?[0-9A-Fa-f]{4}[-]?[0-9A-Fa-f]{4}[-]?[0-9A-Fa-f]{4}[-]?[0-9A-Fa-f]{12})$$`

func CreateGetContinuousFileRoute(configuration handlers.ContinuousFileByIdHandlerConfiguration) router.Route {
	return router.Route{
		Path:    getContinuousFileByIdPath,
		Method:  http.MethodGet,
		Handler: handlers.NewGetContinuousFileByIdHandler(configuration),
	}
}
