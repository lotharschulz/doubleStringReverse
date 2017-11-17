package myStringUtil

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDoubleStringReverseToReturnDoubleReversedInputString(t *testing.T) {
	actualResult := DoubleStringReverse("abcde")
	var expectedResult = "EDCBA"
	assert.Equal(t, expectedResult, actualResult)
	require.Equal(t, expectedResult, actualResult)
}
