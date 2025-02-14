/**
 * @file b64.go
 * @brife Wrapper for encoding/base64
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package b64

import (
    "bytes"
    "encoding/base64"
)

//
// Decode from base64 encoded string to raw byte in byte array []byte
//
func Decode(b64 string) ([]byte, error) {
    // convert string to []byte
    return DecodeB([]byte(b64))
}
/// Decode byte array to 
func DecodeB(b64 []byte) ([]byte, error) {
    // create destination
    dest := make([]byte, base64.StdEncoding.DecodedLen(len(b64)))
    // execute decode
    n, err := base64.StdEncoding.Decode(dest, b64)
    if err != nil {
        return nil, err
    }
    // slog.Info(fmt.Sprintln(n))
    return dest[:n], nil
}
/// Encode to base64 from byte array
func EncodeB(raw []byte) []byte {
    dest := make([]byte, base64.StdEncoding.EncodedLen(len(raw)))
    base64.StdEncoding.Encode(dest, raw)
    return dest
}
//
/// Decode from base64 URL encoded string to raw byte in byte array []byte
func DecodeURL(url64 string) ([]byte, error) {
    return DecodeURLB([]byte(url64))
}
///
func DecodeURLB(url64 []byte) ([]byte, error) {
    dest := make([]byte, base64.RawURLEncoding.DecodedLen(len(url64)))
    n, err := base64.RawURLEncoding.Decode(dest, url64); if err != nil {
        return nil, err
    }
    // For some reason,
    // the padding character "=" is converted to null, so cut it out.
    return bytes.Trim(dest[:n], "\x00"), nil
}
/// Encode to base64 URL safe 
func EncodeURLB(raw []byte) []byte {
    dest := make([]byte, base64.RawURLEncoding.EncodedLen(len(raw)))
    base64.RawURLEncoding.Encode(dest, raw)
    return dest
}

/* //<-- b64.go ends here */
