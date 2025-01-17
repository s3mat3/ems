/**
 * @file router.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package web


import(
    "time"
    
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"

    "ems/ss/driver/web/handler"
    "ems/ss/adapter/controller"
)
//
func SetupRoute() *gin.Engine {
    controller := controller.NewManageController()
    handler    := handler.NewManageHandler(controller)

    router     := gin.Default()
    // setup for cors
    router.Use(cors.New(cors.Config{
        AllowMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE",},
        AllowHeaders: []string{
            "Access-Control-Allow-Headers",
            "Content-Type",
            "Content-Length",
            "Accept-Encoding",
            "X-CSRF-Token",
            "Authorization",
        },
        AllowOrigins: []string{"http://localhost:28080"},
        MaxAge: 24 * time.Hour,
    }))


    router.GET(    "/helth", handler.Helth())
    router.GET(    "/v1",    handler.Hello())
    return router
}

/* //<-- router.go ends here */
