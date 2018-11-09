package daoi

type VehicleDriverDaoI interface {
	Create(vehicleId, driverId *string) (err error)
}
