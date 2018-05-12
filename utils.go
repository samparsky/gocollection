package main

import(
	"reflect"
)
// IsValidFunction confirms wether the second argument is 
// a function
func IsValidFunction( function interface{})(bool){
	rf := reflect.TypeOf(function)
	if rf.Kind() != reflect.Func || rf.NumIn() != 1  || rf.NumOut() != 1 {
		return false
	}
	return true
}

// IsSlice verify if input is a valid slice
func IsSlice(collection collector) bool {
	rf := reflect.TypeOf(collection)
	if rf.Kind() != reflect.Slice {
		return false
	}
	return true
}

// IsArray verify if input is a valid array
func IsArray(collection collector) bool {
	rf := reflect.TypeOf(collection)
	if rf.Kind() != reflect.Array {
		return false
	}
	return true
}
