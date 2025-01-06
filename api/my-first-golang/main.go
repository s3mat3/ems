package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
//    "github.com/jmoiron/sqlx"
)


type todo struct {
    ID string `json: "id"`
    Subject string `json: "subject"`
    Status string `json: "status"`
}

var todoList = []todo {
    {ID: "1", Subject: "hoge", Status: "new"},
    {ID: "2", Subject: "fuga", Status: "open"},
    {ID: "3", Subject: "sleep", Status: "in progress"},
    {ID: "4", Subject: "wakeup", Status: "done"},
    {ID: "5", Subject: "eat", Status: "in progress"},
}

func readTodoAll(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, todoList)
}

func readTodoByID(c *gin.Context) {
    id := c.Param("id")
    for _, v := range todoList {
        if v.ID == id {
            c.IndentedJSON(http.StatusOK, v)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func appendTodo(c *gin.Context, targetTodo todo) {
    todoList = append(todoList, targetTodo)
    c.IndentedJSON(http.StatusCreated, targetTodo)
}


func createTodo(c *gin.Context) {
    var newTodo todo
    if err := c.BindJSON(&newTodo); err != nil {
        return
    }
    for _, v := range todoList {
        if v.ID == newTodo.ID {
            c.IndentedJSON(http.StatusConflict, gin.H{
                "status"  : "insert error",
                "source" : v,
                "target" : newTodo,
            })
            return
        }
    }
    appendTodo(c, newTodo)
}

func updateTodo(c * gin.Context) {
    var targetTodo todo
    ref := todoList
    if err := c.BindJSON(&targetTodo); err != nil {
        return
    }
    for k, _ := range todoList {
        // found same ID -> update
        if ref[k].ID == targetTodo.ID {
            ref[k].Subject = targetTodo.Subject
            ref[k].Status = targetTodo.Status
            c.IndentedJSON(http.StatusNoContent, nil)
            return
        }
    }
    // never found insert
    appendTodo(c, targetTodo)
}

// func remove(list []todo, id string) []todo {
//     temp := []todo{}
//     for _, v := range list {
//         if v.ID != id {
//             temp = append(temp, v)
//         }
//     }
//     return temp;
// }

func deleteTodo(c *gin.Context) {
    var targetTodo todo
    if err := c.BindJSON(&targetTodo); err != nil {
        return
    }
    temp := []todo{}
    flg := false
    for _, v := range todoList {
        if v.ID != targetTodo.ID {
            temp = append(temp, v)
        } else {
            flg = true
        }
    }
    if flg {
        todoList = temp
        c.IndentedJSON(http.StatusNoContent, nil)
    } else {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "delte target not found"})
    }
}

func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "I'm alive and well!",
        })
    })
    r.GET("/todo", readTodoAll) // read all todoList
    r.GET("/todo/:id", readTodoByID) // read all todoList
    r.DELETE("/todo", deleteTodo) // update or create
    r.POST("/todo", createTodo) // create
    r.PUT("/todo", updateTodo) // update or create
    // serveer start
    r.Run(":20080")
}
