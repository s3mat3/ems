/**
 * @file status.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package model

type Status struct {
    code   int
    name string
}

var statusList = []Status {
    {code: -10, name: "reject"},
    {code:  -1, name: "close"},
    {code:   0, name: "new"},
    {code:   1, name: "open"},
    {code:   2, name: "in progress"},
}
//
// dummy
//
func GetStatusCode(name string) int {
    for _, v := range statusList {
        if v.name == name {
            return v.code
        }
    }
    return 0 // mean new
}

func GetStatusName(code int) string {
    for _, v := range statusList {
        if v.code == code {
            return v.name
        }
    }
    return "new"
}
/* //<-- status.go ends here */
