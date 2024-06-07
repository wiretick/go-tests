package reflect

import "reflect"

func Walk(x interface{}, fn func(string)) {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.Struct:
		for i := range val.NumField() {
			Walk(val.Field(i).Interface(), fn)
		}
	case reflect.Slice:
		for i := range val.Len() {
			Walk(val.Index(i).Interface(), fn)
		}
	case reflect.String:
		fn(val.String())
	}
}
