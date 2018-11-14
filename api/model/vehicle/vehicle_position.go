package vehicle

type VehiclePosition struct {
	VehicleId string `json:"vehicle_id"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}
