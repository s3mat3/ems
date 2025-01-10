/**
 * @file todo.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package presenter

import (
    "todo/entity/model"
    "todo/usecase/port"
)


type TodoPresenter struct {
}


func NewTodoPresenter() port.TodoOutputPort {
    return &TodoPresenter{}
}

func (self *TodoPresenter) ConvertEntity(entity model.Todo) (port.TodoResponse, error) {
    return port.TodoResponse{
        Id: entity.GetID(),
        Subject: entity.GetSubject(),
        Status: model.GetStatusName(entity.GetStatus()),
    }, nil
}
func (self *TodoPresenter) ConvertEntityList(entities []model.Todo) ([]port.TodoResponse, error) {
    var list []port.TodoResponse
    for _, v := range entities {
        t, err := self.ConvertEntity(v)
        if err != nil {
            return nil, err
        }
        list = append(list, t)
    }
    return list, nil
}

func (self *TodoPresenter) ResponseMessage(msg string) port.TodoMessageResponse {
    return port.TodoMessageResponse {
        Message:  msg,
    }
}

// func (self *TodoPresenter) GenTodoList() ([]usecases.TodoOutputData,error) {
//     return nil, nil
// }
/* //<-- todo.go ends here */
