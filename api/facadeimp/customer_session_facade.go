package facadeimp

import (
	"gbmchallenge/api/constants"
	"gbmchallenge/api/daoi"
	"gbmchallenge/api/errorhandler"
	"gbmchallenge/api/facadei"
	"gbmchallenge/api/model"
	"gbmchallenge/api/security"
	"gbmchallenge/api/security/jwttasks"
)

type customerSessionFacade struct {
	customerDao daoi.CustomerDaoI
}

func NewCustomerSessionFacade(dao daoi.CustomerDaoI) facadei.CustomerSessionFacadeI {
	return &customerSessionFacade{
		customerDao: dao,
	}
}

func (c *customerSessionFacade) Authorize(user *model.User) (s model.Session) {
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
	s.Res.ResCode, s.Res.Msg, s.Res.HttpCode = constants.SUCCESS_C, constants.SUCCESS_M, 200
	return
}
