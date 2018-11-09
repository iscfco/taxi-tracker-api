package routes

import "taxi-tracker-api/api/service"

func VehicleRoutes() []Route {
	ws := service.NewVehicleWS()
	routes := []Route{
		{
			Method:  "GET",
			Path:    "/api/vehicle",
			Handler: ws.GetVehiclesHandler,
		},
		{
			Method:    "GET",
			Path:      "/api/vehicle/{vehicleId}/position",
			Handler:   ws.GetVehiclePositionHandler,
			Protected: true,
		},
		{
			Method:    "PATCH",
			Path:      "/api/vehicle/{vehicleId}/position",
			Handler:   ws.UpdateVehiclePositionHandler,
			Protected: true,
		},
	}

	return routes
}
