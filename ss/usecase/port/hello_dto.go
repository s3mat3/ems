/**
 * @file hello_dto.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package port
/////////////////////////
/// Response data struct
/////////////////////////

/// Pubkey response
type Pubkey struct {
    CommonMessage
    Pub string `json:"pub"`
}
/// Ssid response
type Ssid struct {
    CommonMessage
    Cause string `json:"cause"`
    Ssid string `json:"ssid"`
}

////////////////////////
/// Request data struct
////////////////////////

/// Himitsu(secret key) request
type Himitsu struct {
    Himitsu string `json:"himitsu"`
}
/* //<-- hello_dto.go ends here */
