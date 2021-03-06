package routes

import "taxi-tracker-api/api/service"

func TaxiServiceRoutes() []Route {
	ws := service.NewTaxiServiceWS()
	routes := []Route{
		{
			Method:  "POST",
			Path:    "/api/taxi_service/",
			Handler: ws.CreateService,
			Protected: true,
		},
		{
			Method:  "GET",
			Path:    "/api/taxi_service/",
			Handler: ws.GetCustomerService,
			Protected: true,
		},
	}

	return routes
}
