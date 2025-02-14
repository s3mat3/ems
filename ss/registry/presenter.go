/**
 * @file presenter.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package registry

import(
    "ems/ss/adapter/presenter"
    "ems/ss/usecase/port"
)

func CreateHelloPresenter() port.HelloOut {
    return presenter.NewHello()
}
/* //<-- presenter.go ends here */
