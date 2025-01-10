/**
 * @file todo.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package port

import (
    "todo/entity/model"
)

type TodoInputPort interface {
    GetHello() TodoMessageResponse
    GetAll() ([]TodoResponse, error)
    GetById(id string) (TodoResponse, error)
    Delete(id string) (TodoMessageResponse, error)
}

// impl in adapter/presenter
type TodoOutputPort interface {
    ConvertEntity(model.Todo) (TodoResponse, error)
    ConvertEntityList([]model.Todo) ([]TodoResponse, error)
    ResponseMessage(msg string) TodoMessageResponse
}

/* //<-- todo.go ends here */
