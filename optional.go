package optional

import "reflect"

type Optional interface {
	IsPresent() bool
	Get() interface{}
	Map(mapFunc interface{}) Optional
	Filter(predicate interface{}) Optional
	OrElse(defaultValue interface{}) interface{}
}

type optionalImpl struct {
	isPresent bool
	value interface{}
}

func (o *optionalImpl) Get() interface{} {
	return o.value
}

func (o *optionalImpl) OrElse(defaultValue interface{}) interface{} {
	if o.isPresent {
		return o.value
	} else {
		return defaultValue
	}
}

func (o *optionalImpl) IsPresent() bool {
	return o.isPresent
}

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

func OfPossibleNil(value interface{}) Optional {
	if isNil(value) {
		return Empty()
	} else {
		return Of(value)
	}
}

func Empty() Optional {
	return &optionalImpl{isPresent:false, value:nil}
}