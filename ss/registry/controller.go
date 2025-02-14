/**
 * @file controller.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package registry

import (
    "ems/ss/adapter/controller"
)

func CreateHelloController() *controller.Hello {
    return controller.NewHello(CreateHelloUsecase())
}

/* //<-- controller.go ends here */
