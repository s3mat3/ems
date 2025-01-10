/**
 * @file memory.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package memory

import (
    _"fmt"
    "errors"
    "strconv"
    "strings"
    "todo/driver/persistent"
    "todo/entity/model"
)

var inMemory = []model.Todo {
    {Id: 1, Subject: "Herohero", Status: 0},
    {Id: 2, Subject: "Herohero", Status: 1},
    {Id: 3, Subject: "Herohero", Status: 2},
    {Id: 4, Subject: "Herohero", Status: 2},
    {Id: 5, Subject: "Herohero", Status: -1},
}

type memoryFunction struct {}

func NewMemory() persistent.Functional {
    return &memoryFunction{}
}
// Serch() select all
// Serch("id = 1") select one record by id is 1 (cf. sql select * from x where id = 1)
func (self *memoryFunction) Search(cond ...string) ([]model.Todo, error) {
    if cond == nil {
        return inMemory, nil
    } else {
        if strings.Contains(cond[0], "id") {
            s := strings.SplitAfter(cond[0], "=")
            if len(s) < 2 {
                return nil, errors.New("Need id = xx")
            }
            sn := strings.TrimSpace(s[1])
            id, err := strconv.Atoi(sn)
            if err != nil { // sn not a number string
                return nil, err
            }
            var l []model.Todo
            for _, v := range inMemory {
                if v.Id == id {
                    l = append(l, v)
                    return l, nil
                }
            }
            // no id
            return nil, errors.New("no id")
        } else { // has not search condition
            return nil, errors.New("no condition id = ")
        }
    }
    // maybe unreachable
    // return nil, nil
}

func (self *memoryFunction) Delete(id string) (int, error) {
    key, err := strconv.Atoi(id)
    if err != nil { // sn not a number string
        return -1, err
    }
    flg := false
    tmp := []model.Todo{}
    for _, v := range inMemory {
        if v.Id != key {
            tmp = append(tmp, v)
        } else {
            flg = true
        }
    }
    if flg {
        inMemory = tmp
        return key, nil
    }
    return -1, errors.New("Not found target id")
}
/* //<-- memory.go ends here */
