/**
 * @file login.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package interactor

import (
    "ems/ss/usecase/port"
    "ems/ss/entity/model"
)

type Login struct {}

func NewLogin() port.LoginInput {
    return &Login{}
}

func (s *Login) Authenticate(req port.LoginRequest) (model.Permission, error) {
    // extract json req.info
    // json info to Authentication struct
    // invoke repository and get permission when match user
    return model.Permission{}, nil
}

func (s *Login) ProvideToken(perm model.Permission, ssid string) (string, error) {
    return "", nil
}
/* //<-- login.go ends here */
