package utilities

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetEnvironmentStringSuccessWithoutDefault(t *testing.T) {
	seeds := []string{"test", "    test", "test     ", "12 3445"}

	for _, seed := range seeds {
		title := fmt.Sprintf("Seed '%s' successfully without default value", seed)

		t.Run(title, func(t *testing.T) {
			t.Setenv("envVariable", seed)

			result, err := GetEnvironmentString("envVariable", false, nil)
			if err != nil {
				t.Errorf("Expected no error but received '%s'", err.Error())
			}

			expectedValue := strings.TrimSpace(seed)
			if result != expectedValue {
				t.Errorf("Expected '%s' but received '%s'", expectedValue, result)
			}
		})
	}
}

func FuzzGetEnvironmentStringDefaultValue(f *testing.F) {
	for _, seeds := range []string{"test", "    test", "test     ", "12 3445"} {
		f.Add(seeds)
	}

	f.Fuzz(func(t *testing.T, seed string) {
		result, err := GetEnvironmentString("envVariable", false, &seed)
		if err != nil {
			t.Errorf("Expected no error but received '%s'", err.Error())
		}

		expectedValue := strings.TrimSpace(seed)
		if result != expectedValue {
			t.Errorf("Expected '%s' but received '%s'", expectedValue, result)
		}
	})
}

// TODO TEST REQUIRED
// TODO TEST INT AND BOOL
