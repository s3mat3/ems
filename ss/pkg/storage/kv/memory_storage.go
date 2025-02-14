/**
 * @file memory_storage.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package kv

import "sync"

type MemoryStorage[T any] struct {
    rooms map[string]T
    locker sync.RWMutex
}

func NewMemoryStorage[T any] () *MemoryStorage[T] {
    return &MemoryStorage[T] {
        rooms: make(map[string]T) ,
    }
}

func (self *MemoryStorage[T]) Load(key string) (T, bool) {
    self.locker.RLock()
    defer self.locker.RUnlock()
    value, exists := self.rooms[key]
    return value, exists
}

func (self *MemoryStorage[T]) Store(key string, value T) bool {
    self.locker.Lock()
    defer self.locker.Unlock()
    self.rooms[key] = value
    return true
}

func (self *MemoryStorage[T]) Remove(key string) bool {
    self.locker.Lock()
    defer self.locker.Unlock()
    _, exists := self.rooms[key]
    if exists != true {return false}
    delete(self.rooms, key)
    return true
}

func (self *MemoryStorage[T]) Rooms() map[string]T {
    return self.rooms
}

/* //<-- memory_storage.go ends here */
