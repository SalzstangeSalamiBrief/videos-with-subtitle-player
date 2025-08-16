package routes

import (
	"backend/internal/router"
	"backend/pkg/api/handlers"
	"net/http"
)

// TODO REFACTOR: CURRENTLY A HARD COUPLING BETWEEN THE HANDLER AND ROUTE FILE EXISTS => MAYBE MERGE?
func NewGetFileTreeRoute(configuration handlers.FileTreeHandlerConfiguration) router.Route {
	return router.Route{
		Path:    "/file-tree",
		Handler: handlers.CreateGetFileTreeHandler(configuration),
		Method:  http.MethodGet,
	}
}
