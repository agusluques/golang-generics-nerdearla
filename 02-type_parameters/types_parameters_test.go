package typeparameters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert := assert.New(t)

	assert.Equalf(3, Add(1, 2), "Add(1, 2) should be 3")

	assert.Equalf(4, Add(2, 2), "Add(2, 2) should be 4")

	assert.Equalf(1, Add(-2, 3), "Add(-2, 3) should be 1")

	// assert.Equalf(2.3, Add(1.1, 1.2), "Add(1.1, 1.2) should be 2.3")

	// assert.Equalf("ab", Add("a", "b"), "Add(\"a\", \"b\") should be \"ab\"")
}
