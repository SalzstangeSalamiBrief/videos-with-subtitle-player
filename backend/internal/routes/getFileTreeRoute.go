package routes

import (
	"backend/internal/router"
	"backend/pkg/api/handlers"
	"net/http"
)

func NewGetFileTreeRoute(configuration handlers.FileTreeHandlerConfiguration) router.Route {
	return router.Route{
		Path:    "/file-tree",
		Handler: handlers.CreateGetFileTreeHandler(configuration),
		Method:  http.MethodGet,
	}
}
