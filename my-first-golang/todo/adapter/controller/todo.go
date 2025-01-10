/**
 * @file todo.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package controller

import (
    "net/http"
    "todo/adapter/presenter"
    "todo/usecase/port"
)

// This idea is from https://github.com/arakawamoriyuki/go-clean-handson,
// so the framework will not leak than the infrastructure.
//
type Context interface {
    Abort()
    BindJSON(obj interface{}) error
    JSON(code int, obj interface{})
    Param(ket string) string
}

type TodoController struct {
    uc port.TodoInputPort
}

func NewTodoController(usecase port.TodoInputPort)  *TodoController {
    return &TodoController{
        uc: usecase,
    }
}

func (self *TodoController) Hello(c Context) {
    c.JSON(http.StatusOK, self.uc.GetHello())
}

func (self *TodoController) GetAll(c Context) {
    l, err := self.uc.GetAll()
    if err != nil {
        c.JSON(http.StatusNotFound, presenter.NewErrorResponse("Not found"))
        return
    }
    c.JSON(http.StatusOK, l)
}

func (self *TodoController) GetById(c Context) {
    id := c.Param("id")
    d, err := self.uc.GetById(id)
    if err != nil {
        c.JSON(http.StatusNotFound, presenter.NewErrorResponse("Not found"))
    } else {
        c.JSON(http.StatusOK, d)
    }
}

func (self *TodoController) Delete(c Context) {
    id := c.Param("id")
    d, err := self.uc.Delete(id)
    if err != nil {
        c.JSON(http.StatusNotFound, presenter.NewErrorResponse("Can not delete, cause key not found"))
    } else {
        c.JSON(http.StatusOK, d)
    }
}
// func (self *TodoController) GetByStatusTodoController(c *gin.Context) {
//     c.JSON(http.StatusOK, gin.H{"message": "Get 1 todo by status"})
// }

// func (self *TodoController) PostTodoController(c *gin.Context) {
//     c.JSON(http.StatusOK, gin.H{"message": "Post new todo"})
// }

// func (self *TodoController) UpdateTodoController(c *gin.Context) {
//     c.JSON(http.StatusOK, gin.H{"message": "Update to already exist todo"})
// }


/* //<-- todo.go ends here */
