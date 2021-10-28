package walkmap

import (
	"reflect"
)

// Visitor is a function that will be called on each non-map node
// keyPath is a slice containing consecutive keys used to arrive at the given value
// value is a non-map value, corresponding to the above path
// kind is `reflect.TypeOf(value).Kind()`
type Visitor func(keyPath []interface{}, value interface{}, kind reflect.Kind)

// Walk walks the given map `m` and calls the visitor at every non-map node
func Walk(m interface{}, visitor Visitor) {
	if reflect.TypeOf(m).Kind() != reflect.Map {
		panic("m must be a map")
	}

	walk([]interface{}{}, m, visitor)
}

func walk(keyPath []interface{}, m interface{}, visitor Visitor) {
	mapVal := reflect.ValueOf(m)
	for _, refKeyValue := range mapVal.MapKeys() {
		val := mapVal.MapIndex(refKeyValue).Interface()
		key := refKeyValue.Interface()

		kind := reflect.TypeOf(val).Kind()
		path := make([]interface{}, len(keyPath))
		copy(path, keyPath)
		path = append(path, key)

		if kind == reflect.Map {
			walk(path, val, visitor)
		} else {
			visitor(path, val, kind)
		}
	}
}
