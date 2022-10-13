package application_test

import (
	"engineecore/demobank-server/domain/application"
	"testing"
)

func TestApplication(t *testing.T) {

	t.Run("Register a new application", func(t *testing.T) {
		// Given
		testCases := [...]string{
			"homerApplication",
			"simpsonApplication",
		}
		im := application.NewInMemoryApplicationStore()
		store := application.ApplicationStoreFactory(im)
		read := application.ApplicationReadStoreFactory(im)

		for _, name := range testCases {
			// When
			app := store(name)
			expected := read(name)

			// Then
			if app.ClientId != expected.ClientId {
				t.Errorf("Invalid result clientId %v, expected %v", app.ClientId, expected.ClientId)
			}

			if app.ClientSecret != expected.ClientSecret {
				t.Errorf("Invalid result clientId %v, expected %v", app.ClientSecret, expected.ClientSecret)
			}
		}

	})

}
