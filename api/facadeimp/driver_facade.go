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

type driverFacade struct {
	dao daoi.DriverDaoI
}

func NewDriverFacade(d daoi.DriverDaoI) facadei.DriverFacadeI {
	return &driverFacade{
		dao: d,
	}
}

func (c *driverFacade) CreateAccount(driver *model.Driver) model.Result {
	pwdHashed, err := security.HashPassword(driver.Password)
	if err != nil {
		// handle
		log.Print(err)
		return util.GetServerErr()
	}

	driverId, _ := uuid.NewV4()
	driver.Id = driverId.String()
	driver.Password = pwdHashed
	res, err := c.dao.CreateAccount(driver)
	if err != nil {
		// handle
		log.Print(err)
		return util.GetServerErr()
	}
	return res
}
