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
    "log"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
    "ems/ss/driver/web"
)

//////////////////////////////////////////////////////////////////////////////
/// create server certificate and key
/// **CAUTION** this key and certificate only use in local test.
// cmd-line$ openssl genrsa -out server.key 4096
// cmd-line$ openssl req -new -x509 -days 365 -key server.key -out server.pem
// You are about to be asked to enter information that will be incorporated
// into your certificate request.
// What you are about to enter is what is called a Distinguished Name or a DN.
// There are quite a few fields but you can leave some blank
// For some fields there will be a default value,
// If you enter '.', the field will be left blank.
// -----
// Country Name (2 letter code) [AU]: <-- suitably
// State or Province Name (full name) [Some-State]: <-- suitably
// Locality Name (eg, city) []: <-- suitably
// Organization Name (eg, company) [Internet Widgits Pty Ltd]: <-- suitably
// Organizational Unit Name (eg, section) []: <-- suitably
// Common Name (e.g. server FQDN or YOUR name) []:localhost
// Email Address []: <-- suitably
//////////////////////////////////////////////////////////////////////////////
func main() {
    _, err := sqlx.Open("postgres", "postgres@/test_db")
    if err != nil {
        log.Fatal(err)
    }

    router := web.SetupRoute()
    if err := router.RunTLS(":10443", "server.pem", "server.key"); err != nil {
        panic(err)
    }

}

