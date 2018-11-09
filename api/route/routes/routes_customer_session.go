package routes

import "taxi-tracker-api/api/service"

func CustomerSessionRoutes() []Route {
	ws := service.NewCustomerSessionWS()
	routes := []Route{
		{
			Method:  "POST",
			Path:    "/api/customer_session/",
			Handler: ws.CustomerLogin,
		},
	}

	return routes
}
