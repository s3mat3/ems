/**
 * @file handler.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package registry

import (
    "ems/ss/driver/web/handler"
    "ems/ss/adapter/controller"
)

/// create HelthHandler
func CreateHelthHandler() *handler.Helth {
    return handler.NewHelth(controller.NewHelth())
}

func CreateHelloHandler() *handler.Hello {
    return handler.NewHello(CreateHelloController())
}

func CreateLoginHandler() *handler.Login {
    return handler.NewLogin()
}
/* //<-- handler.go ends here */
