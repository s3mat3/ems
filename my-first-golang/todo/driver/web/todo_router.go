/**
 * @file todo_router.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package web

import(
    // "net/http"
    "github.com/gin-gonic/gin"
    "todo/driver/web/handler"
    "todo/adapter/controller"
    "todo/adapter/gateway"
    "todo/adapter/presenter"
    "todo/usecase/interactor"
)

func SetupRoute() *gin.Engine {
    repos      := gateway.NewTodoRepository("memory")
    presenter  := presenter.NewTodoPresenter()
    usecase    := interactor.NewTodoUsecase(presenter, repos)
    controller := controller.NewTodoController(usecase)
    handler    := handlers.NewTodoHandler(controller)
    router     := gin.Default()
    router.GET(    "/",         handler.Hello())
    router.GET(    "/todo",     handler.GetAll())
    router.GET(    "/todo/:id", handler.GetById())
    router.DELETE( "/todo/:id", handler.Delete())
    // router.POST(   "/todo",     handler.Post)
    // router.PUT(    "/todo",     handler.Update)
    return router
}
/* //<-- todo_router.go ends here */
