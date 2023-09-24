package constraints

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	assert := assert.New(t)

	assert.Equalf(2, Min[int](3, 2), "Min(3, 2) should be 2")
	assert.Equalf(int32(2), Min[int32](int32(3), int32(2)), "Min(3, 2) should be 2")
	assert.Equalf(int64(2), Min[int64](int64(3), int64(2)), "Min(3, 2) should be 2")

	assert.Equalf(1.1, Min[float64](1.1, 1.2), "Min(1.1, 1.2) should be 1.1")
	assert.Equalf(float32(1.1), Min[float32](float32(1.1), float32(1.2)), "Min(1.1, 1.2) should be 1.1")

	assert.Equalf("hello", Min[string]("hello", "nerdearla"), "Min(\"hello\", \"nerdearla\") should be \"hello\"")
}

func TestGenMin(t *testing.T) {
	assert := assert.New(t)

	assert.Equalf(2, GenMin[int](3, 2), "GenMin(3, 2) should be 2")
	assert.Equalf(int32(2), GenMin[int32](int32(3), int32(2)), "GenMin(3, 2) should be 2")
	assert.Equalf(int64(2), GenMin[int64](int64(3), int64(2)), "GenMin(3, 2) should be 2")

	assert.Equalf(1.1, GenMin[float64](1.1, 1.2), "GenMin(1.1, 1.2) should be 1.1")
	assert.Equalf(float32(1.1), GenMin[float32](float32(1.1), float32(1.2)), "GenMin(1.1, 1.2) should be 1.1")

	assert.Equalf("hello", GenMin[string]("hello", "nerdearla"), "GenMin(\"hello\", \"nerdearla\") should be \"hello\"")
}

func TestSuma(t *testing.T) {
	assert.Equal(t, 5, Suma(2, 3), "Suma(2, 3) should equal 5")

	assert.Equal(t, int32(5), Suma(int32(2), int32(3)), "Suma(int32(2), int32(3)) should equal 5")

	assert.Equal(t, int64(5), Suma(int64(2), int64(3)), "Suma(int64(2), int64(3)) should equal 5")

	assert.Equal(t, float32(6.0), Suma(float32(2.5), float32(3.5)), "Suma(float32(2.5), float32(3.5)) should equal 6.0")

	assert.Equal(t, float64(6.0), Suma(float64(2.5), float64(3.5)), "Suma(float64(2.5), float64(3.5)) should equal 6.0")
}
