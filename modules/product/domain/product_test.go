package domain

import (
	"testing"
)

func TestProduct(t *testing.T) {
	t.Run("should return success test validate product", func(t *testing.T) {
		product := &Product{
			ID:         "1",
			Name:       "Mac Book Pro 2018",
			CategoryID: "3",
			Quantity:   5,
		}

		err := product.Validate()

		if err != nil {
			t.Error("validate product should return nil of error")
		}
	})

	t.Run("should return error test validate product when field name empty", func(t *testing.T) {
		product := &Product{
			ID:         "1",
			Name:       "",
			CategoryID: "3",
			Quantity:   5,
		}

		err := product.Validate()

		if err == nil {
			t.Error("validate product should return error")
		}
	})
}
