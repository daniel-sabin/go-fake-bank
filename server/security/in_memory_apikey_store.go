package security

import (
	"fmt"
)

type InMemoryApiKeyStore struct {
	store []string
}

func NewInMemoryApiKeyStore() *InMemoryApiKeyStore {
	return &InMemoryApiKeyStore{make([]string, 0)}
}

func (i *InMemoryApiKeyStore) save(key string) {
	fmt.Printf("trying to save key %v\r\n", key)
	i.store = append(i.store, key)
}

func (i *InMemoryApiKeyStore) exists(key string) bool {
	for _, k := range i.store {
		if k == key {
			return true
		}
	}
	return false
}
