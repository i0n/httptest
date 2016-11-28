package utils

import "reflect"

// StringInSlice check string against strings in a slice returning true or false
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// StringInSliceMapKey check string against value of a key in a slice of string maps
func StringInSliceMapKey(a string, m []map[string]string, key string) bool {
	for _, b := range m {
		if b[key] == a {
			return true
		}
	}
	return false
}

// StringInSliceInterfaceKey check string against the value of a kay in a slice of empty interfaces
func StringInSliceInterfaceKey(a interface{}, m []map[string]string, key string) bool {
	switch reflect.TypeOf(a).Kind() {
	case reflect.String:
		x := reflect.ValueOf(a)
		for _, b := range m {
			if b[key] == x.String() {
				return true
			}
		}
		return false
	default:
		return false
	}
}
