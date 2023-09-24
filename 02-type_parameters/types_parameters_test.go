package typeparameters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuma(t *testing.T) {
	assert.Equalf(t, 3, Suma(1, 2), "Add(1, 2) should be 3")
	assert.Equalf(t, 4, Suma(2, 2), "Add(2, 2) should be 4")
	assert.Equalf(t, 1, Suma(-2, 3), "Add(-2, 3) should be 1")
}

func TestContiene(t *testing.T) {
	t.Run("ElementPresent", func(t *testing.T) {
		lista := []int{1, 2, 3, 4, 5}
		element := 3
		assert.True(t, Contiene(lista, element))
	})

	t.Run("ElementNotPresent", func(t *testing.T) {
		lista := []int{1, 2, 3, 4, 5}
		element := 6
		assert.False(t, Contiene(lista, element))
	})

	t.Run("EmptyList", func(t *testing.T) {
		lista := []int{}
		element := 1
		assert.False(t, Contiene(lista, element))
	})

	t.Run("EmptyListElementNotPresent", func(t *testing.T) {
		lista := []int{}
		element := 0
		assert.False(t, Contiene(lista, element))
	})

	t.Run("DuplicateElements", func(t *testing.T) {
		lista := []int{1, 2, 2, 3, 4, 4, 5}
		element := 3
		assert.True(t, Contiene(lista, element))
	})
}
