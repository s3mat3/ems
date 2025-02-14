/**
 * @file login.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package presenter

type LoginResponse struct {
    Message string `json:"message"`
    Token string `json:"token"`
}

func BuildLoginResponse(m string, t string) *LoginResponse {
    return &LoginResponse {
        Message: m,
        Token: t,
    }
}
/* //<-- login.go ends here */
