package typeparameters

import (
	"reflect"
	"testing"
)

func TestMapFunctions(t *testing.T) {
	testCases := []struct {
		name     string
		input    interface{}
		expected []interface{}
		f        interface{}
	}{
		{
			name:     "MapStrings",
			input:    []string{"apple", "banana", "cherry"},
			expected: []interface{}{5, 6, 6},
			f: func(input interface{}) interface{} {
				return len(input.(string))
			},
		},
		{
			name:     "MapInts",
			input:    []int{1, 2, 3, 4},
			expected: []interface{}{2, 4, 6, 8},
			f: func(input interface{}) interface{} {
				return input.(int) * 2
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := callMapFunction(testCase.input, testCase.f)
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("Expected %v, but got %v", testCase.expected, result)
			}
		})
	}
}

func callMapFunction(input interface{}, f interface{}) []interface{} {
	inputValue := reflect.ValueOf(input)
	funcValue := reflect.ValueOf(f)

	if inputValue.Kind() != reflect.Slice {
		panic("Input must be a slice")
	}

	if funcValue.Kind() != reflect.Func {
		panic("Function must be a function")
	}

	result := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(f).Out(0)), inputValue.Len(), inputValue.Len())

	for i := 0; i < inputValue.Len(); i++ {
		params := []reflect.Value{inputValue.Index(i)}
		output := funcValue.Call(params)[0]
		result.Index(i).Set(output)
	}

	resultInterface := make([]interface{}, result.Len())
	for i := 0; i < result.Len(); i++ {
		resultInterface[i] = result.Index(i).Interface()
	}

	return resultInterface
}
