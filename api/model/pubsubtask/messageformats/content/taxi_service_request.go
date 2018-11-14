package content

type TaxiServiceRequest struct {
	UserName      string  `json:"user_name"`
	UserLatitude  float32 `json:"user_latitude"`
	UserLongitude float32 `json:"use_longitude"`
}
