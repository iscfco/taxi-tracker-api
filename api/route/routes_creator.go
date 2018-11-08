package route

import (
	"gbmchallenge/api/middleware"
	r "gbmchallenge/api/route/routes"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func CreateRoutes(router *mux.Router) {
	var routes []r.Route
	routes = append(routes, r.VehicleRoutes()...)
	routes = append(routes, r.CustomerRoutes()...)
	routes = append(routes, r.DriverRoutes()...)
	routes = append(routes, r.CustomerSessionRoutes()...)

	for _, route := range routes {
		if route.Protected {
			router.Handle(route.Path,
				negroni.New(
					negroni.HandlerFunc(middleware.ValidateToken),
					negroni.WrapFunc(route.Handler),
				),
			).Methods(route.Method)
			continue
		}
		router.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}
}
