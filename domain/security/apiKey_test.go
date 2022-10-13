package security_test

import (
	"engineecore/demobank-server/domain/security"
	"testing"
)

type DumbStore struct {
	Exist *bool
}

func (i *DumbStore) Save(key string) {
}

func (i *DumbStore) Exists(key string) bool {
	return *i.Exist
}

func TestChekApiKey(t *testing.T) {
	t.Run("Should allow api key", func(t *testing.T) {
		t.Parallel()

		type TestCase struct {
			allowed bool
		}

		// Given
		testCases := []TestCase{
			{true},
			{false},
		}

		for _, testCase := range testCases {
			isKeyAllowed := security.IsKeyAllowedFactory(&DumbStore{Exist: &testCase.allowed})

			// When
			allowed, _ := isKeyAllowed("fake")

			// Then
			if allowed != testCase.allowed {
				t.Errorf("error, apiKey allowed should be %v", testCase.allowed)
			}
		}
	})

}
