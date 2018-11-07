package route

import (
	r "gbmchallenge/api/route/routes"
	"github.com/gorilla/mux"
)

func CreateRoutes(router *mux.Router) {
	var routes []r.Route
	routes = append(routes, r.VehicleRoutes()...)
	routes = append(routes, r.CustomerRoutes()...)
	routes = append(routes, r.DriverRoutes()...)

	for _, route := range routes {
		router.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}
}
