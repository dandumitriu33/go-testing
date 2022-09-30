package reflection

import "reflect"

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	// Back to classic - refactor
	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walkValue(v)
		}
	}

	// // Different refactor
	// numberOfValues := 0
	// var getField func(int) reflect.Value

	// switch val.Kind() {
	// case reflect.String:
	// 	fn(val.String())
	// case reflect.Struct:
	// 	numberOfValues = val.NumField()
	// 	getField = val.Field
	// case reflect.Slice, reflect.Array:
	// 	numberOfValues = val.Len()
	// 	getField = val.Index
	// case reflect.Map:
	// 	for _, key := range val.MapKeys() {
	// 		walk(val.MapIndex(key).Interface(), fn)
	// 	}
	// }

	// for i := 0; i < numberOfValues; i++ {
	// 	walk(getField(i).Interface(), fn)
	// }

	// // Classic refactor
	// switch val.Kind() {
	// case reflect.Struct:
	// 	for i := 0; i < val.NumField(); i++ {
	// 		walk(val.Field(i).Interface(), fn)
	// 	}
	// case reflect.Slice:
	// 	for i := 0; i < val.Len(); i++ {
	// 		walk(val.Index(i).Interface(), fn)
	// 	}
	// case reflect.String:
	// 	fn(val.String())
	// }

	// if val.Kind() == reflect.Slice {
	// 	for i := 0; i < val.Len(); i++ {
	// 		walk(val.Index(i).Interface(), fn)
	// 	}
	// 	return
	// }

	// for i := 0; i < val.NumField(); i++ {
	// 	field := val.Field(i)

	// 	switch field.Kind() {
	// 	case reflect.String:
	// 		fn(field.String())
	// 	case reflect.Struct:
	// 		walk(field.Interface(), fn)
	// 	}
	// 	// if field.Kind() == reflect.String {
	// 	// 	fn(field.String())
	// 	// }
	// 	// if field.Kind() == reflect.Struct {
	// 	// 	walk(field.Interface(), fn)
	// 	// }
	// }
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
