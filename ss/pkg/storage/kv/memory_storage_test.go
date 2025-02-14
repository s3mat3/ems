/**
 * @file memory_storage_test.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package kv_test

import (
	"exp/storage/kv"
    "reflect"
	"testing"
    "github.com/stretchr/testify/assert"
)

type keys struct {
    himitsu string
    naisho string
}

func TestSaveLoad(t *testing.T) {
    strStorage := kv.NewMemoryStorage[string]()
    strStorage.Store("Hello", "World")
    v, exists := strStorage.Load("Hello")
    assert.Equal(t, exists, true, "They shud be equal")
    assert.Equal(t, v, "World", "They shud be equal")
}

func updateNaisho(s *kv.MemoryStorage[keys], k string, v string) *kv.MemoryStorage[keys] {
    t, e := s.Load(k)
    if !e {
        t = keys{himitsu: "", naisho: ""}
    }
    t.naisho = v
    s.Store(k, t)
    return s
}

func TestTisStruct(t *testing.T) {
    key := keys{
        himitsu: "himitsu",
        naisho: "hoge",
    }
    storage := kv.NewMemoryStorage[keys]()
    uv, e := storage.Load("aaaa")
    assert.Equal(t, e, false, "They shud be equal")
    storage.Store("aaaa", key)
    value, exists := storage.Load("aaaa")
    assert.Equal(t, exists, true, "They shud be equal")
    tn := reflect.TypeOf(value)
    assert.Equal(t, exists, true, "They shud be equal")
    assert.Equal(t, tn.Name(), "keys", "They shud be equal")
    assert.Equal(t, value.himitsu, "himitsu", "They shud be equal")
    assert.Equal(t, value.naisho, "hoge", "They shud be equal")
    updateNaisho(storage, "aaaa", "hello")
    uv, e = storage.Load("aaaa")
    assert.Equal(t, e, true, "They shud be equal")
    assert.Equal(t, uv.naisho, "hello", "They shud be equal")
}

/* //<-- kv_memory_storage_test.go ends here */
