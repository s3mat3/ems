/**
 * @file context.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package controller
// This idea is from https://github.com/arakawamoriyuki/go-clean-handson,
// so the gin framework will not leak than the infrastructure (driver directory).
//
type Context interface {
    Abort()
    BindJSON(obj interface{}) error
    ShouldBindJSON(obj interface{}) error
    JSON(code int, obj interface{})
    Param(key string) string
    Header(string, string)
}


/* //<-- context.go ends here */
