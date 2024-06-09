package reflect

import "reflect"

func Walk(x interface{}, fn func(string)) {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := range val.NumField() {
			Walk(val.Field(i).Interface(), fn)
		}
	case reflect.Slice, reflect.Array:
		for i := range val.Len() {
			Walk(val.Index(i).Interface(), fn)
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			Walk(val.MapIndex(k).Interface(), fn)
		}
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				Walk(v.Interface(), fn)
			} else {
				break
			}
		}
	}
}
