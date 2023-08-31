package mapsutils

import (
	"errors"
	"reflect"
)

func DeepEqualMap[TKey comparable, TValue comparable](left map[TKey]TValue, right map[TKey]TValue) (bool, error) {
	if len(left) != len(right) {
		return false, errors.New("not equal length")
	}
	for lKey, lValue := range left {
		if _, ok := right[lKey]; !ok {
			return false, errors.New("element not found")
		}
		if lValue != right[lKey] {
			return false, errors.New("not equal value")
		}
	}
	return true, nil
}

func MapDeleteKey[TKey comparable, TValue any](m map[TKey]TValue, key TKey) error {
	if IsZeroOfUnderlyingType(m) {
		return errors.New("map is default value")
	}

	if IsZeroOfUnderlyingType(key) {
		return errors.New("key is default value")
	}
	if _, ok := m[key]; !ok {
		return errors.New("element not found")
	}

	delete(m, key)
	return nil
}

func IsZeroOfUnderlyingType(x interface{}) bool {
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}
