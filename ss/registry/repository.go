/**
 * @file repository.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package registry

import (
    "ems/ss/driver/persistent"
    "ems/ss/entity/repository"
)

func CreateSessionDataRepository() repository.SessionData {
    return persistent.NewSessionData()
}


/* //<-- repository.go ends here */
