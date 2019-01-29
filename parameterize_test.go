package go_parameterize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParameterize(t *testing.T) {
	assert.Equal(t, Parameterize("ABC 123"), "abc-123")
	assert.Equal(t, Parameterize(" ABC 123 "), "abc-123")
	assert.Equal(t, Parameterize("    ABC      123     "), "abc-123")
	assert.Equal(t, Parameterize("--- ABC **** 123   "), "abc-123")
	assert.Equal(t, Parameterize("&$@#!$@#ABC@@@@123"), "abc-123")
	assert.Equal(t, Parameterize("__ABC 123"), "abc-123")
	assert.Equal(t, Parameterize("ABC _ 123"), "abc-123")
	assert.Equal(t, Parameterize("ABC^^^ 123"), "abc-123")
	assert.Equal(t, Parameterize("ABC:123"), "abc-123")
	assert.Equal(t, Parameterize("abc-123"), "abc-123")
	assert.Equal(t, Parameterize("$%&*"), "")
	assert.Equal(t, Parameterize("$#@1#$%"), "1")
}
