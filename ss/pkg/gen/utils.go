/**
 * @file utils.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package gen

import (
    "crypto/rand"
    "time"
)

// Generate random letter byte stream length to size
func RandomLetters(size uint32) []byte {
    // 64 letters
    const base = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+/"
    const bs = len(base)
    // Random number
    rnum := make([]byte, size)
    // Random vector
    rv   := make([]byte, size)
    if _, err := rand.Read(rnum); err != nil {
        return nil
    }
    for k, v := range rnum {
        rv[k] = base[int(v) % bs]
    }
    return rv
}
// \param[in] from start in time.Time
// \param[in] after_sec in millisecond
// Genarate expiry time stamp (Unix time stamp come from epoch 1970-01-01T00:00:00Z)
func ExpiryTimeStamp(from time.Time, after_sec uint32) int64 {
    return from.Add(time.Duration(after_sec) * time.Second).Unix()
}



/* //<-- utils.go ends here */
