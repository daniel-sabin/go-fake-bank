package security

import "errors"

type ApiKeyStore interface {
	Save(string)
	Exists(string) bool
}

func ApiKeyStoreFactory(s ApiKeyStore) func(key string) {
	return func(key string) {
		s.Save(key)
	}
}

func IsKeyAllowedFactory(s ApiKeyStore) func(k string) (bool, error) {
	return func(k string) (bool, error) {
		if !s.Exists(k) {
			return false, errors.New("api ley is not allowed")
		}
		return true, nil
	}
}
