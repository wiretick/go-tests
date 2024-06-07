package reflect

import "reflect"

func Walk(x interface{}, fn func(string)) {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	count := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		getField = val.Field
		count = val.NumField()
	case reflect.Slice, reflect.Array:
		getField = val.Index
		count = val.Len()
	}

	for i := range count {
		Walk(getField(i).Interface(), fn)
	}
}
