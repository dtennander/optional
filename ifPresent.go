package optional

import "reflect"

func (o *optionalImpl) IfPresent(consumer interface{}) {
	function := reflect.ValueOf(consumer)
	if !isConsumer(function.Type()) {
		panic("Argument must be a consumer.")
	}
	if o.isPresent {
		consumeValue(o, function)
	}
}

func isConsumer(consumer reflect.Type) bool {
	return isFunction(consumer) && consumer.NumIn() == 1 && consumer.NumOut() == 0
}

func consumeValue(o *optionalImpl, function reflect.Value) {
	value := reflect.ValueOf(o.value)
	if !value.Type().AssignableTo(function.Type().In(0)) {
		panic("Consumer function's argument is not compatible with this optional.")
	}
	function.Call([]reflect.Value{value})
}
