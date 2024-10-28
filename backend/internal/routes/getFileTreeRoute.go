package routes

import (
	"backend/internal/router"
	"backend/pkg/api/handlers"
	"net/http"
)

var GetFileTreeRoute = router.Route{
	Path:    "/file-tree",
	Handler: handlers.GetFileTreeHandler,
	Method:  http.MethodGet,
}
