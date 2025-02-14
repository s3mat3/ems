/**
 * @file login.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package controller

import (
    "strconv"
    "net/http"
    "ems/ss/adapter/presenter"
	"ems/ss/usecase/interactor"
	"ems/ss/usecase/port"
)

type Login struct {
    uc *interactor.Login
}
func NewLogin(uc *interactor.Login) *Login {
    return &Login{
        uc: uc,
    }
}

func (s *Login) ProcessLogin(c Context) {
    lr := port.LoginRequest{}
    if err := c.ShouldBindJSON(&lr); err != nil {
        c.JSON(http.StatusBadRequest, port.CreateMessage(err.Error()))
        return
    }
    p, err := s.uc.Authenticate(lr); if err != nil {
        c.JSON(http.StatusUnauthorized, port.CreateMessage("User not registerd"))
        return
    }
    t, err := s.uc.ProvideToken(p, "ssssssdddddffffdsaqedww3"); if err != nil {
        c.JSON(http.StatusUnauthorized, port.CreateMessage(err.Error()))
        return
    }
    ns := strconv.Itoa(p.GetID())
    c.Header("X-Temp-Range", ns)
    j := presenter.BuildLoginResponse("Authorized", t)
    c.JSON(http.StatusOK, j)
}
/* //<-- login.go ends here */
