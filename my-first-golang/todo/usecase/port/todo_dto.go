/**
 * @file todo_dto.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package port

type TodoResponse struct {
    Id      int    `json:"id"`
    Subject string `json:"subject"`
    Status  string `json:"status"`
}

type TodoMessageResponse struct {
    Message string `json:"message"`
}
/* //<-- todo_dto.go ends here */
