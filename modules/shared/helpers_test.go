package shared

import (
	"testing"
)

func TestStringInSlice(t *testing.T) {
	list := []string{"Wury", "Alex", "Dony"}

	t.Run("should return true test string in slice", func(t *testing.T) {

		expected := true

		result := StringInSlice("Wury", list)

		if expected != result {
			t.Error("data not in list")
		}
	})

	t.Run("should return true test string in slice", func(t *testing.T) {

		expected := false

		result := StringInSlice("Ben", list)

		if expected != result {
			t.Error("data should not in list")
		}
	})
}
