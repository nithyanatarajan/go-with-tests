package reflection

import "reflect"

func Walk(x any, fn func(input string)) {
	value := getValue(x)
	switch value.Kind() {
	case reflect.String:
		fn(value.String())
	case reflect.Struct:
		{
			for i := 0; i < value.NumField(); i++ {
				Walk(value.Field(i), fn)
			}
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
