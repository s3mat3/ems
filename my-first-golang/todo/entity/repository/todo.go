/**
 * @file todo.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package repository

import (
    "todo/entity/model"
)

type TodoRepository interface {
    GetAll() ([]model.Todo, error)
    FindById(id string) (model.Todo, error)
    Delete(id string) (int, error)
    // Insert(todo entity.TodoEntity) error
    // Update(todo entity.TodoEntity) error
}



/* //<-- todo.go ends here */
