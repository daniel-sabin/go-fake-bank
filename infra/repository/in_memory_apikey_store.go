package repository

import (
	"engineecore/demobank-server/domain/security"
	"fmt"
)

type inMemoryApiKeyStore struct {
	store []string
}

func (i *inMemoryApiKeyStore) Save(key string) {
	fmt.Printf("trying to save key %v\r\n", key)
	i.store = append(i.store, key)
}

func (i *inMemoryApiKeyStore) Exists(key string) bool {
	for _, k := range i.store {
		if k == key {
			return true
		}
	}
	return false
}

func NewInMemoryApiKeyStore() security.ApiKeyStore {
	return &inMemoryApiKeyStore{make([]string, 0)}
}
