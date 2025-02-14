/**
 * @file session_data.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package repository

import(
    "ems/ss/entity/model"
)

type SessionData interface {
    GetAll() (map[string]model.SessionData, error)
    FindByID(id string) (model.SessionData, error)
    Delete(id string) error
    Update(id string, sd model.SessionData) error
}
/* //<-- session_data.go ends here */
