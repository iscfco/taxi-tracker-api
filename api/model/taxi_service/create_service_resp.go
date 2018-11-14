package taxi_service

import "taxi-tracker-api/api/model"

type CreateServiceResp struct {
	Result    model.Result `json:"result"`
	VehicleId string       `json:"vehicle_id"`
}
