package reflection

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)
	length := 0

	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.Slice:
		getField = val.Index
		length = val.Len()
	case reflect.Struct:
		getField = val.Field
		length = val.NumField()
	case reflect.String:
		fn(val.String())
	case reflect.Array:
		getField = val.Index
		length = val.Len()
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	}

	for i := 0; i < length; i++ {
		walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
