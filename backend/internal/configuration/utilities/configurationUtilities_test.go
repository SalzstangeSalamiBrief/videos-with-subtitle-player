package utilities

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetEnvironmentStringInput(t *testing.T) {
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

func FuzzGetEnvironmentIntInput(f *testing.F) {
	for _, seeds := range []int64{1, 2, 283, -12} {
		f.Add(seeds)
	}

	f.Fuzz(func(t *testing.T, seed int64) {
		t.Setenv("envVariable", fmt.Sprintf("%d", seed))
		result, err := GetEnvironmentInt("envVariable", false, nil)
		if err != nil {
			t.Errorf("Expected no error but received '%s'", err.Error())
		}

		if result != seed {
			t.Errorf("Expected '%d' but received '%d'", seed, result)
		}
	})
}

func FuzzGetEnvironmentIntDefaultValue(f *testing.F) {
	for _, seeds := range []int64{1, 2, 283, -12} {
		f.Add(seeds)
	}

	f.Fuzz(func(t *testing.T, seed int64) {
		result, err := GetEnvironmentInt("envVariable", false, &seed)
		if err != nil {
			t.Errorf("Expected no error but received '%s'", err.Error())
		}

		if result != seed {
			t.Errorf("Expected '%d' but received '%d'", seed, result)
		}
	})
}

func FuzzGetEnvironmentIntEmptyButRequiredThrowsError(f *testing.F) {
	for _, seeds := range []int64{1, 2, 283, -12} {
		f.Add(seeds)
	}

	f.Fuzz(func(t *testing.T, seed int64) {
		result, err := GetEnvironmentInt("envVariable", true, &seed)
		if err == nil {
			t.Errorf("Expected no error but received '%s'", err.Error())
		}

		if result != 0 {
			t.Errorf("Expected '%d' but received '%d'", seed, result)
		}
	})
}

// TODO TEST REQUIRED
// TODO TEST INT AND BOOL
