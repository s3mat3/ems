/**
 * @file helth.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package controller

import (
    "net/http"

    "ems/ss/usecase/port"
)

type Helth struct {}

func NewHelth() *Helth {
    return &Helth{}
}

func (s *Helth) Check(c Context) {
    c.JSON(http.StatusOK, port.CreateMessage("I'm fine!"))
}
/* //<-- helth.go ends here */
