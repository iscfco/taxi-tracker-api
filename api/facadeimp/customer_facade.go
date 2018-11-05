package facadeimp

import (
	"gbmchallenge/api/daoi"
	"gbmchallenge/api/facadei"
	"gbmchallenge/api/model"
	"gbmchallenge/api/security"
	"gbmchallenge/api/util"
	"github.com/satori/go.uuid"
	"log"
)

type customerFacade struct {
	customerDao daoi.CustomerDaoI
}

func NewCustomerFacade(c daoi.CustomerDaoI) facadei.CustomerFacadeI {
	return &customerFacade{
		customerDao: c,
	}
}

func (c *customerFacade) CreateAccount(customer *model.Customer) model.Result {
	pwdHashed, err := security.HashPassword(customer.Password)
	if err != nil {
		// handle
		log.Print(err)
		return util.GetServerErr()
	}

	customerId, _ := uuid.NewV4()
	customer.Id = customerId.String()
	customer.Password = pwdHashed
	res, err := c.customerDao.CreateAccount(customer)
	if err != nil {
		// handle
		log.Print(err)
		return util.GetServerErr()
	}
	return res
}
