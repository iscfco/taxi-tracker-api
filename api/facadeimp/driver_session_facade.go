package facadeimp

import (
	"taxi-tracker-api/api/constants"
	"taxi-tracker-api/api/daoi"
	"taxi-tracker-api/api/errorhandler"
	"taxi-tracker-api/api/facadei"
	"taxi-tracker-api/api/model"
	"taxi-tracker-api/api/model/driver"
	"taxi-tracker-api/api/response/prebuilt"
	"taxi-tracker-api/api/security"
	"taxi-tracker-api/api/security/jwttasks"
)

type driverSessionFacade struct {
	driverDao daoi.DriverDaoI
}

func NewDriverSessionFacade(dao daoi.DriverDaoI) facadei.DriverSessionFacadeI {
	return &driverSessionFacade{
		driverDao: dao,
	}
}

func (c *driverSessionFacade) Authorize(user *model.User) (s driver.DriverSession) {
	err, driver := c.driverDao.GetByEmail(&user.User)
	if err != nil {
		s.Res = errorhandler.HandleErr(&err)
		return
	}

	err, match := security.CheckPasswordHash(&driver.Password, &user.Password)
	if err != nil {
		s.Res = errorhandler.HandleErr(&err)
		return
	} else if !match {
		s.Res.ResCode, s.Res.Msg, s.Res.HttpCode = constants.EUS002_C, constants.EUS002_M, 200
		return
	}

	err, jwtTasks := jwttasks.NewJwtTasks()
	if err != nil {
		s.Res = errorhandler.HandleErr(&err)
		return
	}

	s.AccessToken, err = jwtTasks.GenerateAccessToken(&driver.Id)
	if err != nil {
		s.Res = errorhandler.HandleErr(&err)
		return
	}

	s.RefreshToken, err = jwtTasks.GenerateRefreshToken(&driver.Id)
	if err != nil {
		s.Res = errorhandler.HandleErr(&err)
		return
	}

	s.FirstName, s.LastName = driver.FirstName, driver.LastName
	s.Res = prebuilt.GetSuccess()
	return
}
