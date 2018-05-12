package main

import (
	"reflect"
)

type collector interface {
	GetLength() int
	GetItem(i int) reflect.Value
}

type functionType interface{}
// Collection interface
type Collection []interface{}

// GetLength returns the type and kind of the items
// in the collection interface
func(c Collection) GetLength() int {
	val := reflect.ValueOf(c[0])
	return val.Len()
}

// GetItem returns an item 
func(c Collection) GetItem(i int) (reflect.Value) {
	val := reflect.ValueOf(c[0])
	return val.Index(i)
}

// MakeFunc creates a reflect.MakeFunc 
func MakeFunc(f interface{}, value reflect.Value) ([]reflect.Value) {
	funcValue := reflect.ValueOf(f)
	result := funcValue.Call([]reflect.Value{value})
	return result
}

// EachSeries performs series parallel
// operation on input slice
func EachSeries(collection collector, function interface{}) interface{} {
	funcType := reflect.TypeOf(function)
	arrayLength := collection.GetLength()

	if !IsValidFunction(function) {
		panic("expects a function with one argument and one return value")
	}
	
	if !IsSlice(collection) && !IsArray(collection) {
		panic("expects an array or slice of values")
	}

	resultSliceType := reflect.SliceOf(funcType.Out(0))
	resultSlice := reflect.MakeSlice(resultSliceType, 0, arrayLength)

	for i := 0; i < collection.GetLength(); i++ {
		result := MakeFunc(function,  collection.GetItem(i))
		resultSlice = reflect.Append(resultSlice, result[0])
	}
	return resultSlice.Interface()
}