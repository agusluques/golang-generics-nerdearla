package constraintspackage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert := assert.New(t)

	assert.Equalf(2, Min[int](3, 2), "Min(3, 2) should be 2")
	// assert.Equalf(2, Min[int32](int32(3), int32(2)), "Min(3, 2) should be 2")
	// assert.Equalf(2, Min[int64](int64(3), int64(2)), "Min(3, 2) should be 2")

	assert.Equalf(1.1, Min[float64](1.1, 1.2), "Min(1.1, 1.2) should be 1.1")
	// assert.Equalf(1.1, Min[float32](float32(1.1), float32(1.2)), "Min(1.1, 1.2) should be 1.1")

	assert.Equalf("hello", Min[string]("hello", "nerdearla"), "Min(\"hello\", \"nerdearla\") should be \"hello\"")
}
