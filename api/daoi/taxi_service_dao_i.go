package daoi

type TaxiServiceDaoI interface {
	CreateService(customerId *string) (string, error)
}
