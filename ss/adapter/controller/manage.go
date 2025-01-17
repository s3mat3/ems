/**
 * @file manage.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package controller

import (
    "net/http"

    "ems/ss/adapter/presenter"
)

type response struct {
    Message string `json:"message"`
}

type token struct {
    Id string `json:"id"`
}

type ManageController struct {
    // member1 usecase/input port interface
}

func NewManageController(/* usecase port.xxxx */)  *ManageController {
    return &ManageController{
        //uc: usecase,
    }
}

func (self *ManageController) Helth(c Context) {
    c.JSON(http.StatusBadRequest, presenter.H{"message": "hello"})
}

func (self *ManageController) Hello(c Context) {
    c.Header("Authorization", "{\"id\": 12345}")
    c.JSON(http.StatusOK, presenter.H{"message": "hello"})
}



/* //<-- manage.go ends here */
