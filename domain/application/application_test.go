package application_test

import (
	"engineecore/demobank-server/domain/application"
	"testing"
)

type DumbStore struct {
}

func (i *DumbStore) Save(app *application.Application) {
}

func (i *DumbStore) ReadApplication(name string) *application.Application {
	return nil
}

func TestApplication(t *testing.T) {

	t.Run("Register a new application", func(t *testing.T) {
		// Given
		testCases := [...]string{
			"homerApplication",
			"simpsonApplication",
		}

		store := application.ApplicationStoreFactory(new(DumbStore))

		for _, name := range testCases {
			// When
			app := store(name)

			// Then
			if app.ClientId.String() == "" {
				t.Error("client id not set")
			}

			if app.ClientSecret.String() == "" {
				t.Error("client id not set")
			}
		}

		t.Run("Read application", func(t *testing.T) {
			// Given
			read := application.ApplicationReadStoreFactory(new(DumbStore))

			// When - Then
			if read("fake") != nil {
				t.Error("invalid read application")
			}

		})

	})

}
