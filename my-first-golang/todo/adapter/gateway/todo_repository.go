/**
 * @file todo_repository.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package gateway

import(
    "todo/driver/persistent"
    "todo/driver/persistent/memory"
    "todo/entity/model"
    "todo/entity/repository"
)

type todoRepository struct {
    f persistent.Functional
}

func NewTodoRepository(selector string) repository.TodoRepository {
    var pf persistent.Functional
    switch selector {
    case "memory" : pf = memory.NewMemory()
    default: pf = memory.NewMemory()
    }
    return &todoRepository{pf}
}

func (self *todoRepository) GetAll() ([]model.Todo, error) {
    return self.f.Search()
}

func (self *todoRepository) FindById(i string) (model.Todo, error) {
    id := "id = " + i
    l, err := self.f.Search(id)
    if err != nil {
        return model.Todo{}, err
    }
    return l[0], nil
}

func (self *todoRepository) Delete(i string) (int, error) {
    l, err := self.f.Delete(i)
    if err != nil {
        return -1, err
    }
    return l, nil
}
/* //<-- todo_repository.go ends here */
