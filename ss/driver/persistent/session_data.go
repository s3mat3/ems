/**
 * @file session_data.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package persistent

import (
    "errors"
    "sync"
    "ems/ss/entity/model"
    "ems/ss/entity/repository"
    "ems/ss/pkg/storage/kv"
)

var storage *kv.MemoryStorage[model.SessionData]
var once sync.Once

func InitializeMemoryStorage() *kv.MemoryStorage[model.SessionData] {
    once.Do(func () {
        storage = kv.NewMemoryStorage[model.SessionData]()
    })
    return storage
}

type sessionData struct {}

func NewSessionData() repository.SessionData {
    InitializeMemoryStorage()
    return &sessionData{}
}

func (s *sessionData) GetAll() (map[string]model.SessionData, error) {
    return storage.Rooms(), nil
}

func (s *sessionData) FindByID(id string) (model.SessionData, error) {
    data, exists := storage.Load(id)
    if !exists {
        return data, errors.New("no entry")
    }
    return data, nil
}

func (s *sessionData) Delete(id string) error {
    if exists := storage.Remove(id); !exists {
        return errors.New("no entry")
    }
    return nil
}

func (s *sessionData) Update(id string, sd model.SessionData) error {
    storage.Store(id, sd)
    return nil
}

/* //<-- session_data.go ends here */
