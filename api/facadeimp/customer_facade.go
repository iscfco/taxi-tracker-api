package facadeimp

import (
	"gbmchallenge/api/constants"
	"gbmchallenge/api/daoi"
	"gbmchallenge/api/errorhandler"
	"gbmchallenge/api/facadei"
	"gbmchallenge/api/model"
	"gbmchallenge/api/security"
	"github.com/satori/go.uuid"
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
	if len(customer.Password) < 8 {
		return model.Result{
			ResCode:  constants.EDV001_C,
			Msg:      constants.EDV001_M,
			HttpCode: 200,
		}
	}

	pwdHashed, err := security.HashPassword(customer.Password)
	if err != nil {
		return errorhandler.HandleErr(&err)
	}

	customerId, _ := uuid.NewV4()
	customer.Id = customerId.String()
	customer.Password = pwdHashed
	res, err := c.customerDao.CreateAccount(customer)
	if err != nil {
		return errorhandler.HandleErr(&err)
	}
	return res
}
