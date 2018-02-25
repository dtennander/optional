package optional

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestOptionalImpl_Map(t *testing.T) {
	// Given
	optA := getOptional(1)
	// When
	optB := optA.Map(func(a mockObject) mockObject {
		a.value += 1
		return a
	})
	// Then
	assert.Equal(t, mockObject{value:2}, optB.Get())
}

func TestOptionalImpl_Map2(t *testing.T) {
	// Given
	optA := getOptional(1)
	evilMap := func(a *mockObject) *mockObject { return a }
	// When
	evilCall := func() {
		optA.Map(evilMap)
	}
	// Then
	assert.PanicsWithValue(t, "Map function's argument is not compatible with this optional.", evilCall)
}

func TestOptionalImpl_Map3(t *testing.T) {
	// Given
	a := getOptional(1)
	// When
	evilCall := func() {
		a.Map(1) // Not a function
	}
	// Then
	assert.PanicsWithValue(t, "Argument must be a mapping function.", evilCall)
}



func TestMapOnEmpty(t *testing.T) {
	// Given
	a := Empty()
	// When
	b := a.Map(func(a mockObject) mockObject{
		return a
	})
	// Then
	assert.False(t, b.IsPresent())
}