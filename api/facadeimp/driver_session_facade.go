package facadeimp

import (
	"gbmchallenge/api/constants"
	"gbmchallenge/api/daoi"
	"gbmchallenge/api/errorhandler"
	"gbmchallenge/api/facadei"
	"gbmchallenge/api/model"
	"gbmchallenge/api/response/prebuilt"
	"gbmchallenge/api/security"
	"gbmchallenge/api/security/jwttasks"
)

type driverSessionFacade struct {
	driverDao daoi.DriverDaoI
}

func NewDriverSessionFacade(dao daoi.DriverDaoI) facadei.DriverSessionFacadeI {
	return &driverSessionFacade{
		driverDao: dao,
	}
}

func (c *driverSessionFacade) Authorize(user *model.User) (s model.DriverSession) {
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
