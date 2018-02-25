package optional

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptional_IfPresent(t *testing.T) {
	// Given
	a := Empty()
	var i int = 0
	consumer := func(value int) { i += 1 }
	// When
	a.IfPresent(consumer)
	// Then
	assert.Equal(t, 0, i, "Consumer should not get called!")
}

func TestOptional_IfPresent2(t *testing.T) {
	// Given
	a := Of(1)
	var i int = 0
	consumer := func(value int) { i += 1 }
	// When
	a.IfPresent(consumer)
	// Then
	assert.Equal(t, 1, i, "consumer did not get called!")
}

func TestOptional_IfPresent3(t *testing.T) {
	// Given
	optA := Of(1)
	evilMap := func(a string) {}
	// When
	evilCall := func() {
		optA.IfPresent(evilMap)
	}
	// Then
	assert.PanicsWithValue(t, "Consumer function's argument is not compatible with this optional.", evilCall)
}

func TestOptional_IfPresent4(t *testing.T) {
	// Given
	a := Of(1)
	// When
	evilCall := func() {
		a.IfPresent(1) // Not a function
	}
	// Then
	assert.PanicsWithValue(t, "Argument must be a consumer.", evilCall)
}

func TestOptional_IfPresent5(t *testing.T) {
	// Given
	a := Of(1)
	// When
	evilCall := func() {
		a.IfPresent(func() int { return 1 }) // Not a consumer
	}
	// Then
	assert.PanicsWithValue(t, "Argument must be a consumer.", evilCall)
}
