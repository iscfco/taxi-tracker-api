package routes

import "taxi-tracker-api/api/service"

func DriverSessionRoutes() []Route {
	ws := service.NewDriverSessionWS()
	routes := []Route{
		{
			Method:  "POST",
			Path:    "/api/driver_session/",
			Handler: ws.DriverLogin,
		},
	}

	return routes
}
