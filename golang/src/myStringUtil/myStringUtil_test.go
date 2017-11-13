package myStringUtil

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDoubleStringReverseToReturnDoubleReversedInputString(t *testing.T) {
	actualResult := Reverse("abc")
	var expectedResult = "CBA"
	t.Fatal()
	assert.Equal(t, expectedResult, actualResult)
	require.Equal(t, expectedResult, actualResult)
}
