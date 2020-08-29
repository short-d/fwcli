package routes

import (
	"fwcli/app/adapter/routes/handle"
	"net/http"

	"github.com/short-d/app/fw/router"
)

type Repo struct {
	Name string `json:"name"`
}

func NewRoutes() []router.Route {
	return []router.Route{
		{
			Method: http.MethodPost,
			Path:   "/services",
			Handle: handle.CreateService(),
		},
		{
			Method: http.MethodPut,
			Path:   "/services/:service/enable",
			Handle: handle.EnableService(),
		},
	}
}
