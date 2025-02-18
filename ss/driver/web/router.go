/**
 * @file router.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package web

import (
    "time"
    "os"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"

    "ems/ss/registry"
)

//
//
func SetupRoute() *gin.Engine {
    helth := registry.CreateHelthHandler()
    hello := registry.CreateHelloHandler()
    login := registry.CreateLoginHandler()

    router := gin.Default()
    SetupCors(router)
    SetupSecurityHeaders(router)
    // routing
    root := router.Group("/")
    {
        root.GET ( "helth", helth.Check()) // helth check
        root.GET ( "hello", hello.GET()) // given public key
        root.POST( "hello", hello.POST()) // take client secret key and serve ssid
        root.POST( "login", login.TryLogin()) // login request serve api token
        // v1 := router.Group("/v1")
        // {
        //     v1.POST()
        // }
    }
    return router
}

func SetupCors(router *gin.Engine) {
    origins := os.Getenv("PERMIT_URLS")
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
        // temp-id for session id
        // temp-range for access permission use in client
        ExposeHeaders: []string {
            "X-Temp-Id",
            "X-Temp-Range",
        },
        AllowCredentials: true,
        AllowOrigins: []string{origins},
        MaxAge: 24 * time.Hour,
    }))
}

func SetupSecurityHeaders(router *gin.Engine) {
    router.Use(func (c *gin.Context){
        c.Header("X-Frame-Options", "DENY")
        c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
        c.Header("X-XSS-Protection", "1; mode=block")
        c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
        c.Header("Referrer-Policy", "strict-origin")
        c.Header("X-Content-Type-Options", "nosniff")
        c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
    })
}
/* //<-- router.go ends here */
