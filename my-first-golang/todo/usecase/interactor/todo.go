/**
 * @file todo.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package interactor

import (
    "todo/usecase/port"
     "todo/entity/repository"
)

type todoInteractor struct {
    oport port.TodoOutputPort
    repos repository.TodoRepository
}

func NewTodoUsecase(oport port.TodoOutputPort, repos repository.TodoRepository) port.TodoInputPort {
    return &todoInteractor{
        oport: oport,
        repos: repos,
    }
}

func (self *todoInteractor) GetHello() port.TodoMessageResponse {
    return self.oport.ResponseMessage("Hello, How are you?")
}

func (self *todoInteractor) GetAll() ([]port.TodoResponse, error) {
    l, err := self.repos.GetAll()
    if err != nil {
        return nil, err
    }
    return self.oport.ConvertEntityList(l)
}

func (self *todoInteractor) GetById(id string) (port.TodoResponse, error) {
    e, err := self.repos.FindById(id)
    if err != nil {
        return port.TodoResponse{}, err
    }
    return self.oport.ConvertEntity(e)
}

func (self *todoInteractor) Delete(id string) (port.TodoMessageResponse, error) {
    _, err := self.repos.Delete(id)
    if err != nil {
        return self.oport.ResponseMessage(err.Error()), err
    }
    return self.oport.ResponseMessage("ok"), nil
}


/* //<-- todo.go ends here */
