/**
 * @file login_dto.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package port


type LoginRequest struct {
    Info string `json:"info"`
    IV string `json:"iv"`
    Ssid string `json:"id"`
}


/* //<-- login_dto.go ends here */
