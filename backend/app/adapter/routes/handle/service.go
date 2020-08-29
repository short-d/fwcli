package handle

import (
	"fwcli/app/usecase/service"
	"net/http"

	"github.com/short-d/app/fw/router"
)

func CreateService(
	serviceController *service.Controller,
) router.Handle {
	return func(w http.ResponseWriter, r *http.Request, params router.Params) {
		// Create Github repo
		serviceController.CreateService()
	}
}

func EnableService(
	serviceController *service.Controller,
) router.Handle {
	return func(w http.ResponseWriter, r *http.Request, params router.Params) {
		serviceController.EnableService()
	}
}
