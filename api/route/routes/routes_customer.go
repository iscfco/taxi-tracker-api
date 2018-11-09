package routes

import "taxi-tracker-api/api/service"

func CustomerRoutes() []Route {
	ws := service.NewCustomerWS()
	routes := []Route{
		{
			Method:  "POST",
			Path:    "/api/customer",
			Handler: ws.CreateAccountHandler,
		},
	}

	return routes
}
