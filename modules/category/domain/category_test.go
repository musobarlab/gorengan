package domain

import "testing"

func TestCategory(t *testing.T) {
	t.Run("should return success test validate category", func(t *testing.T) {
		category := &Category{
			ID:   "1",
			Name: "Smart Phone",
		}

		err := category.Validate()

		if err != nil {
			t.Error("validate category should return nil of error")
		}
	})

	t.Run("should return success test validate category when field name empty", func(t *testing.T) {
		category := &Category{
			ID:   "1",
			Name: "",
		}

		err := category.Validate()

		if err == nil {
			t.Error("validate category should return error")
		}
	})
}
