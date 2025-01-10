/**
 * @file todo.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package handlers

import (
    "github.com/gin-gonic/gin"
    "todo/adapter/controller"
)

type TodoHandler struct {
    ctrl *controller.TodoController
}

func NewTodoHandler(ctrl *controller.TodoController) *TodoHandler {
    return &TodoHandler{
        ctrl: ctrl,
    }
}

func (self *TodoHandler) Hello() gin.HandlerFunc {
    return func(c *gin.Context) {
        self.ctrl.Hello(c)
    }
}
func (self *TodoHandler) GetAll() gin.HandlerFunc {
    return func(c *gin.Context) {
        self.ctrl.GetAll(c)
    }
}
func (self *TodoHandler) GetById() gin.HandlerFunc {
    return func(c *gin.Context) {
        self.ctrl.GetById(c)
    }
}
func (self *TodoHandler) Delete() gin.HandlerFunc {
    return func(c *gin.Context) {
        self.ctrl.Delete(c)
    }
}
/* //<-- todo.go ends here */
