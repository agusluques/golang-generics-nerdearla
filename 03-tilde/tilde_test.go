package tilde

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintFloats(t *testing.T) {
	assert := assert.New(t)

	assert.Equalf("1.00", GetFloatsIn2Decimals(float64(1.0)), "PrintFloatsIn2Decimals(1) should be '1.00'")

	assert.Equalf("0.00", GetFloatsIn2Decimals(float32(0.0)), "PrintFloatsIn2Decimals(0) should be '0.00'")

	type MyFloat32 float32
	assert.Equalf("2.26", GetFloatsIn2Decimals(MyFloat32(2.260611)), "PrintFloatsIn2Decimals(0) should be '2.26'")

	type MyFloat64 float64
	assert.Equalf("1.00", GetFloatsIn2Decimals(MyFloat64(1.0)), "PrintFloatsIn2Decimals(1) should be '1.00'")
}
