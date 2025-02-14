/**
 * @file hello.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package controller

import (
	"ems/ss/adapter/presenter"
	"ems/ss/usecase/port"
	"fmt"
	"net/http"
)

type Hello struct {
    uc port.HelloInput
    pr presenter.Hello
}

func NewHello(usecase port.HelloInput) *Hello {
    return &Hello{
        uc: usecase,
    }
}

func (s *Hello) ProcessGet(c Context) {
    dto, err := s.uc.ProvidePubkey()
    if err != nil {
        c.JSON(http.StatusInternalServerError, port.CreateMessage(err.Error()))
        return
    }
    c.JSON(http.StatusOK, dto)
}

func (s *Hello) ProcessPost(c Context) {
    sk := port.Himitsu{}
    if err := c.ShouldBindJSON(&sk); err != nil {
        fmt.Println(err.Error())
        c.JSON(http.StatusBadRequest, port.CreateMessage(err.Error()))
        return
    }
    dto, err := s.uc.ProvideSsid(sk)
    if err != nil {
        c.JSON(http.StatusInternalServerError, port.CreateMessage(err.Error()))
    }
    c.Header("X-Temp-Id", dto.Ssid)
    c.JSON(http.StatusOK, dto)
}

/* //<-- hello.go ends here */
