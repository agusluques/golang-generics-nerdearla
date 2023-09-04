package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	assert := assert.New(t)

	assert.Equalf("1", GetStringFromInt(1), "GetStringFromInt(1) should be '1'")

	assert.Equalf("0", GetStringFromInt(0), "GetStringFromInt(0) should be '0'")

	assert.Equalf("true", GetStringFromBoolean(true), "GetStringFromBoolean(true) should be 'true'")

	// create GetString method that receives any type and returns a string
}
