package typeparameters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
}

func TestSuma(t *testing.T) {
	assert.Equalf(t, MapStrings([]string{"1"}, func(s string) interface{} {}), Map(1, 2), "Add(1, 2) should be 3")
}
