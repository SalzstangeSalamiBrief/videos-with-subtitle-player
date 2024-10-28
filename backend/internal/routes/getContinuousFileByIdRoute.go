package routes

import (
	"backend/internal/router"
	"backend/pkg/api/handlers"
	"net/http"
)

const getContinuousFilePath = `\/file\/continuous\/([0-9A-Fa-f]{8}[-]?[0-9A-Fa-f]{4}[-]?[0-9A-Fa-f]{4}[-]?[0-9A-Fa-f]{4}[-]?[0-9A-Fa-f]{12})$$`

var GetContinuousFileRoute = router.Route{
	Path:    getContinuousFilePath,
	Method:  http.MethodGet,
	Handler: handlers.GetContinuousFileByIdHandler,
}
