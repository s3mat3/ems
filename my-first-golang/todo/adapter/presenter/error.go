/**
 * @file error.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package presenter

type ErrorResponse struct {
    Message string `json:"message"`
}

func NewErrorResponse(msg string) *ErrorResponse {
    return &ErrorResponse{
        Message: msg,
    }
}
/* //<-- error.go ends here */
