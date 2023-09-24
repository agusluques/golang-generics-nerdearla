package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	assert := assert.New(t)

	assert.Equalf("1", GetStringFromInt(1), "GetStringFromInt(1) should be '1'")
	assert.Equalf("true", GetStringFromBoolean(true), "GetStringFromBoolean(true) should be 'true'")
	assert.Equalf("string", GetStringFromString("string"), `GetStringFromString("string") should be 'string'`)

	assert.Equalf("1", GetStringFrom(1), "GetStringFrom(1) should be '1'")
	assert.Equalf("true", GetStringFrom(true), "GetStringFrom(true) should be 'true'")
	assert.Equalf("string", GetStringFrom("string"), `GetStringFrom("string") should be 'string'`)
}
