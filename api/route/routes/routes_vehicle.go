package routes

import "taxi-tracker-api/api/service"

func VehicleRoutes() []Route {
	ws := service.NewVehicleWS()
	routes := []Route{
		{
			Method:    "GET",
			Path:      "/api/vehicle",
			Handler:   ws.GetVehiclesHandler,
			Protected: true,
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
		{
			Method:    "PATCH",
			Path:      "/api/v2/vehicle/{vehicleId}/position",
			Handler:   ws.UpdateVehiclePositionHandlerV2,
			Protected: true,
		},
	}

	return routes
}
