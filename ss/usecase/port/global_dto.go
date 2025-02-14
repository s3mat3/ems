/**
 * @file global.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package port

// type Msg struct {
//     Message string `json:"message"`
// }

// func Message(m string) Msg {
//     return Msg{Message: m}
// }

type CommonMessage struct {
    Message string `json:"message"`
}

func CreateMessage(m string) CommonMessage {
    return CommonMessage{Message: m}
}

func Ok() CommonMessage {
    return CommonMessage{Message: "OK"}
}

func InternalError() CommonMessage {
    return CommonMessage{Message: "Internal server error"}
}

func NotfoundError() CommonMessage {
    return CommonMessage{Message: "Not found"}
}

/* //<-- global.go ends here */
