package routes

import "taxi-tracker-api/api/service"

func DriverRoutes() []Route {
	ws := service.NewDriverWS()
	routes := []Route{
		{
			Method:  "POST",
			Path:    "/api/driver",
			Handler: ws.CreateAccountHandler,
		},
	}

	return routes
}
