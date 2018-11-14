package facadeimp

import (
	"taxi-tracker-api/api/constants"
	"taxi-tracker-api/api/daoi"
	"taxi-tracker-api/api/errorhandler"
	"taxi-tracker-api/api/facadei"
	"taxi-tracker-api/api/model"
	"taxi-tracker-api/api/model/driver"
	"taxi-tracker-api/api/security"
	"github.com/satori/go.uuid"
)

type driverFacade struct {
	dao daoi.DriverDaoI
}

func NewDriverFacade(d daoi.DriverDaoI) facadei.DriverFacadeI {
	return &driverFacade{
		dao: d,
	}
}

func (c *driverFacade) CreateAccount(driver *driver.Driver) model.Result {
	if len(driver.Password) < 8 {
		return model.Result{
			ResCode:  constants.EDV001_C,
			Msg:      constants.EDV001_M,
			HttpCode: 200,
		}
	}

	pwdHashed, err := security.HashPassword(driver.Password)
	if err != nil {
		return errorhandler.HandleErr(&err)
	}

	driverId, _ := uuid.NewV4()
	driver.Id = driverId.String()
	driver.Password = pwdHashed
	res, err := c.dao.CreateAccount(driver)
	if err != nil {
		return errorhandler.HandleErr(&err)
	}
	return res
}
