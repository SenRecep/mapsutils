package mapsutils_test

import (
	"github.com/SenRecep/mapsutils"
	"reflect"
	"testing"
)

func TestDeepEqualMap(t *testing.T) {
	testCases := []struct {
		left   map[interface{}]interface{}
		right  map[interface{}]interface{}
		expect bool
	}{
		{
			left:   map[interface{}]interface{}{"a": 1, "b": 2},
			right:  map[interface{}]interface{}{"a": 1, "b": 2},
			expect: true,
		},
		{
			left:   map[interface{}]interface{}{"a": 1, "b": 2},
			right:  map[interface{}]interface{}{"a": 1, "b": 3},
			expect: false,
		},
	}

	for _, tc := range testCases {
		result, _ := mapsutils.DeepEqualMap(tc.left, tc.right)
		if result != tc.expect {
			t.Errorf("Expected %v, but got %v for %v and %v", tc.expect, result, tc.left, tc.right)
		}
	}
}

func TestMapDeleteKey(t *testing.T) {
	testCases := []struct {
		inputMap map[interface{}]interface{}
		key      interface{}
		expect   map[interface{}]interface{}
	}{
		{
			inputMap: map[interface{}]interface{}{"a": 1, "b": 2},
			key:      "a",
			expect:   map[interface{}]interface{}{"b": 2},
		},
		// Add more test cases here
	}

	for _, tc := range testCases {
		err := mapsutils.MapDeleteKey(tc.inputMap, tc.key)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !reflect.DeepEqual(tc.inputMap, tc.expect) {
			t.Errorf("Expected %v, but got %v after deleting key %v from map %v", tc.expect, tc.inputMap, tc.key, tc.inputMap)
		}
	}
}

func TestIsZeroOfUnderlyingType(t *testing.T) {
	testCases := []struct {
		input  interface{}
		expect bool
	}{
		{
			input:  0,
			expect: true,
		},
		{
			input:  "",
			expect: true,
		},
		// Add more test cases here
	}

	for _, tc := range testCases {
		result := mapsutils.IsZeroOfUnderlyingType(tc.input)
		if result != tc.expect {
			t.Errorf("Expected %v, but got %v for input %v", tc.expect, result, tc.input)
		}
	}
}
