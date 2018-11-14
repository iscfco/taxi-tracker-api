package facadeimp

import (
	"taxi-tracker-api/api/constants"
	"taxi-tracker-api/api/daoi"
	"taxi-tracker-api/api/errorhandler"
	"taxi-tracker-api/api/facadei"
	"taxi-tracker-api/api/model"
	"taxi-tracker-api/api/model/customer"
	"taxi-tracker-api/api/response/prebuilt"
	"taxi-tracker-api/api/security"
	"taxi-tracker-api/api/security/jwttasks"
)

type customerSessionFacade struct {
	customerDao daoi.CustomerDaoI
}

func NewCustomerSessionFacade(dao daoi.CustomerDaoI) facadei.CustomerSessionFacadeI {
	return &customerSessionFacade{
		customerDao: dao,
	}
}

func (c *customerSessionFacade) Authorize(user *model.User) (s customer.CustomerSession) {
	err, customer := c.customerDao.GetByEmail(&user.User)
	if err != nil {
		s.Res = errorhandler.HandleErr(&err)
		return
	}

	err, match := security.CheckPasswordHash(&customer.Password, &user.Password)
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

	s.AccessToken, err = jwtTasks.GenerateAccessToken(&customer.Id)
	if err != nil {
		s.Res = errorhandler.HandleErr(&err)
		return
	}

	s.RefreshToken, err = jwtTasks.GenerateRefreshToken(&customer.Id)
	if err != nil {
		s.Res = errorhandler.HandleErr(&err)
		return
	}

	s.FirstName, s.LastName = customer.FirstName, customer.LastName
	s.Res = prebuilt.GetSuccess()
	return
}
