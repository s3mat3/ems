/**
 * @file hello.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package handler

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "ems/ss/adapter/controller"
)

type Hello struct {
    ctrl *controller.Hello
}

func NewHello(ctrl *controller.Hello) *Hello {
    return &Hello{
        ctrl: ctrl,
    }
}

func (self *Hello) GET() gin.HandlerFunc {
    fmt.Println("Hello")
    return func(c *gin.Context) {
        self.ctrl.ProcessGet(c)
    }
}

func (self *Hello) POST() gin.HandlerFunc {
    return func(c *gin.Context) {
        self.ctrl.ProcessPost(c)
    }
}

/* //<-- hello.go ends here */
