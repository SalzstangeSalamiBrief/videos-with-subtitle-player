package utilities

import (
	"fmt"
	"strings"
	"testing"

	"backend/pkg/models"
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

func TestGetEnvironmentStringUsesDefaultValue(t *testing.T) {
	seeds := []string{"test", "    test", "test     ", "12 3445"}

	for _, seed := range seeds {
		title := fmt.Sprintf("Seed '%s' uses default value when env var is not set", seed)

		t.Run(title, func(t *testing.T) {
			result, err := GetEnvironmentString("envVariableNotSet", false, &seed)
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

func TestGetEnvironmentStringRequiredThrowsError(t *testing.T) {
	t.Run("isRequired but no env var and no default value throws error", func(t *testing.T) {
		result, err := GetEnvironmentString("envVariableNotSet", true, nil)
		if err == nil {
			t.Errorf("Expected an error but received none")
		}

		if result != "" {
			t.Errorf("Expected empty string but received '%s'", result)
		}
	})
}

func TestGetEnvironmentIntThrowsParsingError(t *testing.T) {
	seeds := []models.TestData[string, bool]{
		{Title: "non-numeric string 'notAnInt' throws parsing error", Input: "notAnInt", Expected: true},
		{Title: "float string '12.34' throws parsing error", Input: "12.34", Expected: true},
		{Title: "alphanumeric string 'abc123' throws parsing error", Input: "abc123", Expected: true},
	}

	for _, seed := range seeds {
		t.Run(seed.Title, func(t *testing.T) {
			t.Setenv("envVariable", seed.Input)

			_, err := GetEnvironmentInt("envVariable", false, nil)
			if err == nil {
				t.Errorf("Expected an error but received none")
			}
		})
	}
}

func TestGetEnvironmentBooleanThrowsParsingError(t *testing.T) {
	seeds := []models.TestData[string, bool]{
		{Title: "arbitrary string 'notABool' throws parsing error", Input: "notABool", Expected: true},
		{Title: "out-of-range integer '12' throws parsing error", Input: "12", Expected: true},
		{Title: "word 'yes' throws parsing error", Input: "yes", Expected: true},
		{Title: "word 'no' throws parsing error", Input: "no", Expected: true},
	}

	for _, seed := range seeds {
		t.Run(seed.Title, func(t *testing.T) {
			t.Setenv("envVariable", seed.Input)

			_, err := GetEnvironmentBoolean("envVariable", false, nil)
			if err == nil {
				t.Errorf("Expected an error but received none")
			}
		})
	}
}

func TestGetEnvironmentBooleanInput(t *testing.T) {
	seeds := []models.TestData[string, bool]{
		{Title: "lowercase 'true' parses to true", Input: "true", Expected: true},
		{Title: "lowercase 'false' parses to false", Input: "false", Expected: false},
		{Title: "'1' parses to true", Input: "1", Expected: true},
		{Title: "'0' parses to false", Input: "0", Expected: false},
		{Title: "uppercase 'TRUE' parses to true", Input: "TRUE", Expected: true},
		{Title: "uppercase 'FALSE' parses to false", Input: "FALSE", Expected: false},
	}

	for _, seed := range seeds {
		title := fmt.Sprintf("Seed '%s' successfully parsed as '%v'", seed.Input, seed.Expected)

		t.Run(title, func(t *testing.T) {
			t.Setenv("envVariable", seed.Input)

			result, err := GetEnvironmentBoolean("envVariable", false, nil)
			if err != nil {
				t.Errorf("Expected no error but received '%s'", err.Error())
			}

			if result != seed.Expected {
				t.Errorf("Expected '%v' but received '%v'", seed.Expected, result)
			}
		})
	}
}

func FuzzGetEnvironmentBooleanDefaultValue(f *testing.F) {
	for _, seed := range []bool{true, false} {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, seed bool) {
		result, err := GetEnvironmentBoolean("envVariableNotSet", false, &seed)
		if err != nil {
			t.Errorf("Expected no error but received '%s'", err.Error())
		}

		if result != seed {
			t.Errorf("Expected '%v' but received '%v'", seed, result)
		}
	})
}
