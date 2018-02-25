package optional

import f "github.com/DiTo04/optional/functions"

type optionalImpl struct {
	isPresent bool
	value     interface{}
}

func (o *optionalImpl) Get() interface{} {
	if !o.isPresent {
		panic("Accessing empty Optional.")
	}
	return o.value
}

func (o *optionalImpl) OrElse(defaultValue interface{}) interface{} {
	if o.isPresent {
		return o.value
	}
	return defaultValue
}

func (o *optionalImpl) IsPresent() bool {
	return o.isPresent
}

func (o *optionalImpl) Filter(predicateFunc interface{}) Optional {
	if !f.IsValid(f.Predicate, predicateFunc) {
		panic("Argument must be a predicate function.")
	} else if o.isPresent {
		return applyPredicate(o, predicateFunc)
	}
	return Empty()
}

func applyPredicate(o *optionalImpl, predicate interface{}) Optional {
	if !f.TakesArgument(predicate, o.value) {
		panic("Predicate's argument is not compatible with this optional.")
	}
	if f.CallFunction(predicate, o.value).Bool() {
		return o
	}
	return Empty()
}

func (o *optionalImpl) IfPresent(consumer interface{}) {
	if !f.IsValid(f.Consumer, consumer) {
		panic("Argument must be a consumer.")
	}
	if o.isPresent {
		consumeValue(consumer, o.value)
	}
}

func consumeValue(function interface{}, argument interface{}) {
	if !f.TakesArgument(function, argument) {
		panic("Consumer function's argument is not compatible with this optional.")
	}
	f.Consume(function, argument)
}

func (o *optionalImpl) Map(mapFunc interface{}) Optional {
	if !f.IsValid(f.Map, mapFunc) {
		panic("Argument must be a mapping function.")
	} else if o.isPresent {
		return mapValue(mapFunc, o.value)
	}
	return Empty()
}

func mapValue(mapFunc interface{}, argument interface{}) Optional {
	// Checking whether element type is convertible to function's first argument's type.
	if !f.TakesArgument(mapFunc, argument) {
		panic("Map function's argument is not compatible with this optional.")
	}
	result := f.CallFunction(mapFunc, argument)
	return OfPossibleNil(result.Interface())
}
