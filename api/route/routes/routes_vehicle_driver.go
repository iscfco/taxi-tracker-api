package routes

import "taxi-tracker-api/api/service"

func VehicleDriverRoutes() []Route {
	ws := service.NewVehicleDriverWS()
	routes := []Route{
		{
			Method:    "POST",
			Path:      "/api/vehicle_driver/{vehicleId}",
			Handler:   ws.Create,
			Protected: true,
		},
	}

	return routes
}
