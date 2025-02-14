/**
 * @file interactor.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package registry

import (
    "ems/ss/usecase/interactor"
    "ems/ss/usecase/port"
)

func CreateHelloUsecase() port.HelloInput {
    return interactor.NewHelloUsecase(
        CreateHelloPresenter(),
        CreateSessionDataRepository())
}


/* //<-- interactor.go ends here */
