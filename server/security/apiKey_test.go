package security_test

import (
	"engineecore/demobank-server/server/security"
	"fmt"
	"testing"
)

func TestChekApiKey(t *testing.T) {

	t.Run("Api key", func(t *testing.T) {
		t.Parallel()

		type TestCase struct {
			key     string
			allowed bool
		}

		// Given
		testCases := []TestCase{
			{"1233456", true},
			{"1233457", true},
			{"1233455", false},
		}

		im := security.NewInMemoryApiKeyStore()
		store := security.ApiKeyStoreFactory(im)
		isKeyAllowed := security.IsKeyAllowedFactory(im)

		for _, testCase := range testCases {
			// When
			if testCase.allowed {
				store(testCase.key)
			}
			allowed, err := isKeyAllowed(testCase.key)

			fmt.Print(err)

			// Then
			if allowed != testCase.allowed {
				t.Errorf("error, apiKey %v check to %v must be %v", testCase.key, allowed, testCase.allowed)
			}
		}
	})

}
