/**
 * @file login.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package handler
import (
    "github.com/gin-gonic/gin"
)
type Login struct {}

func NewLogin() *Login {
    return &Login{}
}

func (s *Login) TryLogin() gin.HandlerFunc {
    return func(c *gin.Context) {
        
    }
}
/* //<-- login.go ends here */
