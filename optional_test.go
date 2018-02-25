package optional

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockObject struct {
	value int32
}

func TestOptionalImpl_IsPresent(t *testing.T) {
	// Given
	opt := Of(1)
	// When
	result := opt.IsPresent()
	// Then
	assert.True(t, result)
}

func TestOptionalImpl_Get(t *testing.T) {
	// Given
	a := Of(1)
	// When
	value := a.Get()
	// Then
	assert.Equal(t, 1, value)
}

func TestOptionalImpl_Get2(t *testing.T) {
	// Given
	a := Empty()
	// When
	evilCode := func() {
		a.Get()
	}
	// Then
	assert.PanicsWithValue(t, "Accessing empty Optional.", evilCode)
}

func TestOf(t *testing.T) {
	// Given
	a := mockObject{value: 1}
	// When
	result := Of(a)
	// Then
	assert.Implements(t, (*Optional)(nil), result)
	assert.Equal(t, a, result.Get())
}

func TestOfNil(t *testing.T) {
	// Given
	var a *mockObject = nil
	// When
	evilCode := func() {
		Of(a)
	}
	// Then
	assert.PanicsWithValue(t, "nonNilValue can not be nil!", evilCode)

}

func TestOfPossibleNil(t *testing.T) {
	// Given
	var a *mockObject = nil
	// When
	optional := OfPossibleNil(a)
	// Then
	assert.False(t, optional.IsPresent())
}

func TestOfPossibleNil2(t *testing.T) {
	// Given
	a := mockObject{1}
	// When
	optional := OfPossibleNil(a)
	// Then
	assert.True(t, optional.IsPresent())
	assert.Equal(t, a, optional.Get())
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
