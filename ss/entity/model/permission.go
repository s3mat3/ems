/**
 * @file permission.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package model

type Permission struct {
    id int
}

func (s *Permission) SetID(id int) {
    s.id = id
}

func (s *Permission) GetID() int {
    return s.id
}

/* //<-- permission.go ends here */
