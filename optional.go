package optional

import "reflect"

type Optional interface {
	IsPresent() bool
	Get() interface{}
	Map(mapFunc interface{}) Optional
	Filter(predicate interface{}) Optional
}

type optionalImpl struct {
	isPresent bool
	value interface{}
}

func (o *optionalImpl) Get() interface{} {
	return o.value
}

func (o *optionalImpl) IsPresent() bool {
	return o.isPresent
}

func Of(nonNilValue interface{}) Optional {
	assertNotNil(nonNilValue)
	return &optionalImpl{isPresent: true, value: nonNilValue}
}

func assertNotNil(nonNilValue interface{}) {
	value := reflect.ValueOf(nonNilValue)
	if value.Kind() == reflect.Ptr && value.IsNil() {
		panic("nonNilValue can not be nil!")
	}
}

func Empty() Optional {
	return &optionalImpl{isPresent:false, value:nil}
}