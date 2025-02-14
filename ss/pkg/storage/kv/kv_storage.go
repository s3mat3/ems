/**
 * @file kv_storge.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package kv

type KVStorage[T any] interface {
    Load(string) (T, bool)
    Store(string, T) bool
    Remove(string) bool
}

/* //<-- kv_storge.go ends here */
