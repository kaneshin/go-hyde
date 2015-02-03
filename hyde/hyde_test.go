package hyde

import (
	"reflect"
	"testing"
)

func TestFromHyde(t *testing.T) {
	var (
		result   float64
		expected float64
	)

	expected = 156.0
	result = FromHyde(1.0)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected %v, but %v:", expected, result)
	}

	expected = 199.68
	result = FromHyde(1.28)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected %v, but %v:", expected, result)
	}
}

func TestToHyde(t *testing.T) {
	var (
		result   float64
		expected float64
	)

	expected = 1.0
	result = ToHyde(156.0)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected %v, but %v:", expected, result)
	}

	expected = 1.28
	result = ToHyde(199.68)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected %v, but %v:", expected, result)
	}
}
