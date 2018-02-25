package optional

import "reflect"

// The Optional type represents a possible value that might be there.
// It offers a API that allows for manipulating the possible value returning defaults.
type Optional interface {
	IsPresent() bool
	Get() interface{}
	Map(mapFunc interface{}) Optional
	Filter(predicate interface{}) Optional
	OrElse(defaultValue interface{}) interface{}
}

type optionalImpl struct {
	isPresent bool
	value     interface{}
}

// Returns the value contained in the Optional.
// Will panic if called on an empty Optional.
func (o *optionalImpl) Get() interface{} {
	if o.isPresent {
		return o.value
	} else {
		panic("Accessing empty Optional.")
	}
}

// Returns the value if a value is present else it returns the defaultValue.
func (o *optionalImpl) OrElse(defaultValue interface{}) interface{} {
	if o.isPresent {
		return o.value
	} else {
		return defaultValue
	}
}

// Returns true if the Optional contains a value.
func (o *optionalImpl) IsPresent() bool {
	return o.isPresent
}

// Of constructs an Optional from a value separate from nil.
func Of(nonNilValue interface{}) Optional {
	assertNotNil(nonNilValue)
	return &optionalImpl{isPresent: true, value: nonNilValue}
}

func assertNotNil(nonNilValue interface{}) {
	if isNil(nonNilValue) {
		panic("nonNilValue can not be nil!")
	}
}

func isNil(possibleNil interface{}) bool {
	value := reflect.ValueOf(possibleNil)
	return value.Kind() == reflect.Ptr && value.IsNil()
}

// OfPossibleNil construct an Optional from a given value.
// If the value is nil a empty Optional is returned.
func OfPossibleNil(value interface{}) Optional {
	if isNil(value) {
		return Empty()
	} else {
		return Of(value)
	}
}

// Empty constructs an empty optional.
func Empty() Optional {
	return &optionalImpl{isPresent: false, value: nil}
}
