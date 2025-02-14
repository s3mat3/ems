/**
 * @file login.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package port

import (
	"ems/ss/entity/model"
	// "ems/ss/usecase/port"
)

type LoginInput interface {
    Authenticate(req LoginRequest) (model.Permission, error)
    ProvideToken(perm model.Permission, ssid string) (string, error)
}

type LoginOutput interface {
    Response()
}
/* //<-- login.go ends here */
