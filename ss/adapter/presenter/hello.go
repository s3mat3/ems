/**
 * @file hello.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package presenter

import (
    "ems/ss/usecase/port"
)

// type Pubkey struct {
//     CommonMessage
//     Pub string `json:"pub"`
// }

// type Ssid struct {
//     CommonMessage
//     Ssid string `json:"ssid"`
// }

type Hello struct {}

func NewHello() port.HelloOut {
    return &Hello{}
}

func (s *Hello) ResponsePubkey(key string) port.Pubkey {
    o := port.Pubkey{
        Pub: key,
    }
    if key == "" {
        o.CommonMessage = port.InternalError()
    } else {
        o.CommonMessage = port.Ok()
    }
    return o
}

func (s *Hello) ResponseSsid(id string) port.Ssid {
    o := port.Ssid{
        Ssid: id,
    }
    if id == "" {
        o.CommonMessage = port.InternalError()
    } else {
        o.CommonMessage = port.Ok()
    }
    return o
}




/* //<-- hello.go ends here */
