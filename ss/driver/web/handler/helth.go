/**
 * @file helth.go
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

type Helth struct {
    ctrl *controller.Helth
}

func NewHelth(ctrl *controller.Helth) *Helth {
    return &Helth{ctrl: ctrl}
}

func (s *Helth) Check() gin.HandlerFunc {
    return func(c *gin.Context) {
        s.ctrl.Check(c)
    }
}
/* //<-- helth.go ends here */
