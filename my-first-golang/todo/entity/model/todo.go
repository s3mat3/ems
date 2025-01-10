/**
 * @file todo_entity.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package model

type Todo struct {
    Id      int
    Subject string
    Status  int
}

func NewTodo() *Todo {
    return &Todo{}
}

func (ent *Todo) Entry(i int, s string, sts int) {
    ent.Id = i
    ent.Subject = s
    ent.Status = sts
}

func (ent *Todo) GetID() int {
    return ent.Id
}
func (ent *Todo) GetSubject() string {
    return ent.Subject
}
func (ent *Todo) GetStatus() int {
    return ent.Status
}

/* //<-- todo_entity.go ends here */
