/**
 * @file functional.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package persistent

import(
    "todo/entity/model"
)

type Functional interface {
    Search(...string) ([]model.Todo, error)
    Delete(string) (int, error)
    // Insert(string) (int, error)
    // Update(string) (int, error)
}

/* //<-- functional.go ends here */
