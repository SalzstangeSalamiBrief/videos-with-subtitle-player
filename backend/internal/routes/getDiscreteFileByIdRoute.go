package routes

import (
	"backend/internal/router"
	"backend/pkg/api/handlers"
	"net/http"
)

const getDiscreteFileByIdPath = `\/file\/discrete\/([0-9A-Fa-f]{8}[-]?[0-9A-Fa-f]{4}[-]?[0-9A-Fa-f]{4}[-]?[0-9A-Fa-f]{4}[-]?[0-9A-Fa-f]{12})$$`

func NewGetDiscreteFileByIdRoute(config handlers.DiscreteFileByIdHandlerConfig) router.Route {
	return router.Route{
		Path:    getDiscreteFileByIdPath,
		Method:  http.MethodGet,
		Handler: handlers.CreateDiscreteFileByIdHandler(config),
	}
}
