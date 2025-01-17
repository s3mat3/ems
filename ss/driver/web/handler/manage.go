/**
 * @file manage.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package handler

import (
    "github.com/gin-gonic/gin"
    "ems/ss/adapter/controller"
)

type ManageHandler struct {
    ctrl *controller.ManageController
}

func NewManageHandler(ctrl *controller.ManageController) *ManageHandler {
    return &ManageHandler{
        ctrl: ctrl,
    }
}

func (self *ManageHandler) Helth() gin.HandlerFunc {
    return func(c *gin.Context) {
        self.ctrl.Helth(c)
    }
}

func (self *ManageHandler) Hello() gin.HandlerFunc {
    return func(c *gin.Context) {
        self.ctrl.Hello(c)
    }
}



/* //<-- manage.go ends here */
