/**
 * @file main.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package main

import (
    "todo/driver/web"
)

func main() {
    router := web.SetupRoute()
    if err := router.Run(":20080"); err != nil {
        panic(err)
    }

}
