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

func TestOfNil(t *testing.T) {
	// Given
	var a *mockObject = nil
	// When
	evilCode := func() {
		Of(a)
	}
	// Then
	assert.PanicsWithValue(t,"nonNilValue can not be nil!", evilCode)

}

func TestOptionalImpl_OrElse(t *testing.T) {
	// Given
	a := Of(1)
	// When
	result := a.OrElse(2)

	// Then
	assert.Equal(t, 1, result)
}

func TestOptionalImpl_OrElse2(t *testing.T) {
	// Given
	a := Empty()
	// When
	result := a.OrElse("Default")

	// Then
	assert.Equal(t, "Default", result)
}

func TestEmpty(t *testing.T) {
	// Given
	a := Empty()
	// When
	isPresent := a.IsPresent()
	// Then
	assert.False(t, isPresent)
}