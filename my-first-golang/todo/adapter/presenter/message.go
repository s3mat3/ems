/**
 * @file message.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package presenter

type MessageResponse struct {
    Message string `json:"message"`
}

func NewMessageResponse(msg string) * MessageResponse {
    return &MessageResponse{
        Message: msg,
    }
}
/* //<-- message.go ends here */
