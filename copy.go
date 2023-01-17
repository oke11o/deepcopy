package deepcopy

import "reflect"

// CopySlice copies a slice of any type.
func CopySlice(in []any) []any {
	out := make([]any, len(in))
	for i, v := range in {
		out[i] = Copy(v)
	}
	return out
}

// Copy copies any type.
func Copy(source any) any {
	sourceValue := reflect.ValueOf(source)
	if sourceValue.Kind() != reflect.Ptr {
		return sourceValue.Interface()
	}

	sourceType := sourceValue.Type()
	destinationElem := reflect.New(sourceType).Elem()
	destination := destinationElem.Interface()

	if sourceValue.IsNil() {
		return destination
	}

	sourceValueElem := sourceValue.Elem()

	newValue := reflect.New(sourceValueElem.Type())
	newValueElem := newValue.Elem()
	newValueElem.Set(sourceValueElem)

	destinationElem.Set(newValue)

	return destinationElem.Interface()
}
