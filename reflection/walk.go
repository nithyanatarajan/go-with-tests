package reflection

import "reflect"

func Walk(x any, fn func(input string)) {
	value := getValue(x)
	switch value.Kind() {
	case reflect.String:
		fn(value.String())
	case reflect.Struct:
		walkStruct(value, fn)
	case reflect.Slice, reflect.Array:
		walkArray(value, fn)
	case reflect.Map:
		walkMap(value, fn)
	case reflect.Func:
		walkFunc(value, fn)
	}
}

func walkFunc(value reflect.Value, fn func(input string)) {
	{
		for _, val := range value.Call(nil) {
			Walk(val, fn)
		}
	}
}

func walkMap(value reflect.Value, fn func(input string)) {
	{
		for _, key := range value.MapKeys() {
			Walk(value.MapIndex(key), fn)
		}
	}
}

func walkArray(value reflect.Value, fn func(input string)) {
	{
		for i := 0; i < value.Len(); i++ {
			Walk(value.Index(i), fn)
		}
	}
}

func walkStruct(value reflect.Value, fn func(input string)) {
	{
		for i := 0; i < value.NumField(); i++ {
			Walk(value.Field(i), fn)
		}
	}
}

func getValue(x any) reflect.Value {
	if v, ok := x.(reflect.Value); ok {
		return v
	}
	value := reflect.ValueOf(x)
	if value.Kind() == reflect.Pointer {
		return value.Elem()
	}
	return value
}
