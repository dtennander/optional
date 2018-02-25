package optional

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type mockObject struct {
	value int32
}

func getOptional(value int32) Optional {
	a := mockObject{value: value}
	optA := Of(a)
	return optA
}

func TestOptionalImpl_IsPresent(t *testing.T) {
	// Given
	opt := getOptional(1)
	// When
	result := opt.IsPresent()
	// Then
	assert.True(t, result)
}

func TestOf(t *testing.T) {
	// Given
	a := mockObject{value: 1}
	// When
	result := Of(a)
	// Then
	assert.Implements(t, (*Optional)(nil), result)
	assert.Equal(t,a, result.Get())
}

func TestEmpty(t *testing.T) {
	// Given
	a := Empty()
	// When
	isPresent := a.IsPresent()
	// Then
	assert.False(t, isPresent)
}

