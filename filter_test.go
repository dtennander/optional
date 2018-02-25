package optional

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestOptionalImpl_Filter(t *testing.T) {
	// Given
	a := getOptional(1)
	// When
	b := a.Filter(func(a mockObject) bool {
		return a.value == 1
	})
	// Then
	assert.True(t, b.IsPresent())
}

func TestOptionalImpl_Filter2(t *testing.T) {
	// Given
	a := getOptional(1)
	// When
	b := a.Filter(func(a mockObject) bool {
		return a.value > 1
	})
	// Then
	assert.False(t, b.IsPresent())
}

func TestOptionalImpl_Filter3(t *testing.T) {
	// Given
	optA := getOptional(1)
	evilPredicate := func(a *mockObject) bool { return true }
	// When
	evilCall := func() {
		optA.Filter(evilPredicate)
	}
	// Then
	assert.PanicsWithValue(t, "Predicate's argument is not compatible with this optional.", evilCall)
}

func TestOptionalImpl_Filter4(t *testing.T) {
	// Given
	a := getOptional(1)
	evilPredicate := func(a mockObject) mockObject { return a }
	// When
	evilCall := func() {
		a.Filter(evilPredicate) // Not a function
	}
	// Then
	assert.PanicsWithValue(t, "Argument must be a predicate function.", evilCall)
}

func TestFilterOnEmpty(t *testing.T) {
	// Given
	a := Empty()
	// When
	b := a.Filter(func(a mockObject) bool {
		return true
	})
	// Then
	assert.False(t, b.IsPresent())
}
