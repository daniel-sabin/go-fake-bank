package security

import "errors"

type ApiKeyStore interface {
	save(key string)
	exists(key string) bool
}

func ApiKeyStoreFactory(s ApiKeyStore) func(key string) {
	return func(key string) {
		s.save(key)
	}
}

func IsKeyAllowedFactory(s ApiKeyStore) func(k string) (bool, error) {
	return func(k string) (bool, error) {
		if !s.exists(k) {
			return false, errors.New("api ley is not allowed")
		}
		return true, nil
	}
}
